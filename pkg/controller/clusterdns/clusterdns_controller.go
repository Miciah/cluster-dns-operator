package clusterdns

import (
	"context"
	"fmt"

	dnsv1alpha1 "github.com/openshift/cluster-dns-operator/pkg/apis/dns/v1alpha1"
	"github.com/openshift/cluster-dns-operator/pkg/manifests"
	"github.com/openshift/cluster-dns-operator/pkg/util"
	"github.com/openshift/cluster-dns-operator/pkg/util/slice"
	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller"
	// "sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	logf "sigs.k8s.io/controller-runtime/pkg/runtime/log"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller_clusterdns")

// Add creates a new ClusterDNS Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func Add(mgr manager.Manager) error {
	return add(mgr, newReconciler(mgr))
}

// newReconciler returns a new reconcile.Reconciler
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	client := mgr.GetClient()

	ic, err := util.GetInstallConfig(client)
	if err != nil {
		log.Error(err, "couldn't get installconfig")
	}

	return &ReconcileClusterDNS{
		client:          client,
		scheme:          mgr.GetScheme(),
		manifestFactory: manifests.NewFactory(),
		installConfig:   ic,
	}
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	// Create a new controller
	c, err := controller.New("clusterdns-controller", mgr, controller.Options{Reconciler: r})
	if err != nil {
		return err
	}

	// Watch for changes to primary resource ClusterDNS
	err = c.Watch(&source.Kind{Type: &dnsv1alpha1.ClusterDNS{}}, &handler.EnqueueRequestForObject{})
	if err != nil {
		return err
	}

	// Watch for changes to secondary resource DaemonSets and requeue the owner ClusterDNS
	err = c.Watch(&source.Kind{Type: &appsv1.DaemonSet{}}, &handler.EnqueueRequestForOwner{
		IsController: true,
		OwnerType:    &dnsv1alpha1.ClusterDNS{},
	})
	if err != nil {
		return err
	}

	return nil
}

var _ reconcile.Reconciler = &ReconcileClusterDNS{}

// ReconcileClusterDNS reconciles a ClusterDNS object.
type ReconcileClusterDNS struct {
	// This client, initialized using mgr.Client() above, is a split client
	// that reads objects from the cache and writes to the apiserver
	client client.Client
	scheme *runtime.Scheme

	manifestFactory *manifests.Factory
	installConfig   *util.InstallConfig
}

// ClusterDNSFinalizer is applied to all ClusterDNS resources before they are
// considered for processing; this ensures the operator has a chance to handle
// all states.
// TODO: Make this generic and not tied to the "default" clusterdns.
const ClusterDNSFinalizer = "dns.openshift.io/default-cluster-dns"

// Reconcile reads that state of the cluster for a ClusterDNS object and makes
// changes based on the state read and what is in the ClusterDNS.Spec.
func (r *ReconcileClusterDNS) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	reqLogger := log.WithValues("Request.Namespace", request.Namespace, "Request.Name", request.Name)
	reqLogger.Info("Reconciling ClusterDNS")

	defer r.syncOperatorStatus(request.Namespace)

	// Fetch the ClusterDNS instance
	instance := &dnsv1alpha1.ClusterDNS{}
	err := r.client.Get(context.TODO(), request.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}

	// Ensure we have all the necessary scaffolding on which to place DNS
	// instances.
	if err := r.ensureDNSNamespace(); err != nil {
		return reconcile.Result{}, err
	}

	// Handle deleted dns.
	// TODO: Assert/ensure that the dns has a finalizer so we can reliably
	// detect deletion.
	if instance.DeletionTimestamp != nil {
		// Destroy any coredns instance associated with the dns.
		if err := r.ensureDNSDeleted(instance); err != nil {
			return reconcile.Result{}, fmt.Errorf("couldn't delete clusterdns %q: %v", instance.Name, err)
		}

		// Clean up the finalizer to allow the clusterdns to be deleted.
		if slice.ContainsString(instance.Finalizers, ClusterDNSFinalizer) {
			instance.Finalizers = slice.RemoveString(instance.Finalizers, ClusterDNSFinalizer)
			if err = r.client.Update(context.TODO(), instance); err != nil {
				return reconcile.Result{}, fmt.Errorf("couldn't remove finalizer from clusterdns %q: %v", instance.Name, err)
			}
		}

		return reconcile.Result{}, nil
	}

	// Handle active DNS.
	if err := r.ensureCoreDNSForClusterDNS(instance); err != nil {
		return reconcile.Result{}, fmt.Errorf("couldn't ensure clusterdns %q: %v", instance.Name, err)
	}

	return reconcile.Result{}, nil
}

