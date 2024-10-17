// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

type PrivateDnsZonesVirtualNetworkLink_STATUS_ARM struct {
	// Etag: The ETag of the virtual network link.
	Etag *string `json:"etag,omitempty"`

	// Id: Fully qualified resource Id for the resource. Example -
	// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/privateDnsZones/{privateDnsZoneName}'.
	Id *string `json:"id,omitempty"`

	// Location: The Azure Region where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the virtual network link to the Private DNS zone.
	Properties *VirtualNetworkLinkProperties_STATUS_ARM `json:"properties,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. Example - 'Microsoft.Network/privateDnsZones'.
	Type *string `json:"type,omitempty"`
}

// Represents the properties of the Private DNS zone.
type VirtualNetworkLinkProperties_STATUS_ARM struct {
	// ProvisioningState: The provisioning state of the resource. This is a read-only property and any attempt to set this
	// value will be ignored.
	ProvisioningState *VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`

	// RegistrationEnabled: Is auto-registration of virtual machine records in the virtual network in the Private DNS zone
	// enabled?
	RegistrationEnabled *bool `json:"registrationEnabled,omitempty"`

	// VirtualNetwork: The reference of the virtual network.
	VirtualNetwork *SubResource_STATUS_ARM `json:"virtualNetwork,omitempty"`

	// VirtualNetworkLinkState: The status of the virtual network link to the Private DNS zone. Possible values are
	// 'InProgress' and 'Done'. This is a read-only property and any attempt to set this value will be ignored.
	VirtualNetworkLinkState *VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM `json:"virtualNetworkLinkState,omitempty"`
}

// Reference to another subresource.
type SubResource_STATUS_ARM struct {
	// Id: Resource ID.
	Id *string `json:"id,omitempty"`
}

type VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM string

const (
	VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Canceled  = VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM("Canceled")
	VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Creating  = VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM("Creating")
	VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Deleting  = VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM("Deleting")
	VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Failed    = VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM("Failed")
	VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Succeeded = VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM("Succeeded")
	VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Updating  = VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM("Updating")
)

// Mapping from string to VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM
var virtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Values = map[string]VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM{
	"canceled":  VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Canceled,
	"creating":  VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Creating,
	"deleting":  VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Deleting,
	"failed":    VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Failed,
	"succeeded": VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Succeeded,
	"updating":  VirtualNetworkLinkProperties_ProvisioningState_STATUS_ARM_Updating,
}

type VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM string

const (
	VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM_Completed  = VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM("Completed")
	VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM_InProgress = VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM("InProgress")
)

// Mapping from string to VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM
var virtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM_Values = map[string]VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM{
	"completed":  VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM_Completed,
	"inprogress": VirtualNetworkLinkProperties_VirtualNetworkLinkState_STATUS_ARM_InProgress,
}
