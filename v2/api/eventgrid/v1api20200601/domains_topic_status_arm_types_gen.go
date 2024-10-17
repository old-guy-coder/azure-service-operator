// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20200601

type DomainsTopic_STATUS_ARM struct {
	// Id: Fully qualified identifier of the resource.
	Id *string `json:"id,omitempty"`

	// Name: Name of the resource.
	Name *string `json:"name,omitempty"`

	// Properties: Properties of the Domain Topic.
	Properties *DomainTopicProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: The system metadata relating to Domain Topic resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: Type of the resource.
	Type *string `json:"type,omitempty"`
}

// Properties of the Domain Topic.
type DomainTopicProperties_STATUS_ARM struct {
	// ProvisioningState: Provisioning state of the domain topic.
	ProvisioningState *DomainTopicProperties_ProvisioningState_STATUS_ARM `json:"provisioningState,omitempty"`
}

type DomainTopicProperties_ProvisioningState_STATUS_ARM string

const (
	DomainTopicProperties_ProvisioningState_STATUS_ARM_Canceled  = DomainTopicProperties_ProvisioningState_STATUS_ARM("Canceled")
	DomainTopicProperties_ProvisioningState_STATUS_ARM_Creating  = DomainTopicProperties_ProvisioningState_STATUS_ARM("Creating")
	DomainTopicProperties_ProvisioningState_STATUS_ARM_Deleting  = DomainTopicProperties_ProvisioningState_STATUS_ARM("Deleting")
	DomainTopicProperties_ProvisioningState_STATUS_ARM_Failed    = DomainTopicProperties_ProvisioningState_STATUS_ARM("Failed")
	DomainTopicProperties_ProvisioningState_STATUS_ARM_Succeeded = DomainTopicProperties_ProvisioningState_STATUS_ARM("Succeeded")
	DomainTopicProperties_ProvisioningState_STATUS_ARM_Updating  = DomainTopicProperties_ProvisioningState_STATUS_ARM("Updating")
)

// Mapping from string to DomainTopicProperties_ProvisioningState_STATUS_ARM
var domainTopicProperties_ProvisioningState_STATUS_ARM_Values = map[string]DomainTopicProperties_ProvisioningState_STATUS_ARM{
	"canceled":  DomainTopicProperties_ProvisioningState_STATUS_ARM_Canceled,
	"creating":  DomainTopicProperties_ProvisioningState_STATUS_ARM_Creating,
	"deleting":  DomainTopicProperties_ProvisioningState_STATUS_ARM_Deleting,
	"failed":    DomainTopicProperties_ProvisioningState_STATUS_ARM_Failed,
	"succeeded": DomainTopicProperties_ProvisioningState_STATUS_ARM_Succeeded,
	"updating":  DomainTopicProperties_ProvisioningState_STATUS_ARM_Updating,
}
