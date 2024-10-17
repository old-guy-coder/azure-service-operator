// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230630

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersConfiguration_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a configuration.
	Properties *ConfigurationProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersConfiguration_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-06-30"
func (configuration FlexibleServersConfiguration_Spec_ARM) GetAPIVersion() string {
	return "2023-06-30"
}

// GetName returns the Name of the resource
func (configuration *FlexibleServersConfiguration_Spec_ARM) GetName() string {
	return configuration.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/configurations"
func (configuration *FlexibleServersConfiguration_Spec_ARM) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/configurations"
}

// The properties of a configuration.
type ConfigurationProperties_ARM struct {
	// CurrentValue: Current value of the configuration.
	CurrentValue *string `json:"currentValue,omitempty"`

	// Source: Source of the configuration.
	Source *ConfigurationProperties_Source_ARM `json:"source,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:validation:Enum={"system-default","user-override"}
type ConfigurationProperties_Source_ARM string

const (
	ConfigurationProperties_Source_ARM_SystemDefault = ConfigurationProperties_Source_ARM("system-default")
	ConfigurationProperties_Source_ARM_UserOverride  = ConfigurationProperties_Source_ARM("user-override")
)

// Mapping from string to ConfigurationProperties_Source_ARM
var configurationProperties_Source_ARM_Values = map[string]ConfigurationProperties_Source_ARM{
	"system-default": ConfigurationProperties_Source_ARM_SystemDefault,
	"user-override":  ConfigurationProperties_Source_ARM_UserOverride,
}
