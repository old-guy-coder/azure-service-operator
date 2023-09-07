// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220702storage

import (
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime/conditions"
	"github.com/pkg/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +kubebuilder:rbac:groups=compute.azure.com,resources=diskencryptionsets,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=compute.azure.com,resources={diskencryptionsets/status,diskencryptionsets/finalizers},verbs=get;update;patch

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="Severity",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].severity"
// +kubebuilder:printcolumn:name="Reason",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].reason"
// +kubebuilder:printcolumn:name="Message",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].message"
// Storage version of v1api20220702.DiskEncryptionSet
// Generator information:
// - Generated from: /compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/diskEncryptionSet.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}
type DiskEncryptionSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DiskEncryptionSet_Spec   `json:"spec,omitempty"`
	Status            DiskEncryptionSet_STATUS `json:"status,omitempty"`
}

var _ conditions.Conditioner = &DiskEncryptionSet{}

// GetConditions returns the conditions of the resource
func (encryptionSet *DiskEncryptionSet) GetConditions() conditions.Conditions {
	return encryptionSet.Status.Conditions
}

// SetConditions sets the conditions on the resource status
func (encryptionSet *DiskEncryptionSet) SetConditions(conditions conditions.Conditions) {
	encryptionSet.Status.Conditions = conditions
}

var _ genruntime.KubernetesResource = &DiskEncryptionSet{}

// AzureName returns the Azure name of the resource
func (encryptionSet *DiskEncryptionSet) AzureName() string {
	return encryptionSet.Spec.AzureName
}

// GetAPIVersion returns the ARM API version of the resource. This is always "2022-07-02"
func (encryptionSet DiskEncryptionSet) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetResourceScope returns the scope of the resource
func (encryptionSet *DiskEncryptionSet) GetResourceScope() genruntime.ResourceScope {
	return genruntime.ResourceScopeResourceGroup
}

// GetSpec returns the specification of this resource
func (encryptionSet *DiskEncryptionSet) GetSpec() genruntime.ConvertibleSpec {
	return &encryptionSet.Spec
}

// GetStatus returns the status of this resource
func (encryptionSet *DiskEncryptionSet) GetStatus() genruntime.ConvertibleStatus {
	return &encryptionSet.Status
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Compute/diskEncryptionSets"
func (encryptionSet *DiskEncryptionSet) GetType() string {
	return "Microsoft.Compute/diskEncryptionSets"
}

// NewEmptyStatus returns a new empty (blank) status
func (encryptionSet *DiskEncryptionSet) NewEmptyStatus() genruntime.ConvertibleStatus {
	return &DiskEncryptionSet_STATUS{}
}

// Owner returns the ResourceReference of the owner
func (encryptionSet *DiskEncryptionSet) Owner() *genruntime.ResourceReference {
	group, kind := genruntime.LookupOwnerGroupKind(encryptionSet.Spec)
	return encryptionSet.Spec.Owner.AsResourceReference(group, kind)
}

// SetStatus sets the status of this resource
func (encryptionSet *DiskEncryptionSet) SetStatus(status genruntime.ConvertibleStatus) error {
	// If we have exactly the right type of status, assign it
	if st, ok := status.(*DiskEncryptionSet_STATUS); ok {
		encryptionSet.Status = *st
		return nil
	}

	// Convert status to required version
	var st DiskEncryptionSet_STATUS
	err := status.ConvertStatusTo(&st)
	if err != nil {
		return errors.Wrap(err, "failed to convert status")
	}

	encryptionSet.Status = st
	return nil
}

// Hub marks that this DiskEncryptionSet is the hub type for conversion
func (encryptionSet *DiskEncryptionSet) Hub() {}

// OriginalGVK returns a GroupValueKind for the original API version used to create the resource
func (encryptionSet *DiskEncryptionSet) OriginalGVK() *schema.GroupVersionKind {
	return &schema.GroupVersionKind{
		Group:   GroupVersion.Group,
		Version: encryptionSet.Spec.OriginalVersion,
		Kind:    "DiskEncryptionSet",
	}
}

// +kubebuilder:object:root=true
// Storage version of v1api20220702.DiskEncryptionSet
// Generator information:
// - Generated from: /compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/diskEncryptionSet.json
// - ARM URI: /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}
type DiskEncryptionSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DiskEncryptionSet `json:"items"`
}

// Storage version of v1api20220702.APIVersion
// +kubebuilder:validation:Enum={"2022-07-02"}
type APIVersion string

const APIVersion_Value = APIVersion("2022-07-02")

