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

// +kubebuilder:rbac:groups=network.azure.com,resources=dnsresolversoutboundendpoints,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=network.azure.com,resources={dnsresolversoutboundendpoints/status,dnsresolversoutboundendpoints/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220701.DnsResolversOutboundEndpoint
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints/{outboundEndpointName}
type DnsResolversOutboundEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DnsResolversOutboundEndpoint_Spec   `json:"spec,omitempty"`
	Status            DnsResolversOutboundEndpoint_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DnsResolversOutboundEndpoint{}

// GetConditions returns the conditions of the resource
func (endpoint *DnsResolversOutboundEndpoint) GetConditions() conditions.Conditions {
	return endpoint.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (endpoint *DnsResolversOutboundEndpoint) SetConditions(conditions conditions.Conditions) {
	endpoint.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DnsResolversOutboundEndpoint{}

// AzureName returns the Azure name of the resource
func (endpoint *DnsResolversOutboundEndpoint) AzureName() string {
	return endpoint.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-01"
func (endpoint DnsResolversOutboundEndpoint) GetAPIVersion() string {
	return "2022-07-01"
}

// GetResourceScope returns the scope of the resource
func (endpoint *DnsResolversOutboundEndpoint) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (endpoint *DnsResolversOutboundEndpoint) GetSpec() genruntime.ConvertibleSpec {
	return &endpoint.Spec
}

// GetStatus returns the status of this resource
func (endpoint *DnsResolversOutboundEndpoint) GetStatus() genruntime.ConvertibleStatus {
	return &endpoint.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (endpoint *DnsResolversOutboundEndpoint) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationDelete,
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Network/dnsResolvers/outboundEndpoints"
func (endpoint *DnsResolversOutboundEndpoint) GetType() string {
	return "Microsoft.Network/dnsResolvers/outboundEndpoints"
}

// NewEmptyStatus returns a new empty (blank) status
func (endpoint *DnsResolversOutboundEndpoint) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DnsResolversOutboundEndpoint_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (endpoint *DnsResolversOutboundEndpoint) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(endpoint.Spec)
	return endpoint.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (endpoint *DnsResolversOutboundEndpoint) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DnsResolversOutboundEndpoint_STATUS); ok {
		endpoint.Status = *st
		return nil
	}

	// Convert status to required version
	var st DnsResolversOutboundEndpoint_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	endpoint.Status = st
	return nil
}

// Hub marks that this DnsResolversOutboundEndpoint is the hub type for conversion
func (endpoint *DnsResolversOutboundEndpoint) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (endpoint *DnsResolversOutboundEndpoint) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: endpoint.Spec.OriginalVersion,
		Kind:    "DnsResolversOutboundEndpoint",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220701.DnsResolversOutboundEndpoint
// Generator information:
// - Generated from: /dnsresolver/resource-manager/Microsoft.Network/stable/2022-07-01/dnsresolver.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/dnsResolvers/{dnsResolverName}/outboundEndpoints/{outboundEndpointName}
type DnsResolversOutboundEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DnsResolversOutboundEndpoint `json:"items"`
}

// Storage version of v1api20220701.DnsResolversOutboundEndpoint_Spec
type DnsResolversOutboundEndpoint_Spec struct {
	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName       string  `json:"azureName,omitempty"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a network.azure.com/DnsResolver resource
	Owner       *genruntime.KnownResourceReference `group:"network.azure.com" json:"owner,omitempty" kind:"DnsResolver"`
	PropertyBag genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	Subnet      *DnsresolverSubResource            `json:"subnet,omitempty"`
	Tags        map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DnsResolversOutboundEndpoint_Spec{}

// ConvertSpecFrom populates our DnsResolversOutboundEndpoint_Spec from the provided source
func (endpoint *DnsResolversOutboundEndpoint_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(endpoint)
}

// ConvertSpecTo populates the provided destination from our DnsResolversOutboundEndpoint_Spec
func (endpoint *DnsResolversOutboundEndpoint_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(endpoint)
}

// Storage version of v1api20220701.DnsResolversOutboundEndpoint_STATUS
type DnsResolversOutboundEndpoint_STATUS struct {
	Conditions        []conditions.Condition         `json:"conditions,omitempty"`
	Etag              *string                        `json:"etag,omitempty"`
	Id                *string                        `json:"id,omitempty"`
	Location          *string                        `json:"location,omitempty"`
	Name              *string                        `json:"name,omitempty"`
	PropertyBag       genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	ProvisioningState *string                        `json:"provisioningState,omitempty"`
	ResourceGuid      *string                        `json:"resourceGuid,omitempty"`
	Subnet            *DnsresolverSubResource_STATUS `json:"subnet,omitempty"`
	SystemData        *SystemData_STATUS             `json:"systemData,omitempty"`
	Tags              map[string]string              `json:"tags,omitempty"`
	Type              *string                        `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DnsResolversOutboundEndpoint_STATUS{}

// ConvertStatusFrom populates our DnsResolversOutboundEndpoint_STATUS from the provided source
func (endpoint *DnsResolversOutboundEndpoint_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(endpoint)
}

// ConvertStatusTo populates the provided destination from our DnsResolversOutboundEndpoint_STATUS
func (endpoint *DnsResolversOutboundEndpoint_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == endpoint {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(endpoint)
}

func init() {
	SchemeBuilder.Register(&DnsResolversOutboundEndpoint{}, &DnsResolversOutboundEndpointList{})
}
