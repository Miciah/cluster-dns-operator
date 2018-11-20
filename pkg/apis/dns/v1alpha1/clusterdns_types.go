package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ClusterDNSSpec defines the desired state of ClusterDNS
type ClusterDNSSpec struct {
	ClusterIP *string `json:"clusterIP"`

	ClusterDomain *string `json:"clusterDomain"`
}

// ClusterDNSStatus defines the observed state of ClusterDNS
type ClusterDNSStatus struct {
	// Fill me
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterDNS is the Schema for the clusterdns API
// +k8s:openapi-gen=true
type ClusterDNS struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterDNSSpec   `json:"spec,omitempty"`
	Status ClusterDNSStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterDNSList contains a list of ClusterDNS
type ClusterDNSList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterDNS `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterDNS{}, &ClusterDNSList{})
}
