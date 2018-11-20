package clusterdns

import (
	"context"
	"fmt"
	"strings"

	osv1 "github.com/openshift/api/config/v1"
	dnsv1alpha1 "github.com/openshift/cluster-dns-operator/pkg/apis/dns/v1alpha1"
	"github.com/openshift/cluster-dns-operator/pkg/util/clusteroperator"
	operatorversion "github.com/openshift/cluster-dns-operator/version"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// TODO Make this a controller?

const (
	clusterOperatorName      = "openshift-dns"
	clusterOperatorNamespace = "openshift-cluster-dns-operator"
)

// syncOperatorStatus computes the operator's current status and therefrom
// creates or updates the ClusterOperator resource for the operator.
func (r *ReconcileClusterDNS) syncOperatorStatus(namespace string) {
	log = log.WithValues(
		"Status.Namespace", clusterOperatorNamespace,
		"Status.Name", clusterOperatorName,
	)

	co := &osv1.ClusterOperator{}
	err := r.client.Get(context.TODO(), types.NamespacedName{
		Name:      clusterOperatorName,
		Namespace: clusterOperatorNamespace,
	}, co)
	isNotFound := errors.IsNotFound(err)
	if err != nil && !isNotFound {
		log.Error(err, "syncOperatorStatus: error getting ClusterOperator")

		return
	}

	ns, dnses, daemonsets, err := r.getOperatorState(namespace)
	if err != nil {
		log.Error(err, "syncOperatorStatus: getOperatorState")

		return
	}

	oldConditions := co.Status.Conditions
	co.Status.Conditions = computeStatusConditions(oldConditions, ns,
		dnses, daemonsets)

	if isNotFound {
		co.Status.Version = operatorversion.Version

		if err := r.client.Create(context.TODO(), co); err != nil {
			log.Error(err, "syncOperatorStatus: failed to create ClusterOperator")
		} else {
			log.Info("syncOperatorStatus: created ClusterOperator", "Status.UID", co.UID)
		}
	}

	if clusteroperator.ConditionsEqual(oldConditions, co.Status.Conditions) {
		return
	}

	if err := r.client.Status().Update(context.TODO(), co); err != nil {
		log.Error(err, "syncOperatorStatus: error updating status")
	}
}

// getOperatorState gets and returns the resources necessary to compute the
// operator's current state.
func (r *ReconcileClusterDNS) getOperatorState(namespace string) (*corev1.Namespace, []dnsv1alpha1.ClusterDNS, []appsv1.DaemonSet, error) {
	ns, err := r.manifestFactory.DNSNamespace()
	if err != nil {
		return nil, nil, nil, fmt.Errorf("error building Namespace: %v",
			err)
	}

	if err := r.client.Get(context.TODO(), types.NamespacedName{
		Name:      ns.Name,
		Namespace: ns.Namespace,
	}, ns); err != nil {
		if errors.IsNotFound(err) {
			return nil, nil, nil, nil
		}

		return nil, nil, nil, fmt.Errorf(
			"error getting Namespace %s: %v", ns.Name, err)
	}

	listOpts := &client.ListOptions{Namespace: namespace}

	dnsList := &dnsv1alpha1.ClusterDNSList{}
	if err := r.client.List(context.TODO(), listOpts, dnsList); err != nil {
		return nil, nil, nil, fmt.Errorf(
			"failed to list ClusterDNSes: %v", err)
	}

	daemonsetList := &appsv1.DaemonSetList{}
	if err := r.client.List(context.TODO(), listOpts, daemonsetList); err != nil {
		return nil, nil, nil, fmt.Errorf(
			"failed to list DaemonSets: %v", err)
	}

	return ns, dnsList.Items, daemonsetList.Items, nil
}

// computeStatusConditions computes the operator's current state.
func computeStatusConditions(conditions []osv1.ClusterOperatorStatusCondition, ns *corev1.Namespace, dnses []dnsv1alpha1.ClusterDNS, daemonsets []appsv1.DaemonSet) []osv1.ClusterOperatorStatusCondition {
	failingCondition := &osv1.ClusterOperatorStatusCondition{
		Type:   osv1.OperatorFailing,
		Status: osv1.ConditionUnknown,
	}
	if ns == nil {
		failingCondition.Status = osv1.ConditionTrue
		failingCondition.Reason = "NoNamespace"
		failingCondition.Message = "DNS namespace does not exist"
	} else {
		failingCondition.Status = osv1.ConditionFalse
	}
	conditions = clusteroperator.SetStatusCondition(conditions,
		failingCondition)

	progressingCondition := &osv1.ClusterOperatorStatusCondition{
		Type:   osv1.OperatorProgressing,
		Status: osv1.ConditionUnknown,
	}
	numClusterDNSes := len(dnses)
	numDaemonsets := len(daemonsets)
	if numClusterDNSes == numDaemonsets {
		progressingCondition.Status = osv1.ConditionFalse
	} else {
		progressingCondition.Status = osv1.ConditionTrue
		progressingCondition.Reason = "Reconciling"
		progressingCondition.Message = fmt.Sprintf(
			"have %d DaemonSets, want %d",
			numDaemonsets, numClusterDNSes)
	}
	conditions = clusteroperator.SetStatusCondition(conditions,
		progressingCondition)

	availableCondition := &osv1.ClusterOperatorStatusCondition{
		Type:   osv1.OperatorAvailable,
		Status: osv1.ConditionUnknown,
	}
	dsAvailable := map[string]bool{}
	for _, ds := range daemonsets {
		dsAvailable[ds.Name] = ds.Status.NumberAvailable > 0
	}
	unavailable := []string{}
	for _, dns := range dnses {
		// TODO Use the manifest to derive the name, or use labels or
		// owner references.
		name := "dns-" + dns.Name
		if available, exists := dsAvailable[name]; !exists {
			msg := fmt.Sprintf("no DaemonSet for ClusterDNS %q",
				dns.Name)
			unavailable = append(unavailable, msg)
		} else if !available {
			msg := fmt.Sprintf("DaemonSet %q not available", name)
			unavailable = append(unavailable, msg)
		}
	}
	if len(unavailable) == 0 {
		availableCondition.Status = osv1.ConditionTrue
	} else {
		availableCondition.Status = osv1.ConditionFalse
		availableCondition.Reason = "DaemonSetNotAvailable"
		availableCondition.Message = strings.Join(unavailable, "\n")
	}
	conditions = clusteroperator.SetStatusCondition(conditions,
		availableCondition)

	return conditions
}
