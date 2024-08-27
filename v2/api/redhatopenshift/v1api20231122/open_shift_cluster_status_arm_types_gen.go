// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231122

// OpenShiftCluster represents an Azure Red Hat OpenShift cluster.
type OpenShiftCluster_STATUS_ARM struct {
	// Id: Fully qualified resource ID for the resource. Ex -
	// /subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}
	Id *string `json:"id,omitempty"`

	// Location: The geo-location where the resource lives
	Location *string `json:"location,omitempty"`

	// Name: The name of the resource
	Name *string `json:"name,omitempty"`

	// Properties: The cluster properties.
	Properties *OpenShiftClusterProperties_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Azure Resource Manager metadata containing createdBy and modifiedBy information.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Tags: Resource tags.
	Tags map[string]string `json:"tags,omitempty"`

	// Type: The type of the resource. E.g. "Microsoft.Compute/virtualMachines" or "Microsoft.Storage/storageAccounts"
	Type *string `json:"type,omitempty"`
}

// OpenShiftClusterProperties represents an OpenShift cluster's properties.
type OpenShiftClusterProperties_STATUS_ARM struct {
	// ApiserverProfile: The cluster API server profile.
	ApiserverProfile *APIServerProfile_STATUS_ARM `json:"apiserverProfile,omitempty"`

	// ClusterProfile: The cluster profile.
	ClusterProfile *ClusterProfile_STATUS_ARM `json:"clusterProfile,omitempty"`

	// ConsoleProfile: The console profile.
	ConsoleProfile *ConsoleProfile_STATUS_ARM `json:"consoleProfile,omitempty"`

	// IngressProfiles: The cluster ingress profiles.
	IngressProfiles []IngressProfile_STATUS_ARM `json:"ingressProfiles,omitempty"`

	// MasterProfile: The cluster master profile.
	MasterProfile *MasterProfile_STATUS_ARM `json:"masterProfile,omitempty"`

	// NetworkProfile: The cluster network profile.
	NetworkProfile *NetworkProfile_STATUS_ARM `json:"networkProfile,omitempty"`

	// ProvisioningState: The cluster provisioning state.
	ProvisioningState *ProvisioningState_STATUS `json:"provisioningState,omitempty"`

	// ServicePrincipalProfile: The cluster service principal profile.
	ServicePrincipalProfile *ServicePrincipalProfile_STATUS_ARM `json:"servicePrincipalProfile,omitempty"`

	// WorkerProfiles: The cluster worker profiles.
	WorkerProfiles []WorkerProfile_STATUS_ARM `json:"workerProfiles,omitempty"`

	// WorkerProfilesStatus: The cluster worker profiles status.
	WorkerProfilesStatus []WorkerProfile_STATUS_ARM `json:"workerProfilesStatus,omitempty"`
}

// Metadata pertaining to creation and last modification of the resource.
type SystemData_STATUS_ARM struct {
	// CreatedAt: The timestamp of resource creation (UTC).
	CreatedAt *string `json:"createdAt,omitempty"`

	// CreatedBy: The identity that created the resource.
	CreatedBy *string `json:"createdBy,omitempty"`

	// CreatedByType: The type of identity that created the resource.
	CreatedByType *SystemData_CreatedByType_STATUS `json:"createdByType,omitempty"`

	// LastModifiedAt: The timestamp of resource last modification (UTC)
	LastModifiedAt *string `json:"lastModifiedAt,omitempty"`

	// LastModifiedBy: The identity that last modified the resource.
	LastModifiedBy *string `json:"lastModifiedBy,omitempty"`

	// LastModifiedByType: The type of identity that last modified the resource.
	LastModifiedByType *SystemData_LastModifiedByType_STATUS `json:"lastModifiedByType,omitempty"`
}

// APIServerProfile represents an API server profile.
type APIServerProfile_STATUS_ARM struct {
	// Ip: The IP of the cluster API server.
	Ip *string `json:"ip,omitempty"`

	// Url: The URL to access the cluster API server.
	Url *string `json:"url,omitempty"`

	// Visibility: API server visibility.
	Visibility *Visibility_STATUS `json:"visibility,omitempty"`
}

// ClusterProfile represents a cluster profile.
type ClusterProfile_STATUS_ARM struct {
	// Domain: The domain for the cluster.
	Domain *string `json:"domain,omitempty"`

	// FipsValidatedModules: If FIPS validated crypto modules are used
	FipsValidatedModules *FipsValidatedModules_STATUS `json:"fipsValidatedModules,omitempty"`

	// ResourceGroupId: The ID of the cluster resource group.
	ResourceGroupId *string `json:"resourceGroupId,omitempty"`

	// Version: The version of the cluster.
	Version *string `json:"version,omitempty"`
}

// ConsoleProfile represents a console profile.
type ConsoleProfile_STATUS_ARM struct {
	// Url: The URL to access the cluster console.
	Url *string `json:"url,omitempty"`
}

// IngressProfile represents an ingress profile.
type IngressProfile_STATUS_ARM struct {
	// Ip: The IP of the ingress.
	Ip *string `json:"ip,omitempty"`

	// Name: The ingress profile name.
	Name *string `json:"name,omitempty"`

	// Visibility: Ingress visibility.
	Visibility *Visibility_STATUS `json:"visibility,omitempty"`
}

// MasterProfile represents a master profile.
type MasterProfile_STATUS_ARM struct {
	// DiskEncryptionSetId: The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetId *string `json:"diskEncryptionSetId,omitempty"`

	// EncryptionAtHost: Whether master virtual machines are encrypted at host.
	EncryptionAtHost *EncryptionAtHost_STATUS `json:"encryptionAtHost,omitempty"`

	// SubnetId: The Azure resource ID of the master subnet.
	SubnetId *string `json:"subnetId,omitempty"`

	// VmSize: The size of the master VMs.
	VmSize *string `json:"vmSize,omitempty"`
}

