package controller

import (
	"context"
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sirupsen/logrus"

	operatorv1 "github.com/openshift/api/operator/v1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

// syncDNSStatus computes the current status of dns and
// updates status upon any changes since last sync.
func (r *reconciler) syncDNSStatus(dns *operatorv1.DNS, clusterIP, clusterDomain string, ds *appsv1.DaemonSet, endpoints *corev1.Endpoints) error {
	resolverErrors := checkEndpoints(ds, endpoints)
	updated := dns.DeepCopy()
	updated.Status.ClusterIP = clusterIP
	updated.Status.ClusterDomain = clusterDomain
	updated.Status.Conditions = computeDNSStatusConditions(dns.Status.Conditions, clusterIP, ds, resolverErrors)
	if !dnsStatusesEqual(updated.Status, dns.Status) {
		if err := r.client.Status().Update(context.TODO(), updated); err != nil {
			return fmt.Errorf("failed to update dns status: %v", err)
		}
		logrus.Infof("updated DNS %s status: old: %#v, new: %#v", dns.ObjectMeta.Name, dns.Status, updated.Status)
	}

	return nil
}

// checkEndpoints verifies that the provided endpoints by sending a DNS query
// for the API service host name to each endpoint and returns a list of errors.
func checkEndpoints(ds *appsv1.DaemonSet, endpoints *corev1.Endpoints) []error {
	const apiHost = "kubernetes.default.svc.cluster.local."
	var (
		errs      []error
		addresses []string
	)
	for _, ep := range endpoints.Subsets {
		for _, subset := range ep.Addresses {
			for _, port := range ep.Ports {
				if port.Protocol != corev1.ProtocolUDP {
					continue
				}
				a, p := subset.IP, strconv.Itoa(int(port.Port))
				addrPort := net.JoinHostPort(a, p)
				addresses = append(addresses, addrPort)
			}
		}
	}
	if len(addresses) < int(ds.Status.NumberAvailable) {
		errs = append(errs, fmt.Errorf("DaemonSet reports %d pods, but Endpoints has only %d addresses", ds.Status.NumberAvailable, len(addresses)))
	}
	// TODO: Parallelize queries.
	for _, address := range addresses {
		resolver := &net.Resolver{
			PreferGo: true,
			Dial: func(ctx context.Context, network_, address_ string) (net.Conn, error) {
				dialer := net.Dialer{
					Timeout: time.Millisecond * 500,
				}
				return dialer.DialContext(ctx, "udp", address)
			},
		}
		if _, err := resolver.LookupHost(context.Background(), apiHost); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

// computeDNSStatusConditions computes dns status conditions based on the status
// of ds and clusterIP and any provided resolver errors.
func computeDNSStatusConditions(oldConditions []operatorv1.OperatorCondition, clusterIP string, ds *appsv1.DaemonSet, resolverErrors []error) []operatorv1.OperatorCondition {
	var oldDegradedCondition, oldProgressingCondition, oldAvailableCondition *operatorv1.OperatorCondition
	for i := range oldConditions {
		switch oldConditions[i].Type {
		case operatorv1.OperatorStatusTypeDegraded:
			oldDegradedCondition = &oldConditions[i]
		case operatorv1.OperatorStatusTypeProgressing:
			oldProgressingCondition = &oldConditions[i]
		case operatorv1.OperatorStatusTypeAvailable:
			oldAvailableCondition = &oldConditions[i]
		}
	}

	conditions := []operatorv1.OperatorCondition{
		computeDNSDegradedCondition(oldDegradedCondition, clusterIP, ds, resolverErrors),
		computeDNSProgressingCondition(oldProgressingCondition, clusterIP, ds),
		computeDNSAvailableCondition(oldAvailableCondition, clusterIP, ds),
	}

	return conditions
}

// computeDNSDegradedCondition computes the dns Degraded status condition
// based on the status of clusterIP and ds and any provided resolver errors.
func computeDNSDegradedCondition(oldCondition *operatorv1.OperatorCondition, clusterIP string, ds *appsv1.DaemonSet, resolverErrors []error) operatorv1.OperatorCondition {
	var (
		degradedCondition = &operatorv1.OperatorCondition{
			Type:   operatorv1.OperatorStatusTypeDegraded,
			Status: operatorv1.ConditionUnknown,
		}
		reasons []string
		errs    []error
	)

	if len(resolverErrors) != 0 {
		reasons = append(reasons, "EndpointsUnavailable")
		errs = append(errs, resolverErrors...)
	}

	numberUnavailable := ds.Status.DesiredNumberScheduled - ds.Status.NumberAvailable
	if len(clusterIP) == 0 {
		reasons = append(reasons, "NoServiceIP")
		errs = append(errs, errors.New("No IP assigned to DNS service"))
	}
	if ds.Status.DesiredNumberScheduled == 0 {
		reasons = append(reasons, "NoPodsDesired")
		errs = append(errs, errors.New("No CoreDNS pods are scheduled"))
	} else if ds.Status.NumberAvailable == 0 {
		reasons = append(reasons, "NoPodsAvailable")
		errs = append(errs, errors.New("No CoreDNS pods are available"))
	} else if numberUnavailable > ds.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable.IntVal {
		reasons = append(reasons, "MaxUnavailableExceeded")
		errs = append(errs, fmt.Errorf("Too many unavailable CoreDNS pods (%d > %d max unavailable)", numberUnavailable, ds.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable.IntVal))
	}

	if len(errs) != 0 {
		degradedCondition.Status = operatorv1.ConditionTrue
		degradedCondition.Reason = strings.Join(reasons, "And")
		degradedCondition.Message = utilerrors.NewAggregate(errs).Error()
	} else {
		degradedCondition.Status = operatorv1.ConditionFalse
		degradedCondition.Reason = "AsExpected"
		degradedCondition.Message = "IP assigned to DNS service and minimum DaemonSet pods running"
	}

	return setDNSLastTransitionTime(degradedCondition, oldCondition)
}

// computeDNSProgressingCondition computes the dns Progressing status condition
// based on the status of ds.
func computeDNSProgressingCondition(oldCondition *operatorv1.OperatorCondition, clusterIP string, ds *appsv1.DaemonSet) operatorv1.OperatorCondition {
	progressingCondition := &operatorv1.OperatorCondition{
		Type: operatorv1.OperatorStatusTypeProgressing,
	}
	switch {
	case len(clusterIP) == 0 && ds.Status.NumberAvailable != ds.Status.DesiredNumberScheduled:
		progressingCondition.Status = operatorv1.ConditionTrue
		progressingCondition.Reason = "Reconciling"
		progressingCondition.Message = fmt.Sprintf("No IP assigned to DNS service and %d Nodes running a DaemonSet pod, want %d",
			ds.Status.NumberAvailable, ds.Status.DesiredNumberScheduled)
	case len(clusterIP) == 0:
		progressingCondition.Status = operatorv1.ConditionTrue
		progressingCondition.Reason = "Reconciling"
		progressingCondition.Message = "No IP assigned to DNS service"
	case ds.Status.NumberAvailable != ds.Status.DesiredNumberScheduled:
		progressingCondition.Status = operatorv1.ConditionTrue
		progressingCondition.Reason = "Reconciling"
		progressingCondition.Message = fmt.Sprintf("%d Nodes running a DaemonSet pod, want %d",
			ds.Status.NumberAvailable, ds.Status.DesiredNumberScheduled)
	default:
		progressingCondition.Status = operatorv1.ConditionFalse
		progressingCondition.Reason = "AsExpected"
		progressingCondition.Message = "All expected Nodes running DaemonSet pod and IP assigned to DNS service"
	}

	return setDNSLastTransitionTime(progressingCondition, oldCondition)
}

// computeDNSAvailableCondition computes the dns Available status condition
// based on the status of clusterIP and ds.
func computeDNSAvailableCondition(oldCondition *operatorv1.OperatorCondition, clusterIP string, ds *appsv1.DaemonSet) operatorv1.OperatorCondition {
	availableCondition := &operatorv1.OperatorCondition{
		Type: operatorv1.OperatorStatusTypeAvailable,
	}
	switch {
	case len(clusterIP) == 0 && ds.Status.NumberAvailable == 0:
		availableCondition.Status = operatorv1.ConditionFalse
		availableCondition.Reason = "NoServiceIPAndNoDaemonSetPods"
		availableCondition.Message = "No IP assigned to DNS service and no running DaemonSet pods"
	case len(clusterIP) == 0:
		availableCondition.Status = operatorv1.ConditionFalse
		availableCondition.Reason = "NoServiceIP"
		availableCondition.Message = "No IP assigned to DNS service"
	case ds.Status.NumberAvailable == 0:
		availableCondition.Status = operatorv1.ConditionFalse
		availableCondition.Reason = "DaemonSetUnavailable"
		availableCondition.Message = "DaemonSet pod not running on any Nodes"
	default:
		availableCondition.Status = operatorv1.ConditionTrue
		availableCondition.Reason = "AsExpected"
		availableCondition.Message = "Minimum number of Nodes running DaemonSet pod and IP assigned to DNS service"
	}

	return setDNSLastTransitionTime(availableCondition, oldCondition)
}

// setDNSLastTransitionTime sets LastTransitionTime for the given condition.
// If the condition has changed, it will assign a new timestamp otherwise keeps the old timestamp.
func setDNSLastTransitionTime(condition, oldCondition *operatorv1.OperatorCondition) operatorv1.OperatorCondition {
	if oldCondition != nil && condition.Status == oldCondition.Status &&
		condition.Reason == oldCondition.Reason && condition.Message == oldCondition.Message {
		condition.LastTransitionTime = oldCondition.LastTransitionTime
	} else {
		condition.LastTransitionTime = metav1.Now()
	}

	return *condition
}

// dnsStatusesEqual compares two DNSStatus values.  Returns true
// if the provided values should be considered equal for the purpose of determining
// whether an update is necessary, false otherwise.
func dnsStatusesEqual(a, b operatorv1.DNSStatus) bool {
	conditionCmpOpts := []cmp.Option{
		cmpopts.EquateEmpty(),
		cmpopts.SortSlices(func(a, b operatorv1.OperatorCondition) bool { return a.Type < b.Type }),
	}
	if !cmp.Equal(a.Conditions, b.Conditions, conditionCmpOpts...) {
		return false
	}
	if a.ClusterIP != b.ClusterIP {
		return false
	}
	if a.ClusterDomain != b.ClusterDomain {
		return false
	}

	return true
}
