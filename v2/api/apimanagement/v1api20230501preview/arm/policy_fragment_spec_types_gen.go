// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type PolicyFragment_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties of the Policy Fragment.
	Properties *PolicyFragmentContractProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &PolicyFragment_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-05-01-preview"
func (fragment PolicyFragment_Spec) GetAPIVersion() string {
	return "2023-05-01-preview"
}

// GetName returns the Name of the resource
func (fragment *PolicyFragment_Spec) GetName() string {
	return fragment.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.ApiManagement/service/policyFragments"
func (fragment *PolicyFragment_Spec) GetType() string {
	return "Microsoft.ApiManagement/service/policyFragments"
}

// Policy fragment contract properties.
type PolicyFragmentContractProperties struct {
	// Description: Policy fragment description.
	Description *string `json:"description,omitempty"`

	// Format: Format of the policy fragment content.
	Format *PolicyFragmentContractProperties_Format `json:"format,omitempty"`

	// Value: Contents of the policy fragment.
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:validation:Enum={"rawxml","xml"}
type PolicyFragmentContractProperties_Format string

const (
	PolicyFragmentContractProperties_Format_Rawxml = PolicyFragmentContractProperties_Format("rawxml")
	PolicyFragmentContractProperties_Format_Xml    = PolicyFragmentContractProperties_Format("xml")
)

// Mapping from string to PolicyFragmentContractProperties_Format
var policyFragmentContractProperties_Format_Values = map[string]PolicyFragmentContractProperties_Format{
	"rawxml": PolicyFragmentContractProperties_Format_Rawxml,
	"xml":    PolicyFragmentContractProperties_Format_Xml,
}