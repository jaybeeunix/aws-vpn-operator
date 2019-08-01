package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VpnGatewaySpec defines the desired state of VpnGateway
// +k8s:openapi-gen=true
type VpnGatewaySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html
}

// VpnGatewayStatus defines the observed state of VpnGateway
// +k8s:openapi-gen=true
type VpnGatewayStatus struct {
	// Phase is the current status of the VPN Gateway
	// +kubebuilder:validation:Enum=Creating,Detached,Ready
	Phase string `json:"phase,omitempty"`

	// VpnGatewayID is the AWS ID of the VPN Gateway object
	VpnGatewayID string `json:"vpnGatewayID,omitempty"`

	// VpcID is the AWS ID of the VPC that the VPN Gateway is attempting to attach to
	VpcID string `json:"vpcID,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VpnGateway is the Schema for the vpngateways API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type VpnGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VpnGatewaySpec   `json:"spec,omitempty"`
	Status VpnGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VpnGatewayList contains a list of VpnGateway
type VpnGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VpnGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VpnGateway{}, &VpnGatewayList{})
}