// Storage version of v1api20220702.DiskEncryptionSet_Spec
type DiskEncryptionSet_Spec struct {
	ActiveKey *KeyForDiskEncryptionSet `json:"activeKey,omitempty"`

	// AzureName: The name of the resource in Azure. This is often the same as the name of the resource in Kubernetes but it
	// doesn't have to be.
	AzureName                   string                         `json:"azureName,omitempty"`
	EncryptionType              *string                        `json:"encryptionType,omitempty"`
	FederatedClientId           *string                        `json:"federatedClientId,omitempty" optionalConfigMapPair:"FederatedClientId"`
	FederatedClientIdFromConfig *genruntime.ConfigMapReference `json:"federatedClientIdFromConfig,omitempty" optionalConfigMapPair:"FederatedClientId"`
	Identity                    *EncryptionSetIdentity         `json:"identity,omitempty"`
	Location                    *string                        `json:"location,omitempty"`
	OriginalVersion             string                         `json:"originalVersion,omitempty"`

	// +kubebuilder:validation:Required
	// Owner: The owner of the resource. The owner controls where the resource goes when it is deployed. The owner also
	// controls the resources lifecycle. When the owner is deleted the resource will also be deleted. Owner is expected to be a
	// reference to a resources.azure.com/ResourceGroup resource
	Owner                             *genruntime.KnownResourceReference `group:"resources.azure.com" json:"owner,omitempty" kind:"ResourceGroup"`
	PropertyBag                       genruntime.PropertyBag             `json:"$propertyBag,omitempty"`
	RotationToLatestKeyVersionEnabled *bool                              `json:"rotationToLatestKeyVersionEnabled,omitempty"`
	Tags                              map[string]string                  `json:"tags,omitempty"`
}

var _ genruntime.ConvertibleSpec = &DiskEncryptionSet_Spec{}

// ConvertSpecFrom populates our DiskEncryptionSet_Spec from the provided source
func (encryptionSet *DiskEncryptionSet_Spec) ConvertSpecFrom(source genruntime.ConvertibleSpec) error {
	if source == encryptionSet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return source.ConvertSpecTo(encryptionSet)
}

// ConvertSpecTo populates the provided destination from our DiskEncryptionSet_Spec
func (encryptionSet *DiskEncryptionSet_Spec) ConvertSpecTo(destination genruntime.ConvertibleSpec) error {
	if destination == encryptionSet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleSpec")
	}

	return destination.ConvertSpecFrom(encryptionSet)
}

// Storage version of v1api20220702.DiskEncryptionSet_STATUS
// disk encryption set resource.
type DiskEncryptionSet_STATUS struct {
	ActiveKey                         *KeyForDiskEncryptionSet_STATUS  `json:"activeKey,omitempty"`
	AutoKeyRotationError              *ApiError_STATUS                 `json:"autoKeyRotationError,omitempty"`
	Conditions                        []conditions.Condition           `json:"conditions,omitempty"`
	EncryptionType                    *string                          `json:"encryptionType,omitempty"`
	FederatedClientId                 *string                          `json:"federatedClientId,omitempty"`
	Id                                *string                          `json:"id,omitempty"`
	Identity                          *EncryptionSetIdentity_STATUS    `json:"identity,omitempty"`
	LastKeyRotationTimestamp          *string                          `json:"lastKeyRotationTimestamp,omitempty"`
	Location                          *string                          `json:"location,omitempty"`
	Name                              *string                          `json:"name,omitempty"`
	PreviousKeys                      []KeyForDiskEncryptionSet_STATUS `json:"previousKeys,omitempty"`
	PropertyBag                       genruntime.PropertyBag           `json:"$propertyBag,omitempty"`
	ProvisioningState                 *string                          `json:"provisioningState,omitempty"`
	RotationToLatestKeyVersionEnabled *bool                            `json:"rotationToLatestKeyVersionEnabled,omitempty"`
	Tags                              map[string]string                `json:"tags,omitempty"`
	Type                              *string                          `json:"type,omitempty"`
}

var _ genruntime.ConvertibleStatus = &DiskEncryptionSet_STATUS{}

// ConvertStatusFrom populates our DiskEncryptionSet_STATUS from the provided source
func (encryptionSet *DiskEncryptionSet_STATUS) ConvertStatusFrom(source genruntime.ConvertibleStatus) error {
	if source == encryptionSet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return source.ConvertStatusTo(encryptionSet)
}

// ConvertStatusTo populates the provided destination from our DiskEncryptionSet_STATUS
func (encryptionSet *DiskEncryptionSet_STATUS) ConvertStatusTo(destination genruntime.ConvertibleStatus) error {
	if destination == encryptionSet {
		return errors.New("attempted conversion between unrelated implementations of github.com/Azure/azure-service-operator/v2/pkg/genruntime/ConvertibleStatus")
	}

	return destination.ConvertStatusFrom(encryptionSet)
}

