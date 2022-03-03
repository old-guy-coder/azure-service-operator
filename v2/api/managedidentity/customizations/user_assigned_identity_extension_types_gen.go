// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package customizations

import (
	managedidentity "github.com/Azure/azure-service-operator/v2/api/managedidentity/v1alpha1api20181130"
	"github.com/Azure/azure-service-operator/v2/api/managedidentity/v1alpha1api20181130storage"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type UserAssignedIdentityExtension struct {
}

// GetExtendedResources Returns the KubernetesResource slice for Resource versions
func (extension *UserAssignedIdentityExtension) GetExtendedResources() []genruntime.KubernetesResource {
	return []genruntime.KubernetesResource{
		&managedidentity.UserAssignedIdentity{},
		&v1alpha1api20181130storage.UserAssignedIdentity{}}
}