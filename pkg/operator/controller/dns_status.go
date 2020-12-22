package controller

import (
	"context"
	"fmt"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sirupsen/logrus"

	operatorv1 "github.com/openshift/api/operator/v1"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// syncDNSStatus computes the current status of dns and
// updates status upon any changes since last sync.
func (r *reconciler) syncDNSStatus(dns *operatorv1.DNS, clusterIP, clusterDomain string, ds *appsv1.DaemonSet, haveResolverDaemonset bool, resolverDaemonset *appsv1.DaemonSet) error {
	updated := dns.DeepCopy()
	updated.Status.ClusterIP = clusterIP
	updated.Status.ClusterDomain = clusterDomain
	updated.Status.Conditions = computeDNSStatusConditions(dns.Status.Conditions, dns, clusterIP, ds, haveResolverDaemonset, resolverDaemonset)
	if !dnsStatusesEqual(updated.Status, dns.Status) {
		if err := r.client.Status().Update(context.TODO(), updated); err != nil {
			return fmt.Errorf("failed to update dns status: %v", err)
		}
		logrus.Infof("updated DNS %s status: old: %#v, new: %#v", dns.ObjectMeta.Name, dns.Status, updated.Status)
	}

	return nil
}

// computeDNSStatusConditions computes dns status conditions based on
// the status of ds and clusterIP.
func computeDNSStatusConditions(oldConditions []operatorv1.OperatorCondition, dns *operatorv1.DNS, clusterIP string, ds *appsv1.DaemonSet, haveResolverDaemonset bool, resolverDaemonset *appsv1.DaemonSet) []operatorv1.OperatorCondition {
	var (
		oldAvailableCondition   *operatorv1.OperatorCondition
		oldDegradedCondition    *operatorv1.OperatorCondition
		oldProgressingCondition *operatorv1.OperatorCondition
		oldUpgradeableCondition *operatorv1.OperatorCondition
	)
	for i := range oldConditions {
		switch oldConditions[i].Type {
		case operatorv1.OperatorStatusTypeDegraded:
			oldDegradedCondition = &oldConditions[i]
		case operatorv1.OperatorStatusTypeProgressing:
			oldProgressingCondition = &oldConditions[i]
		case operatorv1.OperatorStatusTypeAvailable:
			oldAvailableCondition = &oldConditions[i]
		case operatorv1.OperatorStatusTypeUpgradeable:
			oldUpgradeableCondition = &oldConditions[i]
		}
	}

	conditions := []operatorv1.OperatorCondition{
		computeDNSUpgradeableCondition(oldUpgradeableCondition, dns, haveResolverDaemonset, resolverDaemonset),
		computeDNSDegradedCondition(oldDegradedCondition, clusterIP, ds),
		computeDNSProgressingCondition(oldProgressingCondition, clusterIP, ds),
		computeDNSAvailableCondition(oldAvailableCondition, clusterIP, ds),
	}

	return conditions
}

// computeDNSUpgradeableCondition computes the dns Upgradeable status condition
// based on the status of dns and ds.
func computeDNSUpgradeableCondition(oldCondition *operatorv1.OperatorCondition, dns *operatorv1.DNS, haveDS bool, ds *appsv1.DaemonSet) operatorv1.OperatorCondition {
	condition := &operatorv1.OperatorCondition{
		Type:   operatorv1.OperatorStatusTypeUpgradeable,
		Status: operatorv1.ConditionTrue,
		Reason: "AsExpected",
	}
	if haveDS {
		if err := verifyNodeResolverDaemonIsUpgradeable(dns, ds); err != nil {
			condition.Status = operatorv1.ConditionFalse
			condition.Reason = "NotUpgradeable"
			condition.Message = err.Error()
		}
	}
	return setDNSLastTransitionTime(condition, oldCondition)
}

// computeDNSDegradedCondition computes the dns Degraded status condition
// based on the status of clusterIP and ds.
func computeDNSDegradedCondition(oldCondition *operatorv1.OperatorCondition, clusterIP string,
	ds *appsv1.DaemonSet) operatorv1.OperatorCondition {
	degradedCondition := &operatorv1.OperatorCondition{
		Type: operatorv1.OperatorStatusTypeDegraded,
	}
	numberUnavailable := ds.Status.DesiredNumberScheduled - ds.Status.NumberAvailable
	switch {
	case len(clusterIP) == 0 && ds.Status.NumberAvailable == 0:
		degradedCondition.Status = operatorv1.ConditionTrue
		degradedCondition.Reason = "NoServiceIPAndNoDaemonSetPods"
		degradedCondition.Message = "No IP assigned to DNS service and no DaemonSet pods running"
	case len(clusterIP) == 0:
		degradedCondition.Status = operatorv1.ConditionTrue
		degradedCondition.Reason = "NoServiceIP"
		degradedCondition.Message = "No IP assigned to DNS service"
	case ds.Status.DesiredNumberScheduled == 0:
		degradedCondition.Status = operatorv1.ConditionTrue
		degradedCondition.Reason = "NoPodsDesired"
		degradedCondition.Message = "No CoreDNS pods are desired (this could mean nodes are tainted)"
	case ds.Status.NumberAvailable == 0:
		degradedCondition.Status = operatorv1.ConditionTrue
		degradedCondition.Reason = "NoPodsAvailable"
		degradedCondition.Message = "No CoreDNS pods are available"
	case numberUnavailable > ds.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable.IntVal:
		degradedCondition.Status = operatorv1.ConditionTrue
		degradedCondition.Reason = "MaxUnavailableExceeded"
		degradedCondition.Message = fmt.Sprintf("Too many unavailable CoreDNS pods (%d > %d max unavailable)", numberUnavailable, ds.Spec.UpdateStrategy.RollingUpdate.MaxUnavailable.IntVal)
	default:
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
