// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersConfiguration_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a configuration.
	Properties *ConfigurationProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersConfiguration_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-06-30"
func (configuration FlexibleServersConfiguration_Spec) GetAPIVersion() string {
	return "2023-06-30"
}

// GetName returns the Name of the resource
func (configuration *FlexibleServersConfiguration_Spec) GetName() string {
	return configuration.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforMySQL/flexibleServers/configurations"
func (configuration *FlexibleServersConfiguration_Spec) GetType() string {
	return "Microsoft.DBforMySQL/flexibleServers/configurations"
}

// The properties of a configuration.
type ConfigurationProperties struct {
	// CurrentValue: Current value of the configuration.
	CurrentValue *string `json:"currentValue,omitempty"`

	// Source: Source of the configuration.
	Source *ConfigurationProperties_Source `json:"source,omitempty"`

	// Value: Value of the configuration.
	Value *string `json:"value,omitempty"`
}

// +kubebuilder:validation:Enum={"system-default","user-override"}
type ConfigurationProperties_Source string

const (
	ConfigurationProperties_Source_SystemDefault = ConfigurationProperties_Source("system-default")
	ConfigurationProperties_Source_UserOverride  = ConfigurationProperties_Source("user-override")
)

// Mapping from string to ConfigurationProperties_Source
var configurationProperties_Source_Values = map[string]ConfigurationProperties_Source{
	"system-default": ConfigurationProperties_Source_SystemDefault,
	"user-override":  ConfigurationProperties_Source_UserOverride,
}