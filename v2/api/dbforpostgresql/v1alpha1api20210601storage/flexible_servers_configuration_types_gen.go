// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210601storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources=flexibleserversconfigurations,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=dbforpostgresql.azure.com,resources={flexibleserversconfigurations/status,flexibleserversconfigurations/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210601.FlexibleServersConfiguration
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_configurations
type FlexibleServersConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FlexibleServersConfigurations_Spec `json:"spec,omitempty"`
	Status            Configuration_Status               `json:"status,omitempty"`
}

var _ conditions.Conditioner = &FlexibleServersConfiguration{}

// GetConditions returns the conditions of the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetConditions() conditions.Conditions {
	return flexibleServersConfiguration.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (flexibleServersConfiguration *FlexibleServersConfiguration) SetConditions(conditions conditions.Conditions) {
	flexibleServersConfiguration.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &FlexibleServersConfiguration{}

// AzureName returns the Azure name of the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) AzureName() string {
	return flexibleServersConfiguration.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-06-01"
func (flexibleServersConfiguration FlexibleServersConfiguration) GetAPIVersion() string {
	return "2021-06-01"
}

// GetResourceKind returns the kind of the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetSpec() genruntime.ConvertibleSpec {
	return &flexibleServersConfiguration.Spec
}

// GetStatus returns the status of this resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetStatus() genruntime.ConvertibleStatus {
	return &flexibleServersConfiguration.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
func (flexibleServersConfiguration *FlexibleServersConfiguration) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/configurations"
}

// NewEmptyStatus returns a new empty (blank) status
func (flexibleServersConfiguration *FlexibleServersConfiguration) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &Configuration_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (flexibleServersConfiguration *FlexibleServersConfiguration) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(flexibleServersConfiguration.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  flexibleServersConfiguration.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*Configuration_Status); ok {
		flexibleServersConfiguration.Status = *st
		return nil
	}

	// Convert status to required version
	var st Configuration_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	flexibleServersConfiguration.Status = st
	return nil
}

// Hub marks that this FlexibleServersConfiguration is the hub type for conversion
func (flexibleServersConfiguration *FlexibleServersConfiguration) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (flexibleServersConfiguration *FlexibleServersConfiguration) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: flexibleServersConfiguration.Spec.OriginalVersion,
		Kind:    "FlexibleServersConfiguration",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210601.FlexibleServersConfiguration
//Generated from: https://schema.management.azure.com/schemas/2021-06-01/Microsoft.DBforPostgreSQL.json#/resourceDefinitions/flexibleServers_configurations
type FlexibleServersConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlexibleServersConfiguration `json:"items"`
}

//Storage version of v1alpha1api20210601.Configuration_Status
type Configuration_Status struct {
	AllowedValues *string                `json:"allowedValues,omitempty"`
	Conditions    []conditions.Condition `json:"conditions,omitempty"`
	DataType      *string                `json:"dataType,omitempty"`
	DefaultValue  *string                `json:"defaultValue,omitempty"`
	Description   *string                `json:"description,omitempty"`
	Id            *string                `json:"id,omitempty"`
	Name          *string                `json:"name,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Source        *string                `json:"source,omitempty"`
	SystemData    *SystemData_Status     `json:"systemData,omitempty"`
	Type          *string                `json:"type,omitempty"`
	Value         *string                `json:"value,omitempty"`
}

var _ genruntime.ConvertibleStatus = &Configuration_Status{}

// ConvertStatusFrom populates our Configuration_Status from the provided source
func (configurationStatus *Configuration_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == configurationStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(configurationStatus)
}

// ConvertStatusTo populates the provided destination from our Configuration_Status
func (configurationStatus *Configuration_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == configurationStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(configurationStatus)
}

//Storage version of v1alpha1api20210601.FlexibleServersConfigurations_Spec
type FlexibleServersConfigurations_Spec struct {
	//AzureName: The name of the resource in Azure. This is often the same as the name
	//of the resource in Kubernetes but it doesn't have to be.
	AzureName       string  `json:"azureName"`
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"dbforpostgresql.azure.com" json:"owner" kind:"FlexibleServer"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Source      *string                           `json:"source,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
	Value       *string                           `json:"value,omitempty"`
}

var _ genruntime.ConvertibleSpec = &FlexibleServersConfigurations_Spec{}

// ConvertSpecFrom populates our FlexibleServersConfigurations_Spec from the provided source
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == flexibleServersConfigurationsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(flexibleServersConfigurationsSpec)
}

// ConvertSpecTo populates the provided destination from our FlexibleServersConfigurations_Spec
func (flexibleServersConfigurationsSpec *FlexibleServersConfigurations_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == flexibleServersConfigurationsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(flexibleServersConfigurationsSpec)
}

func init() {
	SchemeBuilder.Register(&FlexibleServersConfiguration{}, &FlexibleServersConfigurationList{})
}