// NetworkProfile represents a network profile.
type NetworkProfile_STATUS_ARM struct {
	// LoadBalancerProfile: The cluster load balancer profile.
	LoadBalancerProfile *LoadBalancerProfile_STATUS_ARM `json:"loadBalancerProfile,omitempty"`

	// OutboundType: The OutboundType used for egress traffic.
	OutboundType *OutboundType_STATUS `json:"outboundType,omitempty"`

	// PodCidr: The CIDR used for OpenShift/Kubernetes Pods.
	PodCidr *string `json:"podCidr,omitempty"`

	// PreconfiguredNSG: Specifies whether subnets are pre-attached with an NSG
	PreconfiguredNSG *PreconfiguredNSG_STATUS `json:"preconfiguredNSG,omitempty"`

	// ServiceCidr: The CIDR used for OpenShift/Kubernetes Services.
	ServiceCidr *string `json:"serviceCidr,omitempty"`
}

// ServicePrincipalProfile represents a service principal profile.
type ServicePrincipalProfile_STATUS_ARM struct {
	// ClientId: The client ID used for the cluster.
	ClientId *string `json:"clientId,omitempty"`
}

type SystemData_CreatedByType_STATUS string

const (
	SystemData_CreatedByType_STATUS_Application     = SystemData_CreatedByType_STATUS("Application")
	SystemData_CreatedByType_STATUS_Key             = SystemData_CreatedByType_STATUS("Key")
	SystemData_CreatedByType_STATUS_ManagedIdentity = SystemData_CreatedByType_STATUS("ManagedIdentity")
	SystemData_CreatedByType_STATUS_User            = SystemData_CreatedByType_STATUS("User")
)

// Mapping from string to SystemData_CreatedByType_STATUS
var systemData_CreatedByType_STATUS_Values = map[string]SystemData_CreatedByType_STATUS{
	"application":     SystemData_CreatedByType_STATUS_Application,
	"key":             SystemData_CreatedByType_STATUS_Key,
	"managedidentity": SystemData_CreatedByType_STATUS_ManagedIdentity,
	"user":            SystemData_CreatedByType_STATUS_User,
}

type SystemData_LastModifiedByType_STATUS string

const (
	SystemData_LastModifiedByType_STATUS_Application     = SystemData_LastModifiedByType_STATUS("Application")
	SystemData_LastModifiedByType_STATUS_Key             = SystemData_LastModifiedByType_STATUS("Key")
	SystemData_LastModifiedByType_STATUS_ManagedIdentity = SystemData_LastModifiedByType_STATUS("ManagedIdentity")
	SystemData_LastModifiedByType_STATUS_User            = SystemData_LastModifiedByType_STATUS("User")
)

// Mapping from string to SystemData_LastModifiedByType_STATUS
var systemData_LastModifiedByType_STATUS_Values = map[string]SystemData_LastModifiedByType_STATUS{
	"application":     SystemData_LastModifiedByType_STATUS_Application,
	"key":             SystemData_LastModifiedByType_STATUS_Key,
	"managedidentity": SystemData_LastModifiedByType_STATUS_ManagedIdentity,
	"user":            SystemData_LastModifiedByType_STATUS_User,
}

// WorkerProfile represents a worker profile.
type WorkerProfile_STATUS_ARM struct {
	// Count: The number of worker VMs.
	Count *int `json:"count,omitempty"`

	// DiskEncryptionSetId: The resource ID of an associated DiskEncryptionSet, if applicable.
	DiskEncryptionSetId *string `json:"diskEncryptionSetId,omitempty"`

	// DiskSizeGB: The disk size of the worker VMs.
	DiskSizeGB *int `json:"diskSizeGB,omitempty"`

	// EncryptionAtHost: Whether master virtual machines are encrypted at host.
	EncryptionAtHost *EncryptionAtHost_STATUS `json:"encryptionAtHost,omitempty"`

	// Name: The worker profile name.
	Name *string `json:"name,omitempty"`

	// SubnetId: The Azure resource ID of the worker subnet.
	SubnetId *string `json:"subnetId,omitempty"`

	// VmSize: The size of the worker VMs.
	VmSize *string `json:"vmSize,omitempty"`
}

// LoadBalancerProfile represents the profile of the cluster public load balancer.
type LoadBalancerProfile_STATUS_ARM struct {
	// EffectiveOutboundIps: The list of effective outbound IP addresses of the public load balancer.
	EffectiveOutboundIps []EffectiveOutboundIP_STATUS_ARM `json:"effectiveOutboundIps,omitempty"`

	// ManagedOutboundIps: The desired managed outbound IPs for the cluster public load balancer.
	ManagedOutboundIps *ManagedOutboundIPs_STATUS_ARM `json:"managedOutboundIps,omitempty"`
}

// EffectiveOutboundIP represents an effective outbound IP resource of the cluster public load balancer.
type EffectiveOutboundIP_STATUS_ARM struct {
	// Id: The fully qualified Azure resource id of an IP address resource.
	Id *string `json:"id,omitempty"`
}

// ManagedOutboundIPs represents the desired managed outbound IPs for the cluster public load balancer.
type ManagedOutboundIPs_STATUS_ARM struct {
	// Count: Count represents the desired number of IPv4 outbound IPs created and managed by Azure for the cluster public load
	// balancer.  Allowed values are in the range of 1 - 20.  The default value is 1.
	Count *int `json:"count,omitempty"`
}