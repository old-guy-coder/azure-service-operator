// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1alpha1api20210401storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountsblobservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountsblobservices/status,storageaccountsblobservices/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
//Storage version of v1alpha1api20210401.StorageAccountsBlobService
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts_blobServices
type StorageAccountsBlobService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsBlobServices_Spec `json:"spec,omitempty"`
	Status            BlobServiceProperties_Status     `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsBlobService{}

// GetConditions returns the conditions of the resource
func (storageAccountsBlobService *StorageAccountsBlobService) GetConditions() conditions.Conditions {
	return storageAccountsBlobService.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (storageAccountsBlobService *StorageAccountsBlobService) SetConditions(conditions conditions.Conditions) {
	storageAccountsBlobService.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &StorageAccountsBlobService{}

// AzureName returns the Azure name of the resource (always "default")
func (storageAccountsBlobService *StorageAccountsBlobService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (storageAccountsBlobService StorageAccountsBlobService) GetAPIVersion() string {
	return "2021-04-01"
}

// GetResourceKind returns the kind of the resource
func (storageAccountsBlobService *StorageAccountsBlobService) GetResourceKind() genruntime.ResourceKind {
	return genruntime.ResourceKindNormal
}

// GetSpec returns the specification of this resource
func (storageAccountsBlobService *StorageAccountsBlobService) GetSpec() genruntime.ConvertibleSpec {
	return &storageAccountsBlobService.Spec
}

// GetStatus returns the status of this resource
func (storageAccountsBlobService *StorageAccountsBlobService) GetStatus() genruntime.ConvertibleStatus {
	return &storageAccountsBlobService.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices"
func (storageAccountsBlobService *StorageAccountsBlobService) GetType() string {
	return "Microsoft.Storage/storageAccounts/blobServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (storageAccountsBlobService *StorageAccountsBlobService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &BlobServiceProperties_Status{}
}

// Owner returns the ResourceReference of the owner, or nil if there is no owner
func (storageAccountsBlobService *StorageAccountsBlobService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(storageAccountsBlobService.Spec)
	return &genruntime.ResourceReference{
		Group: group,
		Kind:  kind,
		Name:  storageAccountsBlobService.Spec.Owner.Name,
	}
}

// SetStatus sets the status of this resource
func (storageAccountsBlobService *StorageAccountsBlobService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*BlobServiceProperties_Status); ok {
		storageAccountsBlobService.Status = *st
		return nil
	}

	// Convert status to required version
	var st BlobServiceProperties_Status
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	storageAccountsBlobService.Status = st
	return nil
}

// Hub marks that this StorageAccountsBlobService is the hub type for conversion
func (storageAccountsBlobService *StorageAccountsBlobService) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (storageAccountsBlobService *StorageAccountsBlobService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: storageAccountsBlobService.Spec.OriginalVersion,
		Kind:    "StorageAccountsBlobService",
	}
}

// +kubebuilder:object:root=true
//Storage version of v1alpha1api20210401.StorageAccountsBlobService
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/resourceDefinitions/storageAccounts_blobServices
type StorageAccountsBlobServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsBlobService `json:"items"`
}

//Storage version of v1alpha1api20210401.BlobServiceProperties_Status
type BlobServiceProperties_Status struct {
	AutomaticSnapshotPolicyEnabled *bool                                `json:"automaticSnapshotPolicyEnabled,omitempty"`
	ChangeFeed                     *ChangeFeed_Status                   `json:"changeFeed,omitempty"`
	Conditions                     []conditions.Condition               `json:"conditions,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy_Status        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRules_Status                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                              `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy_Status        `json:"deleteRetentionPolicy,omitempty"`
	Id                             *string                              `json:"id,omitempty"`
	IsVersioningEnabled            *bool                                `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy_Status `json:"lastAccessTimeTrackingPolicy,omitempty"`
	Name                           *string                              `json:"name,omitempty"`
	PropertyBag                    genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	RestorePolicy                  *RestorePolicyProperties_Status      `json:"restorePolicy,omitempty"`
	Sku                            *Sku_Status                          `json:"sku,omitempty"`
	Type                           *string                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &BlobServiceProperties_Status{}

// ConvertStatusFrom populates our BlobServiceProperties_Status from the provided source
func (blobServicePropertiesStatus *BlobServiceProperties_Status) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == blobServicePropertiesStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(blobServicePropertiesStatus)
}

// ConvertStatusTo populates the provided destination from our BlobServiceProperties_Status
func (blobServicePropertiesStatus *BlobServiceProperties_Status) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == blobServicePropertiesStatus {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(blobServicePropertiesStatus)
}

//Storage version of v1alpha1api20210401.StorageAccountsBlobServices_Spec
type StorageAccountsBlobServices_Spec struct {
	AutomaticSnapshotPolicyEnabled *bool                         `json:"automaticSnapshotPolicyEnabled,omitempty"`
	ChangeFeed                     *ChangeFeed                   `json:"changeFeed,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRules                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                       `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy        `json:"deleteRetentionPolicy,omitempty"`
	IsVersioningEnabled            *bool                         `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy `json:"lastAccessTimeTrackingPolicy,omitempty"`
	Location                       *string                       `json:"location,omitempty"`
	OriginalVersion                string                        `json:"originalVersion"`

	// +kubebuilder:validation:Required
	Owner         genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner" kind:"StorageAccount"`
	PropertyBag   genruntime.PropertyBag            `json:"$propertyBag,omitempty"`
	RestorePolicy *RestorePolicyProperties          `json:"restorePolicy,omitempty"`
	Tags          map[string]string                 `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsBlobServices_Spec{}

// ConvertSpecFrom populates our StorageAccountsBlobServices_Spec from the provided source
func (storageAccountsBlobServicesSpec *StorageAccountsBlobServices_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == storageAccountsBlobServicesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(storageAccountsBlobServicesSpec)
}

// ConvertSpecTo populates the provided destination from our StorageAccountsBlobServices_Spec
func (storageAccountsBlobServicesSpec *StorageAccountsBlobServices_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == storageAccountsBlobServicesSpec {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(storageAccountsBlobServicesSpec)
}

//Storage version of v1alpha1api20210401.ChangeFeed
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/ChangeFeed
type ChangeFeed struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionInDays *int                   `json:"retentionInDays,omitempty"`
}

//Storage version of v1alpha1api20210401.ChangeFeed_Status
type ChangeFeed_Status struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionInDays *int                   `json:"retentionInDays,omitempty"`
}

//Storage version of v1alpha1api20210401.CorsRules
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/CorsRules
type CorsRules struct {
	CorsRules   []CorsRule             `json:"corsRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.CorsRules_Status
