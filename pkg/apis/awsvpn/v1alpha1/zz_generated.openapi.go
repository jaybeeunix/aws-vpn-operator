// +build !

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/awsvpn/v1alpha1.CustomerGateway":       schema_pkg_apis_awsvpn_v1alpha1_CustomerGateway(ref),
		"./pkg/apis/awsvpn/v1alpha1.CustomerGatewaySpec":   schema_pkg_apis_awsvpn_v1alpha1_CustomerGatewaySpec(ref),
		"./pkg/apis/awsvpn/v1alpha1.CustomerGatewayStatus": schema_pkg_apis_awsvpn_v1alpha1_CustomerGatewayStatus(ref),
		"./pkg/apis/awsvpn/v1alpha1.VpnGateway":            schema_pkg_apis_awsvpn_v1alpha1_VpnGateway(ref),
		"./pkg/apis/awsvpn/v1alpha1.VpnGatewaySpec":        schema_pkg_apis_awsvpn_v1alpha1_VpnGatewaySpec(ref),
		"./pkg/apis/awsvpn/v1alpha1.VpnGatewayStatus":      schema_pkg_apis_awsvpn_v1alpha1_VpnGatewayStatus(ref),
	}
}

func schema_pkg_apis_awsvpn_v1alpha1_CustomerGateway(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CustomerGateway is the Schema for the customergateways API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/awsvpn/v1alpha1.CustomerGatewaySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/awsvpn/v1alpha1.CustomerGatewayStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/awsvpn/v1alpha1.CustomerGatewaySpec", "./pkg/apis/awsvpn/v1alpha1.CustomerGatewayStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_awsvpn_v1alpha1_CustomerGatewaySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CustomerGatewaySpec defines the desired state of CustomerGateway",
				Properties: map[string]spec.Schema{
					"name": {
						SchemaProps: spec.SchemaProps{
							Description: "Name for the Customer Gateway.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"region": {
						SchemaProps: spec.SchemaProps{
							Description: "AWS Region for the Customer Gateway.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"publicIP": {
						SchemaProps: spec.SchemaProps{
							Description: "Specify the Internet-routable IP address for your gateway's external interface; the address must be static and may be behind a device performing network address translation (NAT).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "The type of VPN connection that this customer gateway supports (ipsec.1).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"bgpAsn": {
						SchemaProps: spec.SchemaProps{
							Description: "The Border Gateway Protocol (BGP) Autonomous System Number (ASN) of your customer gateway. You can use an existing ASN assigned to your network. If you do not have one, you can use a private ASN in the 64512-65534 range. Default: 65000",
							Type:        []string{"integer"},
							Format:      "int64",
						},
					},
				},
				Required: []string{"name", "region", "publicIP", "type"},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_awsvpn_v1alpha1_CustomerGatewayStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CustomerGatewayStatus defines the observed state of CustomerGateway",
				Properties: map[string]spec.Schema{
					"customerGatewayId": {
						SchemaProps: spec.SchemaProps{
							Description: "The ID of the customer gateway.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"ipAddress": {
						SchemaProps: spec.SchemaProps{
							Description: "The Internet-routable IP address of the customer gateway's outside interface.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"bgpAsn": {
						SchemaProps: spec.SchemaProps{
							Description: "The customer gateway's Border Gateway Protocol (BGP) Autonomous System Number (ASN).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"state": {
						SchemaProps: spec.SchemaProps{
							Description: "The current state of the customer gateway.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "The type of VPN connection the customer gateway supports (ipsec.1).",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"tags": {
						SchemaProps: spec.SchemaProps{
							Description: "Tags assigned to the customer gateway.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/awsvpn/v1alpha1.Tag"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/awsvpn/v1alpha1.Tag"},
	}
}

func schema_pkg_apis_awsvpn_v1alpha1_VpnGateway(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VpnGateway is the Schema for the vpngateways API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/awsvpn/v1alpha1.VpnGatewaySpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/awsvpn/v1alpha1.VpnGatewayStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/awsvpn/v1alpha1.VpnGatewaySpec", "./pkg/apis/awsvpn/v1alpha1.VpnGatewayStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_awsvpn_v1alpha1_VpnGatewaySpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VpnGatewaySpec defines the desired state of VpnGateway",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_awsvpn_v1alpha1_VpnGatewayStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "VpnGatewayStatus defines the observed state of VpnGateway",
				Properties:  map[string]spec.Schema{},
			},
		},
		Dependencies: []string{},
	}
}
