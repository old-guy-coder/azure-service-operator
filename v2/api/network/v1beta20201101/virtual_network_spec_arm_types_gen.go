// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1beta20201101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type VirtualNetwork_Spec_ARM struct {
	// ExtendedLocation: The extended location of the virtual network.
	ExtendedLocation *ExtendedLocation_ARM `json:"extendedLocation,omitempty"`

	// Location: Location to deploy resource to
	Location *string `json:"location,omitempty"`

	// Name: Name of the resource
	Name string `json:"name,omitempty"`

	// Properties: Properties of the virtual network.
	Properties *VirtualNetwork_Properties_Spec_ARM `json:"properties,omitempty"`

	// Tags: Name-value pairs to add to the resource
	Tags map[string]string `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &VirtualNetwork_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (network VirtualNetwork_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (network *VirtualNetwork_Spec_ARM) GetName() string {
	return network.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/virtualNetworks"
func (network *VirtualNetwork_Spec_ARM) GetType() string {
	return "Microsoft.Network/virtualNetworks"
}

type VirtualNetwork_Properties_Spec_ARM struct {
	// AddressSpace: The AddressSpace that contains an array of IP address ranges that can be used by subnets.
	AddressSpace *AddressSpace_ARM `json:"addressSpace,omitempty"`

	// BgpCommunities: Bgp Communities sent over ExpressRoute with each route corresponding to a prefix in this VNET.
	BgpCommunities *VirtualNetworkBgpCommunities_ARM `json:"bgpCommunities,omitempty"`

	// DdosProtectionPlan: The DDoS protection plan associated with the virtual network.
	DdosProtectionPlan *SubResource_ARM `json:"ddosProtectionPlan,omitempty"`

	// DhcpOptions: The dhcpOptions that contains an array of DNS servers available to VMs deployed in the virtual network.
	DhcpOptions *DhcpOptions_ARM `json:"dhcpOptions,omitempty"`

	// EnableDdosProtection: Indicates if DDoS protection is enabled for all the protected resources in the virtual network. It
	// requires a DDoS protection plan associated with the resource.
	EnableDdosProtection *bool `json:"enableDdosProtection,omitempty"`

	// EnableVmProtection: Indicates if VM protection is enabled for all the subnets in the virtual network.
	EnableVmProtection *bool `json:"enableVmProtection,omitempty"`

	// IpAllocations: Array of IpAllocation which reference this VNET.
	IpAllocations []SubResource_ARM `json:"ipAllocations,omitempty"`

	// Subnets: A list of subnets in a Virtual Network.
	Subnets []VirtualNetwork_Properties_Subnets_Spec_ARM `json:"subnets,omitempty"`

	// VirtualNetworkPeerings: A list of peerings in a Virtual Network.
	VirtualNetworkPeerings []VirtualNetwork_Properties_VirtualNetworkPeerings_Spec_ARM `json:"virtualNetworkPeerings,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/DhcpOptions
type DhcpOptions_ARM struct {
	// DnsServers: The list of DNS servers IP addresses.
	DnsServers []string `json:"dnsServers,omitempty"`
}

type VirtualNetwork_Properties_Subnets_Spec_ARM struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the subnet.
	Properties *VirtualNetwork_Properties_Subnets_Properties_Spec_ARM `json:"properties,omitempty"`
}

type VirtualNetwork_Properties_VirtualNetworkPeerings_Spec_ARM struct {
	// Name: The name of the resource that is unique within a resource group. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the virtual network peering.
	Properties *VirtualNetworkPeeringPropertiesFormat_ARM `json:"properties,omitempty"`
}

// Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/VirtualNetworkBgpCommunities
type VirtualNetworkBgpCommunities_ARM struct {
	// VirtualNetworkCommunity: The BGP community associated with the virtual network.
	VirtualNetworkCommunity *string `json:"virtualNetworkCommunity,omitempty"`
}

type VirtualNetwork_Properties_Subnets_Properties_Spec_ARM struct {
	// AddressPrefix: The address prefix for the subnet.
	AddressPrefix *string `json:"addressPrefix,omitempty"`

	// AddressPrefixes: List of address prefixes for the subnet.
	AddressPrefixes []string `json:"addressPrefixes,omitempty"`

	// Delegations: An array of references to the delegations on the subnet.
	Delegations []VirtualNetwork_Properties_Subnets_Properties_Delegations_Spec_ARM `json:"delegations,omitempty"`

	// IpAllocations: Array of IpAllocation which reference this subnet.
	IpAllocations []SubResource_ARM `json:"ipAllocations,omitempty"`

	// NatGateway: Nat gateway associated with this subnet.
	NatGateway *SubResource_ARM `json:"natGateway,omitempty"`

	// NetworkSecurityGroup: The reference to the NetworkSecurityGroup resource.
	NetworkSecurityGroup *SubResource_ARM `json:"networkSecurityGroup,omitempty"`

	// PrivateEndpointNetworkPolicies: Enable or Disable apply network policies on private end point in the subnet.
	PrivateEndpointNetworkPolicies *string `json:"privateEndpointNetworkPolicies,omitempty"`

	// PrivateLinkServiceNetworkPolicies: Enable or Disable apply network policies on private link service in the subnet.
	PrivateLinkServiceNetworkPolicies *string `json:"privateLinkServiceNetworkPolicies,omitempty"`

	// RouteTable: The reference to the RouteTable resource.
	RouteTable *SubResource_ARM `json:"routeTable,omitempty"`

	// ServiceEndpointPolicies: An array of service endpoint policies.
	ServiceEndpointPolicies []SubResource_ARM `json:"serviceEndpointPolicies,omitempty"`

	// ServiceEndpoints: An array of service endpoints.
	ServiceEndpoints []ServiceEndpointPropertiesFormat_ARM `json:"serviceEndpoints,omitempty"`
}

type VirtualNetwork_Properties_Subnets_Properties_Delegations_Spec_ARM struct {
	// Name: The name of the resource that is unique within a subnet. This name can be used to access the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the subnet.
	Properties *ServiceDelegationPropertiesFormat_ARM `json:"properties,omitempty"`
}