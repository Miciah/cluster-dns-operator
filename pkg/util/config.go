package util

import (
	"context"
	"fmt"

	"github.com/ghodss/yaml"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// ClusterConfigNamespace is the namespace containing the cluster config.
	ClusterConfigNamespace = "kube-system"
	// ClusterConfigName is the name of the cluster config configmap.
	ClusterConfigName = "cluster-config-v1"
	// InstallConfigKey is the key in the cluster config configmap containing a
	// serialized InstallConfig.
	InstallConfigKey = "install-config"
)

type InstallConfig struct {
	Networking NetworkingConfig `json:"networking"`
}

type NetworkingConfig struct {
	ServiceCIDR string `json:"serviceCIDR"`
}

// UnmarshalInstallConfig builds an install config from the cluster config.
func UnmarshalInstallConfig(clusterConfig *corev1.ConfigMap) (*InstallConfig, error) {
	icJson, ok := clusterConfig.Data[InstallConfigKey]
	if !ok {
		return nil, fmt.Errorf("missing %q in configmap", InstallConfigKey)
	}
	var ic InstallConfig
	if err := yaml.Unmarshal([]byte(icJson), &ic); err != nil {
		return nil, fmt.Errorf("invalid InstallConfig: %v\njson:\n%s", err, icJson)
	}
	return &ic, nil
}

// GetInstallConfig looks up the install config in the cluster.
func GetInstallConfig(client client.Client) (*InstallConfig, error) {
	cm := &corev1.ConfigMap{}
	if err := client.Get(context.TODO(), types.NamespacedName{
		Name:      ClusterConfigName,
		Namespace: ClusterConfigNamespace,
	}, cm); err != nil {
		return nil, fmt.Errorf("couldn't get clusterconfig %s/%s: %v", ClusterConfigNamespace, ClusterConfigName, err)
	}

	return UnmarshalInstallConfig(cm)
}