type CorsRules_Status struct {
	CorsRules   []CorsRule_Status      `json:"corsRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.DeleteRetentionPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/DeleteRetentionPolicy
type DeleteRetentionPolicy struct {
	Days        *int                   `json:"days,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.DeleteRetentionPolicy_Status
type DeleteRetentionPolicy_Status struct {
	Days        *int                   `json:"days,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.LastAccessTimeTrackingPolicy
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/LastAccessTimeTrackingPolicy
type LastAccessTimeTrackingPolicy struct {
	BlobType                  []string               `json:"blobType,omitempty"`
	Enable                    *bool                  `json:"enable,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TrackingGranularityInDays *int                   `json:"trackingGranularityInDays,omitempty"`
}

//Storage version of v1alpha1api20210401.LastAccessTimeTrackingPolicy_Status
type LastAccessTimeTrackingPolicy_Status struct {
	BlobType                  []string               `json:"blobType,omitempty"`
	Enable                    *bool                  `json:"enable,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TrackingGranularityInDays *int                   `json:"trackingGranularityInDays,omitempty"`
}

//Storage version of v1alpha1api20210401.RestorePolicyProperties
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/RestorePolicyProperties
type RestorePolicyProperties struct {
	Days        *int                   `json:"days,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.RestorePolicyProperties_Status
type RestorePolicyProperties_Status struct {
	Days            *int                   `json:"days,omitempty"`
	Enabled         *bool                  `json:"enabled,omitempty"`
	LastEnabledTime *string                `json:"lastEnabledTime,omitempty"`
	MinRestoreTime  *string                `json:"minRestoreTime,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.CorsRule
//Generated from: https://schema.management.azure.com/schemas/2021-04-01/Microsoft.Storage.json#/definitions/CorsRule
type CorsRule struct {
	AllowedHeaders  []string               `json:"allowedHeaders,omitempty"`
	AllowedMethods  []string               `json:"allowedMethods,omitempty"`
	AllowedOrigins  []string               `json:"allowedOrigins,omitempty"`
	ExposedHeaders  []string               `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

//Storage version of v1alpha1api20210401.CorsRule_Status
type CorsRule_Status struct {
	AllowedHeaders  []string               `json:"allowedHeaders,omitempty"`
	AllowedMethods  []string               `json:"allowedMethods,omitempty"`
	AllowedOrigins  []string               `json:"allowedOrigins,omitempty"`
	ExposedHeaders  []string               `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

func init() {
	SchemeBuilder.Register(&StorageAccountsBlobService{}, &StorageAccountsBlobServiceList{})
}
