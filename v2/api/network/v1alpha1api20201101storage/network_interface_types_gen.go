// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201101storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=networkinterfaces,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={networkinterfaces/status,networkinterfaces/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201101.NetworkInterface
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkInterfaces
type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NetworkInterfaces_Spec                                       `json:"spec,omitempty"`
	Status            NetworkInterface_Status_NetworkInterface_SubResourceEmbedded `json:"status,omitempty"`
}

var _ conditions.Conditioner = &NetworkInterface{}

// GetConditions returns the conditions of the resource
func (networkInterface *NetworkInterface) GetConditions() conditions.Conditions {
	return networkInterface.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (networkInterface *NetworkInterface) SetConditions(conditions conditions.Conditions) {
	networkInterface.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &NetworkInterface{}

// AzureName returns the Azure name of the resource
func (networkInterface *NetworkInterface) AzureName() string {
	return networkInterface.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-11-01"
func (networkInterface NetworkInterface) GetAPIVersion() string {
	return "2020-11-01"
}

// GetResourceKind returns the kind of the resource
func (networkInterface *NetworkInterface) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (networkInterface *NetworkInterface) GetSpec() genruntime.ConvertibleSpec {
	return &networkInterface.Spec
}

// GetStatus returns the status of this resource
func (networkInterface *NetworkInterface) GetStatus() genruntime.ConvertibleStatus {
	return &networkInterface.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/networkInterfaces"
func (networkInterface *NetworkInterface) GetType() string {
	return "Microsoft.Network/networkInterfaces"
}

// NewEmptyStatus returns a new empty (blank) status
func (networkInterface *NetworkInterface) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &NetworkInterface_Status_NetworkInterface_SubResourceEmbedded{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (networkInterface *NetworkInterface) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(networkInterface.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  networkInterface.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (networkInterface *NetworkInterface) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*NetworkInterface_Status_NetworkInterface_SubResourceEmbedded); ok {
		networkInterface.Status = *st
		return nil
	}

	// Convert status to required version
	var st NetworkInterface_Status_NetworkInterface_SubResourceEmbedded
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	networkInterface.Status = st
	return nil
}

// Hub marks that this NetworkInterface is the hub type for conversion
func (networkInterface *NetworkInterface) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (networkInterface *NetworkInterface) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: networkInterface.Spec.OriginalVersion,
		Kind:    "NetworkInterface",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201101.NetworkInterface
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/resourceDefinitions/networkInterfaces
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterface `json:"items"`
}

//Storage version of v1alpha1api20201101.NetworkInterface_Status_NetworkInterface_SubResourceEmbedded
type NetworkInterface_Status_NetworkInterface_SubResourceEmbedded struct {
	Conditions                  []conditions.Condition                                                         `json:"conditions,omitempty"`
	DnsSettings                 *NetworkInterfaceDnsSettings_Status                                            `json:"dnsSettings,omitempty"`
	DscpConfiguration           *SubResource_Status                                                            `json:"dscpConfiguration,omitempty"`
	EnableAcceleratedNetworking *bool                                                                          `json:"enableAcceleratedNetworking,omitempty"`
	EnableIPForwarding          *bool                                                                          `json:"enableIPForwarding,omitempty"`
	Etag                        *string                                                                        `json:"etag,omitempty"`
	ExtendedLocation            *ExtendedLocation_Status                                                       `json:"extendedLocation,omitempty"`
	HostedWorkloads             []string                                                                       `json:"hostedWorkloads,omitempty"`
	Id                          *string                                                                        `json:"id,omitempty"`
	IpConfigurations            []NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded  `json:"ipConfigurations,omitempty"`
	Location                    *string                                                                        `json:"location,omitempty"`
	MacAddress                  *string                                                                        `json:"macAddress,omitempty"`
	MigrationPhase              *string                                                                        `json:"migrationPhase,omitempty"`
	Name                        *string                                                                        `json:"name,omitempty"`
	NetworkSecurityGroup        *NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbedded              `json:"networkSecurityGroup,omitempty"`
	NicType                     *string                                                                        `json:"nicType,omitempty"`
	Primary                     *bool                                                                          `json:"primary,omitempty"`
	PrivateEndpoint             *PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded                   `json:"privateEndpoint,omitempty"`
	PrivateLinkService          *PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded                `json:"privateLinkService,omitempty"`
	PropertyBag                 genruntime.PropertyBag                                                         `json:"$propertyBag,omitempty"`
	ProvisioningState           *string                                                                        `json:"provisioningState,omitempty"`
	ResourceGuid                *string                                                                        `json:"resourceGuid,omitempty"`
	Tags                        map[string]string                                                              `json:"tags,omitempty"`
	TapConfigurations           []NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbedded `json:"tapConfigurations,omitempty"`
	Type                        *string                                                                        `json:"type,omitempty"`
	VirtualMachine              *SubResource_Status                                                            `json:"virtualMachine,omitempty"`
}

var _ genruntime.ConvertibleStatus = &NetworkInterface_Status_NetworkInterface_SubResourceEmbedded{}

// ConvertStatusFrom populates our NetworkInterface_Status_NetworkInterface_SubResourceEmbedded from the provided source
func (networkInterfaceStatusNetworkInterfaceSubResourceEmbedded *NetworkInterface_Status_NetworkInterface_SubResourceEmbedded) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == networkInterfaceStatusNetworkInterfaceSubResourceEmbedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(networkInterfaceStatusNetworkInterfaceSubResourceEmbedded)
}

// ConvertStatusTo populates the provided destination from our NetworkInterface_Status_NetworkInterface_SubResourceEmbedded
func (networkInterfaceStatusNetworkInterfaceSubResourceEmbedded *NetworkInterface_Status_NetworkInterface_SubResourceEmbedded) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == networkInterfaceStatusNetworkInterfaceSubResourceEmbedded {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(networkInterfaceStatusNetworkInterfaceSubResourceEmbedded)
}

//Storage version of v1alpha1api20201101.NetworkInterfaces_Spec
type NetworkInterfaces_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName                   string                                               `json:"azureName"`
	DnsSettings                 *NetworkInterfaceDnsSettings                         `json:"dnsSettings,omitempty"`
	EnableAcceleratedNetworking *bool                                                `json:"enableAcceleratedNetworking,omitempty"`
	EnableIPForwarding          *bool                                                `json:"enableIPForwarding,omitempty"`
	ExtendedLocation            *ExtendedLocation                                    `json:"extendedLocation,omitempty"`
	IpConfigurations            []NetworkInterfaces_Spec_Properties_IpConfigurations `json:"ipConfigurations,omitempty"`
	Location                    *string                                              `json:"location,omitempty"`
	NetworkSecurityGroup        *SubResource                                         `json:"networkSecurityGroup,omitempty"`
	OriginalVersion             string                                               `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner" kind:"ResourceGroup"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &NetworkInterfaces_Spec{}

// ConvertSpecFrom populates our NetworkInterfaces_Spec from the provided source
func (networkInterfacesSpec *NetworkInterfaces_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == networkInterfacesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(networkInterfacesSpec)
}

// ConvertSpecTo populates the provided destination from our NetworkInterfaces_Spec
func (networkInterfacesSpec *NetworkInterfaces_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == networkInterfacesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(networkInterfacesSpec)
}

//Storage version of v1alpha1api20201101.NetworkInterfaceDnsSettings
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/NetworkInterfaceDnsSettings
type NetworkInterfaceDnsSettings struct {
	DnsServers           []string               `json:"dnsServers,omitempty"`
	InternalDnsNameLabel *string                `json:"internalDnsNameLabel,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkInterfaceDnsSettings_Status
type NetworkInterfaceDnsSettings_Status struct {
	AppliedDnsServers        []string               `json:"appliedDnsServers,omitempty"`
	DnsServers               []string               `json:"dnsServers,omitempty"`
	InternalDnsNameLabel     *string                `json:"internalDnsNameLabel,omitempty"`
	InternalDomainNameSuffix *string                `json:"internalDomainNameSuffix,omitempty"`
	InternalFqdn             *string                `json:"internalFqdn,omitempty"`
	PropertyBag              genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded
type NetworkInterfaceIPConfiguration_Status_NetworkInterface_SubResourceEmbedded struct {
	ApplicationGatewayBackendAddressPools []ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded `json:"applicationGatewayBackendAddressPools,omitempty"`
	ApplicationSecurityGroups             []ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbedded             `json:"applicationSecurityGroups,omitempty"`
	Etag                                  *string                                                                            `json:"etag,omitempty"`
	Id                                    *string                                                                            `json:"id,omitempty"`
	LoadBalancerBackendAddressPools       []BackendAddressPool_Status_NetworkInterface_SubResourceEmbedded                   `json:"loadBalancerBackendAddressPools,omitempty"`
	LoadBalancerInboundNatRules           []InboundNatRule_Status_NetworkInterface_SubResourceEmbedded                       `json:"loadBalancerInboundNatRules,omitempty"`
	Name                                  *string                                                                            `json:"name,omitempty"`
	Primary                               *bool                                                                              `json:"primary,omitempty"`
	PrivateIPAddress                      *string                                                                            `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion               *string                                                                            `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod             *string                                                                            `json:"privateIPAllocationMethod,omitempty"`
	PrivateLinkConnectionProperties       *NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status             `json:"privateLinkConnectionProperties,omitempty"`
	PropertyBag                           genruntime.PropertyBag                                                             `json:"$propertyBag,omitempty"`
	ProvisioningState                     *string                                                                            `json:"provisioningState,omitempty"`
	PublicIPAddress                       *PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded                       `json:"publicIPAddress,omitempty"`
	Subnet                                *Subnet_Status_NetworkInterface_SubResourceEmbedded                                `json:"subnet,omitempty"`
	Type                                  *string                                                                            `json:"type,omitempty"`
	VirtualNetworkTaps                    []VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbedded                    `json:"virtualNetworkTaps,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbedded
type NetworkInterfaceTapConfiguration_Status_NetworkInterface_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkInterfaces_Spec_Properties_IpConfigurations
type NetworkInterfaces_Spec_Properties_IpConfigurations struct {
	ApplicationGatewayBackendAddressPools []SubResource          `json:"applicationGatewayBackendAddressPools,omitempty"`
	ApplicationSecurityGroups             []SubResource          `json:"applicationSecurityGroups,omitempty"`
	LoadBalancerBackendAddressPools       []SubResource          `json:"loadBalancerBackendAddressPools,omitempty"`
	LoadBalancerInboundNatRules           []SubResource          `json:"loadBalancerInboundNatRules,omitempty"`
	Name                                  *string                `json:"name,omitempty"`
	Primary                               *bool                  `json:"primary,omitempty"`
	PrivateIPAddress                      *string                `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion               *string                `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod             *string                `json:"privateIPAllocationMethod,omitempty"`
	PropertyBag                           genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	PublicIPAddress                       *SubResource           `json:"publicIPAddress,omitempty"`
	Subnet                                *SubResource           `json:"subnet,omitempty"`
	VirtualNetworkTaps                    []SubResource          `json:"virtualNetworkTaps,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbedded
type NetworkSecurityGroup_Status_NetworkInterface_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded
type PrivateEndpoint_Status_NetworkInterface_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status `json:"extendedLocation,omitempty"`
	Id               *string                  `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded
type PrivateLinkService_Status_NetworkInterface_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status `json:"extendedLocation,omitempty"`
	Id               *string                  `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag   `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.SubResource
//Generated from: https://schema.management.azure.com/schemas/2020-11-01/Microsoft.Network.json#/definitions/SubResource
type SubResource struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// +kubebuilder:validation:Required
	//Reference: Resource ID.
	Reference genruntime.ResourceReference `armReference:"Id" json:"reference"`
}

//Storage version of v1alpha1api20201101.SubResource_Status
type SubResource_Status struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded
type ApplicationGatewayBackendAddressPool_Status_NetworkInterface_SubResourceEmbedded struct {
	BackendAddresses  []ApplicationGatewayBackendAddress_Status `json:"backendAddresses,omitempty"`
	Etag              *string                                   `json:"etag,omitempty"`
	Id                *string                                   `json:"id,omitempty"`
	Name              *string                                   `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag                    `json:"$propertyBag,omitempty"`
	ProvisioningState *string                                   `json:"provisioningState,omitempty"`
	Type              *string                                   `json:"type,omitempty"`
}

//Storage version of v1alpha1api20201101.ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbedded
type ApplicationSecurityGroup_Status_NetworkInterface_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.BackendAddressPool_Status_NetworkInterface_SubResourceEmbedded
type BackendAddressPool_Status_NetworkInterface_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.InboundNatRule_Status_NetworkInterface_SubResourceEmbedded
type InboundNatRule_Status_NetworkInterface_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status
type NetworkInterfaceIPConfigurationPrivateLinkConnectionProperties_Status struct {
	Fqdns              []string               `json:"fqdns,omitempty"`
	GroupId            *string                `json:"groupId,omitempty"`
	PropertyBag        genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RequiredMemberName *string                `json:"requiredMemberName,omitempty"`
}

//Storage version of v1alpha1api20201101.PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded
type PublicIPAddress_Status_NetworkInterface_SubResourceEmbedded struct {
	ExtendedLocation *ExtendedLocation_Status   `json:"extendedLocation,omitempty"`
	Id               *string                    `json:"id,omitempty"`
	PropertyBag      genruntime.PropertyBag     `json:"$propertyBag,omitempty"`
	Sku              *PublicIPAddressSku_Status `json:"sku,omitempty"`
	Zones            []string                   `json:"zones,omitempty"`
}

//Storage version of v1alpha1api20201101.Subnet_Status_NetworkInterface_SubResourceEmbedded
type Subnet_Status_NetworkInterface_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbedded
type VirtualNetworkTap_Status_NetworkInterface_SubResourceEmbedded struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20201101.ApplicationGatewayBackendAddress_Status
type ApplicationGatewayBackendAddress_Status struct {
	Fqdn        *string                `json:"fqdn,omitempty"`
	IpAddress   *string                `json:"ipAddress,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&NetworkInterface{}, &NetworkInterfaceList{})
}
