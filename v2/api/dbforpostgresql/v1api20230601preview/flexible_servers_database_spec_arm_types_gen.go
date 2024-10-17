// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230601preview

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type FlexibleServersDatabase_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a database.
	Properties *DatabaseProperties_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &FlexibleServersDatabase_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-06-01-preview"
func (database FlexibleServersDatabase_Spec_ARM) GetAPIVersion() string {
	return "2023-06-01-preview"
}

// GetName returns the Name of the resource
func (database *FlexibleServersDatabase_Spec_ARM) GetName() string {
	return database.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DBforPostgreSQL/flexibleServers/databases"
func (database *FlexibleServersDatabase_Spec_ARM) GetType() string {
	return "Microsoft.DBforPostgreSQL/flexibleServers/databases"
}

// The properties of a database.
type DatabaseProperties_ARM struct {
	// Charset: The charset of the database.
	Charset *string `json:"charset,omitempty"`

	// Collation: The collation of the database.
	Collation *string `json:"collation,omitempty"`
}
