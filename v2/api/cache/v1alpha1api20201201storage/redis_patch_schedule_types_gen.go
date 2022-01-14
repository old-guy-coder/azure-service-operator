// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20201201storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=cache.azure.com,resources=redispatchschedules,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=cache.azure.com,resources={redispatchschedules/status,redispatchschedules/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20201201.RedisPatchSchedule
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_patchSchedules
type RedisPatchSchedule struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisPatchSchedules_Spec  `json:"spec,omitempty"`
	Status            RedisPatchSchedule_Status `json:"status,omitempty"`
}

var _ conditions.Conditioner = &RedisPatchSchedule{}

// GetConditions returns the conditions of the resource
func (redisPatchSchedule *RedisPatchSchedule) GetConditions() conditions.Conditions {
	return redisPatchSchedule.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (redisPatchSchedule *RedisPatchSchedule) SetConditions(conditions conditions.Conditions) {
	redisPatchSchedule.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &RedisPatchSchedule{}

// AzureName returns the Azure name of the resource (always "default")
func (redisPatchSchedule *RedisPatchSchedule) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2020-12-01"
func (redisPatchSchedule RedisPatchSchedule) GetAPIVersion() string {
	return "2020-12-01"
}

// GetResourceKind returns the kind of the resource
func (redisPatchSchedule *RedisPatchSchedule) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (redisPatchSchedule *RedisPatchSchedule) GetSpec() genruntime.ConvertibleSpec {
	return &redisPatchSchedule.Spec
}

// GetStatus returns the status of this resource
func (redisPatchSchedule *RedisPatchSchedule) GetStatus() genruntime.ConvertibleStatus {
	return &redisPatchSchedule.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/patchSchedules"
func (redisPatchSchedule *RedisPatchSchedule) GetType() string {
	return "Microsoft.Cache/redis/patchSchedules"
}

// NewEmptyStatus returns a new empty (blank) status
func (redisPatchSchedule *RedisPatchSchedule) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &RedisPatchSchedule_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (redisPatchSchedule *RedisPatchSchedule) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(redisPatchSchedule.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  redisPatchSchedule.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (redisPatchSchedule *RedisPatchSchedule) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*RedisPatchSchedule_Status); ok {
		redisPatchSchedule.Status = *st
		return nil
	}

	// Convert status to required version
	var st RedisPatchSchedule_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	redisPatchSchedule.Status = st
	return nil
}

// Hub marks that this RedisPatchSchedule is the hub type for conversion
func (redisPatchSchedule *RedisPatchSchedule) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (redisPatchSchedule *RedisPatchSchedule) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: redisPatchSchedule.Spec.OriginalVersion,
		Kind:    "RedisPatchSchedule",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20201201.RedisPatchSchedule
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/resourceDefinitions/redis_patchSchedules
type RedisPatchScheduleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RedisPatchSchedule `json:"items"`
}

//Storage version of v1alpha1api20201201.RedisPatchSchedule_Status
type RedisPatchSchedule_Status struct {
	Conditions      []conditions.Condition `json:"conditions,omitempty"`
	Id              *string                `json:"id,omitempty"`
	Name            *string                `json:"name,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry_Status `json:"scheduleEntries,omitempty"`
	Type            *string                `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &RedisPatchSchedule_Status{}

// ConvertStatusFrom populates our RedisPatchSchedule_Status from the provided source
func (redisPatchScheduleStatus *RedisPatchSchedule_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == redisPatchScheduleStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(redisPatchScheduleStatus)
}

// ConvertStatusTo populates the provided destination from our RedisPatchSchedule_Status
func (redisPatchScheduleStatus *RedisPatchSchedule_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == redisPatchScheduleStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(redisPatchScheduleStatus)
}

//Storage version of v1alpha1api20201201.RedisPatchSchedules_Spec
type RedisPatchSchedules_Spec struct {
	Location        *string `json:"location,omitempty"`
	OriginalVersion string  `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner           genruntime.KnownResourceReference `group:"cache.azure.com" json:"owner" kind:"Redis"`
	PropertyBag     genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	ScheduleEntries []ScheduleEntry                   `json:"scheduleEntries,omitempty"`
	Tags            map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &RedisPatchSchedules_Spec{}

// ConvertSpecFrom populates our RedisPatchSchedules_Spec from the provided source
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == redisPatchSchedulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(redisPatchSchedulesSpec)
}

// ConvertSpecTo populates the provided destination from our RedisPatchSchedules_Spec
func (redisPatchSchedulesSpec *RedisPatchSchedules_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == redisPatchSchedulesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(redisPatchSchedulesSpec)
}

//Storage version of v1alpha1api20201201.ScheduleEntry
//Generated from: https://schema.management.azure.com/schemas/2020-12-01/Microsoft.Cache.json#/definitions/ScheduleEntry
type ScheduleEntry struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

//Storage version of v1alpha1api20201201.ScheduleEntry_Status
type ScheduleEntry_Status struct {
	DayOfWeek         *string                `json:"dayOfWeek,omitempty"`
	MaintenanceWindow *string                `json:"maintenanceWindow,omitempty"`
	PropertyBag       genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	StartHourUtc      *int                   `json:"startHourUtc,omitempty"`
}

func init() {
	SchemeBuilder.Register(&RedisPatchSchedule{}, &RedisPatchScheduleList{})
}