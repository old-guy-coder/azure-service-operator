// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=network.azure.com,resources=privateendpointsprivatednszonegroups,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={privateendpointsprivatednszonegroups/status,privateendpointsprivatednszonegroups/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220701.PrivateEndpointsPrivateDnsZoneGroup
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2022-07-01/privateEndpoint.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}
type PrivateEndpointsPrivateDnsZoneGroup struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateEndpointsPrivateDnsZoneGroup_Spec   `json:"spec,omitempty"`
	Status            PrivateEndpointsPrivateDnsZoneGroup_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &PrivateEndpointsPrivateDnsZoneGroup{}

// GetConditions returns the conditions of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) GetConditions() conditions.Conditions {
	return group.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (group *PrivateEndpointsPrivateDnsZoneGroup) SetConditions(conditions conditions.Conditions) {
	group.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &PrivateEndpointsPrivateDnsZoneGroup{}

// AzureName returns the Azure name of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) AzureName() string {
	return group.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (group PrivateEndpointsPrivateDnsZoneGroup) GetAPIVersion() string {
	return "2022-07-01"
}

// GetResourceScope returns the scope of the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) GetSpec() genruntime.ConvertibleSpec {
	return &group.Spec
}

// GetStatus returns the status of this resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) GetStatus() genruntime.ConvertibleStatus {
	return &group.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/privateEndpoints/privateDnsZoneGroups"
func (group *PrivateEndpointsPrivateDnsZoneGroup) GetType() string {
	return "Microsoft.Network/privateEndpoints/privateDnsZoneGroups"
}

// NewEmptyStatus returns a new empty (blank) status
func (group *PrivateEndpointsPrivateDnsZoneGroup) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &PrivateEndpointsPrivateDnsZoneGroup_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (group *PrivateEndpointsPrivateDnsZoneGroup) Owner() *genruntime.ResourceReference {
	ownerGroup, ownerKind := genruntime.LookupOwnerGroupKind(group.Spec)
	return group.Spec.Owner.AsResourceReference(ownerGroup, ownerKind)
}

// SetStatus sets the status of this resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*PrivateEndpointsPrivateDnsZoneGroup_STATUS); ok {
		group.Status = *st
		return nil
	}

	// Convert status to required version
	var st PrivateEndpointsPrivateDnsZoneGroup_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	group.Status = st
	return nil
}

// Hub marks that this PrivateEndpointsPrivateDnsZoneGroup is the hub type for conversion
func (group *PrivateEndpointsPrivateDnsZoneGroup) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (group *PrivateEndpointsPrivateDnsZoneGroup) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: group.Spec.OriginalVersion,
		Kind:    "PrivateEndpointsPrivateDnsZoneGroup",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220701.PrivateEndpointsPrivateDnsZoneGroup
// Generator information:
// - Generated from: /network/resource-manager/Microsoft.Network/stable/2022-07-01/privateEndpoint.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateEndpoints/{privateEndpointName}/privateDnsZoneGroups/{privateDnsZoneGroupName}
type PrivateEndpointsPrivateDnsZoneGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PrivateEndpointsPrivateDnsZoneGroup `json:"items"`
}

// Storage version of v1api20220701.PrivateEndpointsPrivateDnsZoneGroup_Spec
type PrivateEndpointsPrivateDnsZoneGroup_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string `json:"azureName,omitempty"`
	OriginalVersion string `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/PrivateEndpoint resource
	Owner                 *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"PrivateEndpoint"`
	PrivateDnsZoneConfigs []PrivateDnsZoneConfig             `json:"privateDnsZoneConfigs,omitempty"`
	PropertyBag           genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
}

var _ genruntime.ConvertibleSpec = &PrivateEndpointsPrivateDnsZoneGroup_Spec{}

// ConvertSpecFrom populates our PrivateEndpointsPrivateDnsZoneGroup_Spec from the provided source
func (group *PrivateEndpointsPrivateDnsZoneGroup_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(group)
}

// ConvertSpecTo populates the provided destination from our PrivateEndpointsPrivateDnsZoneGroup_Spec
func (group *PrivateEndpointsPrivateDnsZoneGroup_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(group)
}

// Storage version of v1api20220701.PrivateEndpointsPrivateDnsZoneGroup_STATUS
type PrivateEndpointsPrivateDnsZoneGroup_STATUS struct {
	Conditions            []conditions.Condition        `json:"conditions,omitempty"`
	Etag                  *string                       `json:"etag,omitempty"`
	Id                    *string                       `json:"id,omitempty"`
	Name                  *string                       `json:"name,omitempty"`
	PrivateDnsZoneConfigs []PrivateDnsZoneConfig_STATUS `json:"privateDnsZoneConfigs,omitempty"`
	PropertyBag           genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	ProvisioningState     *string                       `json:"provisioningState,omitempty"`
}

var _ genruntime.ConvertibleStatus = &PrivateEndpointsPrivateDnsZoneGroup_STATUS{}

// ConvertStatusFrom populates our PrivateEndpointsPrivateDnsZoneGroup_STATUS from the provided source
func (group *PrivateEndpointsPrivateDnsZoneGroup_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(group)
}

// ConvertStatusTo populates the provided destination from our PrivateEndpointsPrivateDnsZoneGroup_STATUS
func (group *PrivateEndpointsPrivateDnsZoneGroup_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == group {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(group)
}

// Storage version of v1api20220701.PrivateDnsZoneConfig
// PrivateDnsZoneConfig resource.
type PrivateDnsZoneConfig struct {
	Name *string `json:"name,omitempty"`

	// PrivateDnsZoneReference: The resource id of the private dns zone.
	PrivateDnsZoneReference *genruntime.ResourceReference `armReference:"PrivateDnsZoneId" json:"privateDnsZoneReference,omitempty"`
	PropertyBag             genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220701.PrivateDnsZoneConfig_STATUS
// PrivateDnsZoneConfig resource.
type PrivateDnsZoneConfig_STATUS struct {
	Name             *string                `json:"name,omitempty"`
	PrivateDnsZoneId *string                `json:"privateDnsZoneId,omitempty"`
	PropertyBag      genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RecordSets       []RecordSet_STATUS     `json:"recordSets,omitempty"`
}

// Storage version of v1api20220701.RecordSet_STATUS
// A collective group of information about the record set information.
type RecordSet_STATUS struct {
	Fqdn              *string                `json:"fqdn,omitempty"`
	IpAddresses       []string               `json:"ipAddresses,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ProvisioningState *string                `json:"provisioningState,omitempty"`
	RecordSetName     *string                `json:"recordSetName,omitempty"`
	RecordType        *string                `json:"recordType,omitempty"`
	Ttl               *int                   `json:"ttl,omitempty"`
}

func init() {
	SchemeBuilder.Register(&PrivateEndpointsPrivateDnsZoneGroup{}, &PrivateEndpointsPrivateDnsZoneGroupList{})
}
