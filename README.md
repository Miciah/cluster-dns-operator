# The DNS Operator

The DNS Operator deploys and manages [CoreDNS](https://coredns.io) to provide a name resolution service to pods that enables [DNS-based Kubernetes Service discovery](https://kubernetes.io/docs/concepts/services-networking/service/#dns) in [OpenShift](https://openshift.io).

The operator tries to be useful out of the box by creating a working default deployment based on the cluster's configuration.

* The default cluster domain is `cluster.local`.
* Configuration of the CoreDNS [Corefile](https://coredns.io/manual/toc/#configuration) or [kubernetes plugin](https://coredns.io/plugins/kubernetes/) is not yet supported.

## How it works

The DNS Operator manages CoreDNS as a Kubernetes DaemonSet exposed as a Service with a static IP — CoreDNS runs on all nodes in the cluster.

## How to help

See [HACKING.md](HACKING.md) for development topics.

## Reporting issues

Bugs are tracked in [Bugzilla](https://bugzilla.redhat.com/enter_bug.cgi?product=OpenShift%20Container%20Platform&component=DNS).
