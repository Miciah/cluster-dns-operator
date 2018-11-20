package controller

import (
	"github.com/openshift/cluster-dns-operator/pkg/controller/clusterdns"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, clusterdns.Add)
}
