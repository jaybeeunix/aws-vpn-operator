package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CustomerGatewaySpec defines the desired state of CustomerGateway
// +k8s:openapi-gen=true
type CustomerGatewaySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// Name for the Customer Gateway.
	Name string `json:"name"`

	// AWS Region for the Customer Gateway.
	Region string `json:"region"`

	// Specify the Internet-routable IP address for your gateway's external interface; the address must be static and may be behind a device performing network address translation (NAT).
	PublicIP string `json:"publicIP"`

	// The type of VPN connection that this customer gateway supports (ipsec.1).
	Type string `json:"type"`

	// The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your customer gateway. You can use an existing ASN assigned to your network. If you do not have one, you can use a private ASN in the 64512-65534 range. Default: 65000
	// +optional
	BgpAsn int64 `json:"bgpAsn,omitempty"`
}

// Tag assigned to the customer gateway.
type Tag struct {
	// The key of the tag.
	Key string `json:"key,omitempty"`

	//The value of the tag.
	Value string `json:"value,omitempty"`
}

// CustomerGatewayStatus defines the observed state of CustomerGateway
// +k8s:openapi-gen=true
type CustomerGatewayStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "operator-sdk generate k8s" to regenerate code after modifying this file
	// Add custom validation using kubebuilder tags: https://book-v1.book.kubebuilder.io/beyond_basics/generating_crd.html

	// The ID of the customer gateway.
	CustomerGatewayId string `json:"customerGatewayId,omitempty"`

	// The Internet-routable IP address of the customer gateway's outside interface.
	IpAddress string `json:"ipAddress,omitempty"`

	// The customer gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).
	BgpAsn string `json:"bgpAsn,omitempty"`

	// The current state of the customer gateway.
	State string `json:"state,omitempty"`

	// The type of VPN connection the customer gateway supports (ipsec.1).
	Type string `json:"type,omitempty"`

	// Tags assigned to the customer gateway.
	Tags []Tag `json:"tags,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomerGateway is the Schema for the customergateways API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type CustomerGateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CustomerGatewaySpec   `json:"spec,omitempty"`
	Status CustomerGatewayStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// CustomerGatewayList contains a list of CustomerGateway
type CustomerGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CustomerGateway `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CustomerGateway{}, &CustomerGatewayList{})
}