// Storage version of v1api20220702.ApiError_STATUS
// Api error.
type ApiError_STATUS struct {
	Code        *string                `json:"code,omitempty"`
	Details     []ApiErrorBase_STATUS  `json:"details,omitempty"`
	Innererror  *InnerError_STATUS     `json:"innererror,omitempty"`
	Message     *string                `json:"message,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Target      *string                `json:"target,omitempty"`
}

// Storage version of v1api20220702.EncryptionSetIdentity
// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used
// to encrypt disks.
type EncryptionSetIdentity struct {
	PropertyBag            genruntime.PropertyBag        `json:"$propertyBag,omitempty"`
	Type                   *string                       `json:"type,omitempty"`
	UserAssignedIdentities []UserAssignedIdentityDetails `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20220702.EncryptionSetIdentity_STATUS
// The managed identity for the disk encryption set. It should be given permission on the key vault before it can be used
// to encrypt disks.
type EncryptionSetIdentity_STATUS struct {
	PrincipalId            *string                                                        `json:"principalId,omitempty"`
	PropertyBag            genruntime.PropertyBag                                         `json:"$propertyBag,omitempty"`
	TenantId               *string                                                        `json:"tenantId,omitempty"`
	Type                   *string                                                        `json:"type,omitempty"`
	UserAssignedIdentities map[string]EncryptionSetIdentity_UserAssignedIdentities_STATUS `json:"userAssignedIdentities,omitempty"`
}

// Storage version of v1api20220702.KeyForDiskEncryptionSet
// Key Vault Key Url to be used for server side encryption of Managed Disks and Snapshots
type KeyForDiskEncryptionSet struct {
	KeyUrl           *string                        `json:"keyUrl,omitempty" optionalConfigMapPair:"KeyUrl"`
	KeyUrlFromConfig *genruntime.ConfigMapReference `json:"keyUrlFromConfig,omitempty" optionalConfigMapPair:"KeyUrl"`
	PropertyBag      genruntime.PropertyBag         `json:"$propertyBag,omitempty"`
	SourceVault      *SourceVault                   `json:"sourceVault,omitempty"`
}

// Storage version of v1api20220702.KeyForDiskEncryptionSet_STATUS
// Key Vault Key Url to be used for server side encryption of Managed Disks and Snapshots
type KeyForDiskEncryptionSet_STATUS struct {
	KeyUrl      *string                `json:"keyUrl,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	SourceVault *SourceVault_STATUS    `json:"sourceVault,omitempty"`
}

// Storage version of v1api20220702.ApiErrorBase_STATUS
// Api error base.
type ApiErrorBase_STATUS struct {
	Code        *string                `json:"code,omitempty"`
	Message     *string                `json:"message,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
	Target      *string                `json:"target,omitempty"`
}

// Storage version of v1api20220702.EncryptionSetIdentity_UserAssignedIdentities_STATUS
type EncryptionSetIdentity_UserAssignedIdentities_STATUS struct {
	ClientId    *string                `json:"clientId,omitempty"`
	PrincipalId *string                `json:"principalId,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220702.InnerError_STATUS
// Inner error details.
type InnerError_STATUS struct {
	Errordetail   *string                `json:"errordetail,omitempty"`
	Exceptiontype *string                `json:"exceptiontype,omitempty"`
	PropertyBag   genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220702.SourceVault
// The vault id is an Azure Resource Manager Resource id in the form
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
type SourceVault struct {
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`

	// Reference: Resource Id
	Reference *genruntime.ResourceReference `armReference:"Id" json:"reference,omitempty"`
}

// Storage version of v1api20220702.SourceVault_STATUS
// The vault id is an Azure Resource Manager Resource id in the form
// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.KeyVault/vaults/{vaultName}
type SourceVault_STATUS struct {
	Id          *string                `json:"id,omitempty"`
	PropertyBag genruntime.PropertyBag `json:"$propertyBag,omitempty"`
}

// Storage version of v1api20220702.UserAssignedIdentityDetails
// Information about the user assigned identity for the resource
type UserAssignedIdentityDetails struct {
	PropertyBag genruntime.PropertyBag       `json:"$propertyBag,omitempty"`
	Reference   genruntime.ResourceReference `armReference:"Reference" json:"reference,omitempty"`
}

func init() {
	SchemeBuilder.Register(&DiskEncryptionSet{}, &DiskEncryptionSetList{})
}