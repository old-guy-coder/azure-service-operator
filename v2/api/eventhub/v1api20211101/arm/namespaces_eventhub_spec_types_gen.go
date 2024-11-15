// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type NamespacesEventhub_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: Properties supplied to the Create Or Update Event Hub operation.
	Properties *Namespaces_Eventhub_Properties_Spec `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &NamespacesEventhub_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2021-11-01"
func (eventhub NamespacesEventhub_Spec) GetAPIVersion() string {
	return "2021-11-01"
}

// GetName returns the Name of the resource
func (eventhub *NamespacesEventhub_Spec) GetName() string {
	return eventhub.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.EventHub/namespaces/eventhubs"
func (eventhub *NamespacesEventhub_Spec) GetType() string {
	return "Microsoft.EventHub/namespaces/eventhubs"
}

type Namespaces_Eventhub_Properties_Spec struct {
	// CaptureDescription: Properties of capture description
	CaptureDescription *CaptureDescription `json:"captureDescription,omitempty"`

	// MessageRetentionInDays: Number of days to retain the events for this Event Hub, value should be 1 to 7 days
	MessageRetentionInDays *int `json:"messageRetentionInDays,omitempty"`

	// PartitionCount: Number of partitions created for the Event Hub, allowed values are from 1 to 32 partitions.
	PartitionCount *int `json:"partitionCount,omitempty"`
}

// Properties to configure capture description for eventhub
type CaptureDescription struct {
	// Destination: Properties of Destination where capture will be stored. (Storage Account, Blob Names)
	Destination *Destination `json:"destination,omitempty"`

	// Enabled: A value that indicates whether capture description is enabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Encoding: Enumerates the possible values for the encoding format of capture description. Note: 'AvroDeflate' will be
	// deprecated in New API Version
	Encoding *CaptureDescription_Encoding `json:"encoding,omitempty"`

	// IntervalInSeconds: The time window allows you to set the frequency with which the capture to Azure Blobs will happen,
	// value should between 60 to 900 seconds
	IntervalInSeconds *int `json:"intervalInSeconds,omitempty"`

	// SizeLimitInBytes: The size window defines the amount of data built up in your Event Hub before an capture operation,
	// value should be between 10485760 to 524288000 bytes
	SizeLimitInBytes *int `json:"sizeLimitInBytes,omitempty"`

	// SkipEmptyArchives: A value that indicates whether to Skip Empty Archives
	SkipEmptyArchives *bool `json:"skipEmptyArchives,omitempty"`
}

// +kubebuilder:validation:Enum={"Avro","AvroDeflate"}
type CaptureDescription_Encoding string

const (
	CaptureDescription_Encoding_Avro        = CaptureDescription_Encoding("Avro")
	CaptureDescription_Encoding_AvroDeflate = CaptureDescription_Encoding("AvroDeflate")
)

// Mapping from string to CaptureDescription_Encoding
var captureDescription_Encoding_Values = map[string]CaptureDescription_Encoding{
	"avro":        CaptureDescription_Encoding_Avro,
	"avrodeflate": CaptureDescription_Encoding_AvroDeflate,
}

// Capture storage details for capture description
type Destination struct {
	// Name: Name for capture destination
	Name *string `json:"name,omitempty"`

	// Properties: Properties describing the storage account, blob container and archive name format for capture destination
	Properties *Destination_Properties `json:"properties,omitempty"`
}

type Destination_Properties struct {
	// ArchiveNameFormat: Blob naming convention for archive, e.g.
	// {Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}. Here all the parameters
	// (Namespace,EventHub .. etc) are mandatory irrespective of order
	ArchiveNameFormat *string `json:"archiveNameFormat,omitempty"`

	// BlobContainer: Blob container Name
	BlobContainer *string `json:"blobContainer,omitempty"`

	// DataLakeAccountName: The Azure Data Lake Store name for the captured events
	DataLakeAccountName *string `json:"dataLakeAccountName,omitempty"`

	// DataLakeFolderPath: The destination folder path for the captured events
	DataLakeFolderPath *string `json:"dataLakeFolderPath,omitempty"`

	// DataLakeSubscriptionId: Subscription Id of Azure Data Lake Store
	DataLakeSubscriptionId   *string `json:"dataLakeSubscriptionId,omitempty"`
	StorageAccountResourceId *string `json:"storageAccountResourceId,omitempty"`
}
