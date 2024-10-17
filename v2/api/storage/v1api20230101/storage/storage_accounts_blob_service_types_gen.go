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

// +kubebuilder:rbac:groups=storage.azure.com,resources=storageaccountsblobservices,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=storage.azure.com,resources={storageaccountsblobservices/status,storageaccountsblobservices/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20230101.StorageAccountsBlobService
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2023-01-01/blob.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default
type StorageAccountsBlobService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountsBlobService_Spec   `json:"spec,omitempty"`
	Status            StorageAccountsBlobService_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &StorageAccountsBlobService{}

// GetConditions returns the conditions of the resource
func (service *StorageAccountsBlobService) GetConditions() conditions.Conditions {
	return service.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (service *StorageAccountsBlobService) SetConditions(conditions conditions.Conditions) {
	service.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &StorageAccountsBlobService{}

// AzureName returns the Azure name of the resource (always "default")
func (service *StorageAccountsBlobService) AzureName() string {
	return "default"
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (service StorageAccountsBlobService) GetAPIVersion() string {
	return "2023-01-01"
}

// GetResourceScope returns the scope of the resource
func (service *StorageAccountsBlobService) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (service *StorageAccountsBlobService) GetSpec() genruntime.ConvertibleSpec {
	return &service.Spec
}

// GetStatus returns the status of this resource
func (service *StorageAccountsBlobService) GetStatus() genruntime.ConvertibleStatus {
	return &service.Status
}

// GetSupportedOperations returns the operations supported by the resource
func (service *StorageAccountsBlobService) GetSupportedOperations() []genruntime.ResourceOperation {
	return []genruntime.ResourceOperation{
		genruntime.ResourceOperationGet,
		genruntime.ResourceOperationPut,
	}
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/blobServices"
func (service *StorageAccountsBlobService) GetType() string {
	return "Microsoft.Storage/storageAccounts/blobServices"
}

// NewEmptyStatus returns a new empty (blank) status
func (service *StorageAccountsBlobService) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &StorageAccountsBlobService_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (service *StorageAccountsBlobService) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(service.Spec)
	return service.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (service *StorageAccountsBlobService) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*StorageAccountsBlobService_STATUS); ok {
		service.Status = *st
		return nil
	}

	// Convert status to required version
	var st StorageAccountsBlobService_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	service.Status = st
	return nil
}

// Hub marks that this StorageAccountsBlobService is the hub type for conversion
func (service *StorageAccountsBlobService) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (service *StorageAccountsBlobService) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: service.Spec.OriginalVersion,
		Kind:    "StorageAccountsBlobService",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20230101.StorageAccountsBlobService
// Generator information:
// - Generated from: /storage/resource-manager/Microsoft.Storage/stable/2023-01-01/blob.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default
type StorageAccountsBlobServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []StorageAccountsBlobService `json:"items"`
}

// Storage version of v1api20230101.StorageAccountsBlobService_Spec
type StorageAccountsBlobService_Spec struct {
	AutomaticSnapshotPolicyEnabled *bool                         `json:"automaticSnapshotPolicyEnabled,omitempty"`
	ChangeFeed                     *ChangeFeed                   `json:"changeFeed,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRules                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                       `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy        `json:"deleteRetentionPolicy,omitempty"`
	IsVersioningEnabled            *bool                         `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy `json:"lastAccessTimeTrackingPolicy,omitempty"`
	OriginalVersion                string                        `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a storage.azure.com/StorageAccount resource
	Owner         *genruntime.KnownResourceReference `group:"storage.azure.com" json:"owner,omitempty" kind:"StorageAccount"`
	PropertyBag   genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RestorePolicy *RestorePolicyProperties           `json:"restorePolicy,omitempty"`
}

var _ genruntime.ConvertibleSpec = &StorageAccountsBlobService_Spec{}

// ConvertSpecFrom populates our StorageAccountsBlobService_Spec from the provided source
func (service *StorageAccountsBlobService_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(service)
}

// ConvertSpecTo populates the provided destination from our StorageAccountsBlobService_Spec
func (service *StorageAccountsBlobService_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(service)
}

// Storage version of v1api20230101.StorageAccountsBlobService_STATUS
type StorageAccountsBlobService_STATUS struct {
	AutomaticSnapshotPolicyEnabled *bool                                `json:"automaticSnapshotPolicyEnabled,omitempty"`
	ChangeFeed                     *ChangeFeed_STATUS                   `json:"changeFeed,omitempty"`
	Conditions                     []conditions.Condition               `json:"conditions,omitempty"`
	ContainerDeleteRetentionPolicy *DeleteRetentionPolicy_STATUS        `json:"containerDeleteRetentionPolicy,omitempty"`
	Cors                           *CorsRules_STATUS                    `json:"cors,omitempty"`
	DefaultServiceVersion          *string                              `json:"defaultServiceVersion,omitempty"`
	DeleteRetentionPolicy          *DeleteRetentionPolicy_STATUS        `json:"deleteRetentionPolicy,omitempty"`
	Id                             *string                              `json:"id,omitempty"`
	IsVersioningEnabled            *bool                                `json:"isVersioningEnabled,omitempty"`
	LastAccessTimeTrackingPolicy   *LastAccessTimeTrackingPolicy_STATUS `json:"lastAccessTimeTrackingPolicy,omitempty"`
	Name                           *string                              `json:"name,omitempty"`
	PropertyBag                    genruntime.PropertyBag               `json:"$propertyBag,omitempty"`
	RestorePolicy                  *RestorePolicyProperties_STATUS      `json:"restorePolicy,omitempty"`
	Sku                            *Sku_STATUS                          `json:"sku,omitempty"`
	Type                           *string                              `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &StorageAccountsBlobService_STATUS{}

// ConvertStatusFrom populates our StorageAccountsBlobService_STATUS from the provided source
func (service *StorageAccountsBlobService_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(service)
}

// ConvertStatusTo populates the provided destination from our StorageAccountsBlobService_STATUS
func (service *StorageAccountsBlobService_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == service {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(service)
}

// Storage version of v1api20230101.ChangeFeed
// The blob service properties for change feed events.
type ChangeFeed struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionInDays *int                   `json:"retentionInDays,omitempty"`
}

// Storage version of v1api20230101.ChangeFeed_STATUS
// The blob service properties for change feed events.
type ChangeFeed_STATUS struct {
	Enabled         *bool                  `json:"enabled,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	RetentionInDays *int                   `json:"retentionInDays,omitempty"`
}

// Storage version of v1api20230101.CorsRules
// Sets the CORS rules. You can include up to five CorsRule elements in the request.
type CorsRules struct {
	CorsRules   []CorsRule             `json:"corsRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.CorsRules_STATUS
// Sets the CORS rules. You can include up to five CorsRule elements in the request.
type CorsRules_STATUS struct {
	CorsRules   []CorsRule_STATUS      `json:"corsRules,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.DeleteRetentionPolicy
// The service properties for soft delete.
type DeleteRetentionPolicy struct {
	AllowPermanentDelete *bool                  `json:"allowPermanentDelete,omitempty"`
	Days                 *int                   `json:"days,omitempty"`
	Enabled              *bool                  `json:"enabled,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.DeleteRetentionPolicy_STATUS
// The service properties for soft delete.
type DeleteRetentionPolicy_STATUS struct {
	AllowPermanentDelete *bool                  `json:"allowPermanentDelete,omitempty"`
	Days                 *int                   `json:"days,omitempty"`
	Enabled              *bool                  `json:"enabled,omitempty"`
	PropertyBag          genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.LastAccessTimeTrackingPolicy
// The blob service properties for Last access time based tracking policy.
type LastAccessTimeTrackingPolicy struct {
	BlobType                  []string               `json:"blobType,omitempty"`
	Enable                    *bool                  `json:"enable,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TrackingGranularityInDays *int                   `json:"trackingGranularityInDays,omitempty"`
}

// Storage version of v1api20230101.LastAccessTimeTrackingPolicy_STATUS
// The blob service properties for Last access time based tracking policy.
type LastAccessTimeTrackingPolicy_STATUS struct {
	BlobType                  []string               `json:"blobType,omitempty"`
	Enable                    *bool                  `json:"enable,omitempty"`
	Name                      *string                `json:"name,omitempty"`
	PropertyBag               genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	TrackingGranularityInDays *int                   `json:"trackingGranularityInDays,omitempty"`
}

// Storage version of v1api20230101.RestorePolicyProperties
// The blob service properties for blob restore policy
type RestorePolicyProperties struct {
	Days        *int                   `json:"days,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.RestorePolicyProperties_STATUS
// The blob service properties for blob restore policy
type RestorePolicyProperties_STATUS struct {
	Days            *int                   `json:"days,omitempty"`
	Enabled         *bool                  `json:"enabled,omitempty"`
	LastEnabledTime *string                `json:"lastEnabledTime,omitempty"`
	MinRestoreTime  *string                `json:"minRestoreTime,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.CorsRule
// Specifies a CORS rule for the Blob service.
type CorsRule struct {
	AllowedHeaders  []string               `json:"allowedHeaders,omitempty"`
	AllowedMethods  []string               `json:"allowedMethods,omitempty"`
	AllowedOrigins  []string               `json:"allowedOrigins,omitempty"`
	ExposedHeaders  []string               `json:"exposedHeaders,omitempty"`
	MaxAgeInSeconds *int                   `json:"maxAgeInSeconds,omitempty"`
	PropertyBag     genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20230101.CorsRule_STATUS
// Specifies a CORS rule for the Blob service.
type CorsRule_STATUS struct {
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
