// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsTableService_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a storage account’s Table service.
	Properties *StorageAccounts_TableService_Properties_Spec_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsTableService_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-01-01"
func (service StorageAccountsTableService_Spec_ARM) GetAPIVersion() string {
	return "2023-01-01"
}

// GetName returns the Name of the resource
func (service *StorageAccountsTableService_Spec_ARM) GetName() string {
	return service.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/tableServices"
func (service *StorageAccountsTableService_Spec_ARM) GetType() string {
	return "Microsoft.Storage/storageAccounts/tableServices"
}

type StorageAccounts_TableService_Properties_Spec_ARM struct {
	// Cors: Specifies CORS rules for the Table service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Table service.
	Cors *CorsRules_ARM `json:"cors,omitempty"`
}