// ensureDNSNamespace ensures all the necessary scaffolding exists for
// CoreDNS generally, including a namespace and all RBAC setup.
func (r *ReconcileClusterDNS) ensureDNSNamespace() error {
	ns, err := r.manifestFactory.DNSNamespace()
	if err != nil {
		return fmt.Errorf("couldn't build dns namespace: %v", err)
	}
	err = r.client.Create(context.TODO(), ns)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("couldn't create dns namespace: %v", err)
	}

	sa, err := r.manifestFactory.DNSServiceAccount()
	if err != nil {
		return fmt.Errorf("couldn't build dns service account: %v", err)
	}
	err = r.client.Create(context.TODO(), sa)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("couldn't create dns service account: %v", err)
	}

	cr, err := r.manifestFactory.DNSClusterRole()
	if err != nil {
		return fmt.Errorf("couldn't build dns cluster role: %v", err)
	}
	err = r.client.Create(context.TODO(), cr)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("couldn't create dns cluster role: %v", err)
	}

	crb, err := r.manifestFactory.DNSClusterRoleBinding()
	if err != nil {
		return fmt.Errorf("couldn't build dns cluster role binding: %v", err)
	}
	err = r.client.Create(context.TODO(), crb)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("couldn't create dns cluster role binding: %v", err)
	}

	return nil
}

// ensureCoreDNSForClusterDNS ensures all necessary CoreDNS resources exist for
// a given clusterdns.
func (r *ReconcileClusterDNS) ensureCoreDNSForClusterDNS(dns *dnsv1alpha1.ClusterDNS) error {
	ds, err := r.manifestFactory.DNSDaemonSet(dns)
	if err != nil {
		return fmt.Errorf("couldn't build daemonset: %v", err)
	}
	err = r.client.Create(context.TODO(), ds)
	if errors.IsAlreadyExists(err) {
		if err = r.client.Get(context.TODO(), types.NamespacedName{
			Name:      ds.Name,
			Namespace: ds.Namespace,
		}, ds); err != nil {
			return fmt.Errorf("failed to fetch daemonset %s, %v", ds.Name, err)
		}
	} else if err != nil {
		return fmt.Errorf("failed to create daemonset: %v", err)
	}
	trueVar := true
	dsRef := metav1.OwnerReference{
		APIVersion: ds.APIVersion,
		Kind:       ds.Kind,
		Name:       ds.Name,
		UID:        ds.UID,
		Controller: &trueVar,
	}

	cm, err := r.manifestFactory.DNSConfigMap(dns)
	if err != nil {
		return fmt.Errorf("couldn't build dns config map: %v", err)
	}
	cm.SetOwnerReferences([]metav1.OwnerReference{dsRef})
	err = r.client.Create(context.TODO(), cm)
	if err != nil && !errors.IsAlreadyExists(err) {
		return fmt.Errorf("couldn't create dns config map: %v", err)
	}

	service, err := r.manifestFactory.DNSService(dns)
	if err != nil {
		return fmt.Errorf("couldn't build service: %v", err)
	}
	err = r.client.Get(context.TODO(), types.NamespacedName{
		Name:      service.Name,
		Namespace: service.Namespace,
	}, service)
	if err != nil {
		if !errors.IsNotFound(err) {
			return fmt.Errorf("failed to fetch service %s, %v", service.Name, err)
		}
		service.SetOwnerReferences([]metav1.OwnerReference{dsRef})
		if err = r.client.Create(context.TODO(), service); err != nil {
			return fmt.Errorf("failed to create service: %v", err)
		}
	}

	return nil
}

// ensureDNSDeleted ensures that any CoreDNS resources associated with the
// clusterdns are deleted.
func (r *ReconcileClusterDNS) ensureDNSDeleted(dns *dnsv1alpha1.ClusterDNS) error {
	// DNS specific configmap and service has owner reference to daemonset.
	// So deletion of daemonset will trigger garbage collection of corresponding
	// configmap and service resources.
	ds, err := r.manifestFactory.DNSDaemonSet(dns)
	if err != nil {
		return fmt.Errorf("failed to build daemonset for deletion, ClusterDNS: %q, %v", dns.Name, err)
	}
	err = r.client.Delete(context.TODO(), ds)
	if !errors.IsNotFound(err) {
		return err
	}
	return nil
}
