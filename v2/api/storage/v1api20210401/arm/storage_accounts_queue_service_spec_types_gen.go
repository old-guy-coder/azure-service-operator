// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type StorageAccountsQueueService_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: The properties of a storage account’s Queue service.
	Properties *StorageAccounts_QueueService_Properties_Spec `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &StorageAccountsQueueService_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-04-01"
func (service StorageAccountsQueueService_Spec) GetAPIVersion() string {
	return "2021-04-01"
}

// GetName returns the Name of the resource
func (service *StorageAccountsQueueService_Spec) GetName() string {
	return service.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Storage/storageAccounts/queueServices"
func (service *StorageAccountsQueueService_Spec) GetType() string {
	return "Microsoft.Storage/storageAccounts/queueServices"
}

type StorageAccounts_QueueService_Properties_Spec struct {
	// Cors: Specifies CORS rules for the Queue service. You can include up to five CorsRule elements in the request. If no
	// CorsRule elements are included in the request body, all CORS rules will be deleted, and CORS will be disabled for the
	// Queue service.
	Cors *CorsRules `json:"cors,omitempty"`
}