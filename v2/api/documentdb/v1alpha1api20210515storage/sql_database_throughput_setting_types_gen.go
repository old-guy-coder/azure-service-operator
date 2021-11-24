// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210515storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=documentdb.azure.com,resources=sqldatabasethroughputsettings,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=documentdb.azure.com,resources={sqldatabasethroughputsettings/status,sqldatabasethroughputsettings/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210515.SqlDatabaseThroughputSetting
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_throughputSettings
type SqlDatabaseThroughputSetting struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseAccountsSqlDatabasesThroughputSettings_Spec `json:"spec,omitempty"`
	Status            ThroughputSettingsGetResults_Status                 `json:"status,omitempty"`
}

var _ conditions.Conditioner = &SqlDatabaseThroughputSetting{}

// GetConditions returns the conditions of the resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetConditions() conditions.Conditions {
	return sqlDatabaseThroughputSetting.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) SetConditions(conditions conditions.Conditions) {
	sqlDatabaseThroughputSetting.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &SqlDatabaseThroughputSetting{}

// AzureName returns the Azure name of the resource (always "default")
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-05-15"
func (sqlDatabaseThroughputSetting SqlDatabaseThroughputSetting) GetAPIVersion() string {
	return "2021-05-15"
}

// GetResourceKind returns the kind of the resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetSpec() genruntime.ConvertibleSpec {
	return &sqlDatabaseThroughputSetting.Spec
}

// GetStatus returns the status of this resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetStatus() genruntime.ConvertibleStatus {
	return &sqlDatabaseThroughputSetting.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/throughputSettings"
}

// NewEmptyStatus returns a new empty (blank) status
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &ThroughputSettingsGetResults_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(sqlDatabaseThroughputSetting.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  sqlDatabaseThroughputSetting.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*ThroughputSettingsGetResults_Status); ok {
		sqlDatabaseThroughputSetting.Status = *st
		return nil
	}

	// Convert status to required version
	var st ThroughputSettingsGetResults_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	sqlDatabaseThroughputSetting.Status = st
	return nil
}

// Hub marks that this SqlDatabaseThroughputSetting is the hub type for conversion
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (sqlDatabaseThroughputSetting *SqlDatabaseThroughputSetting) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: sqlDatabaseThroughputSetting.Spec.OriginalVersion,
		Kind:    "SqlDatabaseThroughputSetting",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210515.SqlDatabaseThroughputSetting
//Generated from: https://schema.management.azure.com/schemas/2021-05-15/Microsoft.DocumentDB.json#/resourceDefinitions/databaseAccounts_sqlDatabases_throughputSettings
type SqlDatabaseThroughputSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SqlDatabaseThroughputSetting `json:"items"`
}

//Storage version of v1alpha1api20210515.DatabaseAccountsSqlDatabasesThroughputSettings_Spec
type DatabaseAccountsSqlDatabasesThroughputSettings_Spec struct {
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner       genruntime.KnownResourceReference `group:"documentdb.azure.com" json:"owner" kind:"SqlDatabase"`
	PropertyBag genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	Resource    *ThroughputSettingsResource       `json:"resource,omitempty"`
	Tags        map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DatabaseAccountsSqlDatabasesThroughputSettings_Spec{}

// ConvertSpecFrom populates our DatabaseAccountsSqlDatabasesThroughputSettings_Spec from the provided source
func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == databaseAccountsSqlDatabasesThroughputSettingsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(databaseAccountsSqlDatabasesThroughputSettingsSpec)
}

// ConvertSpecTo populates the provided destination from our DatabaseAccountsSqlDatabasesThroughputSettings_Spec
func (databaseAccountsSqlDatabasesThroughputSettingsSpec *DatabaseAccountsSqlDatabasesThroughputSettings_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == databaseAccountsSqlDatabasesThroughputSettingsSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(databaseAccountsSqlDatabasesThroughputSettingsSpec)
}

func init() {
	SchemeBuilder.Register(&SqlDatabaseThroughputSetting{}, &SqlDatabaseThroughputSettingList{})
}
