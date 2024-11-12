// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type SqlDatabaseContainerThroughputSetting_Spec struct {
	// Location: The location of the resource group to which the resource belongs.
	Location *string `json:"location,omitempty"`
	Name     string  `json:"name,omitempty"`

	// Properties: Properties to update Azure Cosmos DB resource throughput.
	Properties *ThroughputSettingsUpdateProperties `json:"properties,omitempty"`
	Tags       map[string]string                   `json:"tags,omitempty"`
}

var _ genruntime.ARMResourceSpec = &SqlDatabaseContainerThroughputSetting_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-15"
func (setting SqlDatabaseContainerThroughputSetting_Spec) GetAPIVersion() string {
	return "2023-11-15"
}

// GetName returns the Name of the resource
func (setting *SqlDatabaseContainerThroughputSetting_Spec) GetName() string {
	return setting.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
func (setting *SqlDatabaseContainerThroughputSetting_Spec) GetType() string {
	return "Microsoft.DocumentDB/databaseAccounts/sqlDatabases/containers/throughputSettings"
}