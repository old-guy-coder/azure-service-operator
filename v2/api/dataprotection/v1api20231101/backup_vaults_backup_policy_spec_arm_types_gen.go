// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20231101

import (
	"encoding/json"
	"github.com/Azure/azure-service-operator/v2/pkg/genruntime"
)

type BackupVaults_BackupPolicy_Spec_ARM struct {
	Name string `json:"name,omitempty"`

	// Properties: BaseBackupPolicyResource properties
	Properties *BaseBackupPolicy_ARM `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &BackupVaults_BackupPolicy_Spec_ARM{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-11-01"
func (policy BackupVaults_BackupPolicy_Spec_ARM) GetAPIVersion() string {
	return string(APIVersion_Value)
}

// GetName returns the Name of the resource
func (policy *BackupVaults_BackupPolicy_Spec_ARM) GetName() string {
	return policy.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.DataProtection/backupVaults/backupPolicies"
func (policy *BackupVaults_BackupPolicy_Spec_ARM) GetType() string {
	return "Microsoft.DataProtection/backupVaults/backupPolicies"
}

type BaseBackupPolicy_ARM struct {
	// BackupPolicy: Mutually exclusive with all other properties
	BackupPolicy *BackupPolicy_ARM `json:"backupPolicy,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BaseBackupPolicy_ARM represents a discriminated union (JSON OneOf)
func (policy BaseBackupPolicy_ARM) MarshalJSON() ([]byte, error) {
	if policy.BackupPolicy != nil {
		return json.Marshal(policy.BackupPolicy)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BaseBackupPolicy_ARM
func (policy *BaseBackupPolicy_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "BackupPolicy" {
		policy.BackupPolicy = &BackupPolicy_ARM{}
		return json.Unmarshal(data, policy.BackupPolicy)
	}

	// No error
	return nil
}

type BackupPolicy_ARM struct {
	// DatasourceTypes: Type of datasource for the backup management
	DatasourceTypes []string                `json:"datasourceTypes,omitempty"`
	ObjectType      BackupPolicy_ObjectType `json:"objectType,omitempty"`

	// PolicyRules: Policy rule dictionary that contains rules for each backuptype i.e Full/Incremental/Logs etc
	PolicyRules []BasePolicyRule_ARM `json:"policyRules,omitempty"`
}

// +kubebuilder:validation:Enum={"BackupPolicy"}
type BackupPolicy_ObjectType string

const BackupPolicy_ObjectType_BackupPolicy = BackupPolicy_ObjectType("BackupPolicy")

// Mapping from string to BackupPolicy_ObjectType
var backupPolicy_ObjectType_Values = map[string]BackupPolicy_ObjectType{
	"backuppolicy": BackupPolicy_ObjectType_BackupPolicy,
}

type BasePolicyRule_ARM struct {
	// AzureBackup: Mutually exclusive with all other properties
	AzureBackup *AzureBackupRule_ARM `json:"azureBackupRule,omitempty"`

	// AzureRetention: Mutually exclusive with all other properties
	AzureRetention *AzureRetentionRule_ARM `json:"azureRetentionRule,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BasePolicyRule_ARM represents a discriminated union (JSON OneOf)
func (rule BasePolicyRule_ARM) MarshalJSON() ([]byte, error) {
	if rule.AzureBackup != nil {
		return json.Marshal(rule.AzureBackup)
	}
	if rule.AzureRetention != nil {
		return json.Marshal(rule.AzureRetention)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BasePolicyRule_ARM
func (rule *BasePolicyRule_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AzureBackupRule" {
		rule.AzureBackup = &AzureBackupRule_ARM{}
		return json.Unmarshal(data, rule.AzureBackup)
	}
	if discriminator == "AzureRetentionRule" {
		rule.AzureRetention = &AzureRetentionRule_ARM{}
		return json.Unmarshal(data, rule.AzureRetention)
	}

	// No error
	return nil
}

type AzureBackupRule_ARM struct {
	BackupParameters *BackupParameters_ARM `json:"backupParameters,omitempty"`

	// DataStore: DataStoreInfo base
	DataStore  *DataStoreInfoBase_ARM     `json:"dataStore,omitempty"`
	Name       *string                    `json:"name,omitempty"`
	ObjectType AzureBackupRule_ObjectType `json:"objectType,omitempty"`
	Trigger    *TriggerContext_ARM        `json:"trigger,omitempty"`
}

type AzureRetentionRule_ARM struct {
	IsDefault  *bool                         `json:"isDefault,omitempty"`
	Lifecycles []SourceLifeCycle_ARM         `json:"lifecycles,omitempty"`
	Name       *string                       `json:"name,omitempty"`
	ObjectType AzureRetentionRule_ObjectType `json:"objectType,omitempty"`
}

// +kubebuilder:validation:Enum={"AzureBackupRule"}
type AzureBackupRule_ObjectType string

const AzureBackupRule_ObjectType_AzureBackupRule = AzureBackupRule_ObjectType("AzureBackupRule")

// Mapping from string to AzureBackupRule_ObjectType
var azureBackupRule_ObjectType_Values = map[string]AzureBackupRule_ObjectType{
	"azurebackuprule": AzureBackupRule_ObjectType_AzureBackupRule,
}

// +kubebuilder:validation:Enum={"AzureRetentionRule"}
type AzureRetentionRule_ObjectType string

const AzureRetentionRule_ObjectType_AzureRetentionRule = AzureRetentionRule_ObjectType("AzureRetentionRule")

// Mapping from string to AzureRetentionRule_ObjectType
var azureRetentionRule_ObjectType_Values = map[string]AzureRetentionRule_ObjectType{
	"azureretentionrule": AzureRetentionRule_ObjectType_AzureRetentionRule,
}

type BackupParameters_ARM struct {
	// AzureBackupParams: Mutually exclusive with all other properties
	AzureBackupParams *AzureBackupParams_ARM `json:"azureBackupParams,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupParameters_ARM represents a discriminated union (JSON OneOf)
func (parameters BackupParameters_ARM) MarshalJSON() ([]byte, error) {
	if parameters.AzureBackupParams != nil {
		return json.Marshal(parameters.AzureBackupParams)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupParameters_ARM
func (parameters *BackupParameters_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AzureBackupParams" {
		parameters.AzureBackupParams = &AzureBackupParams_ARM{}
		return json.Unmarshal(data, parameters.AzureBackupParams)
	}

	// No error
	return nil
}

// DataStoreInfo base
type DataStoreInfoBase_ARM struct {
	// DataStoreType: type of datastore; Operational/Vault/Archive
	DataStoreType *DataStoreInfoBase_DataStoreType `json:"dataStoreType,omitempty"`

	// ObjectType: Type of Datasource object, used to initialize the right inherited type
	ObjectType *string `json:"objectType,omitempty"`
}

// Source LifeCycle
type SourceLifeCycle_ARM struct {
	DeleteAfter *DeleteOption_ARM `json:"deleteAfter,omitempty"`

	// SourceDataStore: DataStoreInfo base
	SourceDataStore             *DataStoreInfoBase_ARM  `json:"sourceDataStore,omitempty"`
	TargetDataStoreCopySettings []TargetCopySetting_ARM `json:"targetDataStoreCopySettings,omitempty"`
}

type TriggerContext_ARM struct {
	// Adhoc: Mutually exclusive with all other properties
	Adhoc *AdhocBasedTriggerContext_ARM `json:"adhocBasedTriggerContext,omitempty"`

	// Schedule: Mutually exclusive with all other properties
	Schedule *ScheduleBasedTriggerContext_ARM `json:"scheduleBasedTriggerContext,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because TriggerContext_ARM represents a discriminated union (JSON OneOf)
func (context TriggerContext_ARM) MarshalJSON() ([]byte, error) {
	if context.Adhoc != nil {
		return json.Marshal(context.Adhoc)
	}
	if context.Schedule != nil {
		return json.Marshal(context.Schedule)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the TriggerContext_ARM
func (context *TriggerContext_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AdhocBasedTriggerContext" {
		context.Adhoc = &AdhocBasedTriggerContext_ARM{}
		return json.Unmarshal(data, context.Adhoc)
	}
	if discriminator == "ScheduleBasedTriggerContext" {
		context.Schedule = &ScheduleBasedTriggerContext_ARM{}
		return json.Unmarshal(data, context.Schedule)
	}

	// No error
	return nil
}

type AdhocBasedTriggerContext_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType AdhocBasedTriggerContext_ObjectType `json:"objectType,omitempty"`

	// TaggingCriteria: Tagging Criteria containing retention tag for adhoc backup.
	TaggingCriteria *AdhocBasedTaggingCriteria_ARM `json:"taggingCriteria,omitempty"`
}

type AzureBackupParams_ARM struct {
	// BackupType: BackupType ; Full/Incremental etc
	BackupType *string `json:"backupType,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType AzureBackupParams_ObjectType `json:"objectType,omitempty"`
}

// +kubebuilder:validation:Enum={"ArchiveStore","OperationalStore","VaultStore"}
type DataStoreInfoBase_DataStoreType string

const (
	DataStoreInfoBase_DataStoreType_ArchiveStore     = DataStoreInfoBase_DataStoreType("ArchiveStore")
	DataStoreInfoBase_DataStoreType_OperationalStore = DataStoreInfoBase_DataStoreType("OperationalStore")
	DataStoreInfoBase_DataStoreType_VaultStore       = DataStoreInfoBase_DataStoreType("VaultStore")
)

// Mapping from string to DataStoreInfoBase_DataStoreType
var dataStoreInfoBase_DataStoreType_Values = map[string]DataStoreInfoBase_DataStoreType{
	"archivestore":     DataStoreInfoBase_DataStoreType_ArchiveStore,
	"operationalstore": DataStoreInfoBase_DataStoreType_OperationalStore,
	"vaultstore":       DataStoreInfoBase_DataStoreType_VaultStore,
}

type DeleteOption_ARM struct {
	// AbsoluteDeleteOption: Mutually exclusive with all other properties
	AbsoluteDeleteOption *AbsoluteDeleteOption_ARM `json:"absoluteDeleteOption,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because DeleteOption_ARM represents a discriminated union (JSON OneOf)
func (option DeleteOption_ARM) MarshalJSON() ([]byte, error) {
	if option.AbsoluteDeleteOption != nil {
		return json.Marshal(option.AbsoluteDeleteOption)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the DeleteOption_ARM
func (option *DeleteOption_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AbsoluteDeleteOption" {
		option.AbsoluteDeleteOption = &AbsoluteDeleteOption_ARM{}
		return json.Unmarshal(data, option.AbsoluteDeleteOption)
	}

	// No error
	return nil
}

type ScheduleBasedTriggerContext_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType ScheduleBasedTriggerContext_ObjectType `json:"objectType,omitempty"`

	// Schedule: Schedule for this backup
	Schedule *BackupSchedule_ARM `json:"schedule,omitempty"`

	// TaggingCriteria: List of tags that can be applicable for given schedule.
	TaggingCriteria []TaggingCriteria_ARM `json:"taggingCriteria,omitempty"`
}

// Target copy settings
type TargetCopySetting_ARM struct {
	// CopyAfter: It can be CustomCopyOption or ImmediateCopyOption.
	CopyAfter *CopyOption_ARM `json:"copyAfter,omitempty"`

	// DataStore: Info of target datastore
	DataStore *DataStoreInfoBase_ARM `json:"dataStore,omitempty"`
}

type AbsoluteDeleteOption_ARM struct {
	// Duration: Duration of deletion after given timespan
	Duration *string `json:"duration,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType AbsoluteDeleteOption_ObjectType `json:"objectType,omitempty"`
}

// Adhoc backup tagging criteria
type AdhocBasedTaggingCriteria_ARM struct {
	// TagInfo: Retention tag information
	TagInfo *RetentionTag_ARM `json:"tagInfo,omitempty"`
}

// +kubebuilder:validation:Enum={"AdhocBasedTriggerContext"}
type AdhocBasedTriggerContext_ObjectType string

const AdhocBasedTriggerContext_ObjectType_AdhocBasedTriggerContext = AdhocBasedTriggerContext_ObjectType("AdhocBasedTriggerContext")

// Mapping from string to AdhocBasedTriggerContext_ObjectType
var adhocBasedTriggerContext_ObjectType_Values = map[string]AdhocBasedTriggerContext_ObjectType{
	"adhocbasedtriggercontext": AdhocBasedTriggerContext_ObjectType_AdhocBasedTriggerContext,
}

// +kubebuilder:validation:Enum={"AzureBackupParams"}
type AzureBackupParams_ObjectType string

const AzureBackupParams_ObjectType_AzureBackupParams = AzureBackupParams_ObjectType("AzureBackupParams")

// Mapping from string to AzureBackupParams_ObjectType
var azureBackupParams_ObjectType_Values = map[string]AzureBackupParams_ObjectType{
	"azurebackupparams": AzureBackupParams_ObjectType_AzureBackupParams,
}

// Schedule for backup
type BackupSchedule_ARM struct {
	// RepeatingTimeIntervals: ISO 8601 repeating time interval format
	RepeatingTimeIntervals []string `json:"repeatingTimeIntervals,omitempty"`

	// TimeZone: Time zone for a schedule. Example: Pacific Standard Time
	TimeZone *string `json:"timeZone,omitempty"`
}

type CopyOption_ARM struct {
	// CopyOnExpiry: Mutually exclusive with all other properties
	CopyOnExpiry *CopyOnExpiryOption_ARM `json:"copyOnExpiryOption,omitempty"`

	// CustomCopy: Mutually exclusive with all other properties
	CustomCopy *CustomCopyOption_ARM `json:"customCopyOption,omitempty"`

	// ImmediateCopy: Mutually exclusive with all other properties
	ImmediateCopy *ImmediateCopyOption_ARM `json:"immediateCopyOption,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because CopyOption_ARM represents a discriminated union (JSON OneOf)
func (option CopyOption_ARM) MarshalJSON() ([]byte, error) {
	if option.CopyOnExpiry != nil {
		return json.Marshal(option.CopyOnExpiry)
	}
	if option.CustomCopy != nil {
		return json.Marshal(option.CustomCopy)
	}
	if option.ImmediateCopy != nil {
		return json.Marshal(option.ImmediateCopy)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the CopyOption_ARM
func (option *CopyOption_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "CopyOnExpiryOption" {
		option.CopyOnExpiry = &CopyOnExpiryOption_ARM{}
		return json.Unmarshal(data, option.CopyOnExpiry)
	}
	if discriminator == "CustomCopyOption" {
		option.CustomCopy = &CustomCopyOption_ARM{}
		return json.Unmarshal(data, option.CustomCopy)
	}
	if discriminator == "ImmediateCopyOption" {
		option.ImmediateCopy = &ImmediateCopyOption_ARM{}
		return json.Unmarshal(data, option.ImmediateCopy)
	}

	// No error
	return nil
}

// +kubebuilder:validation:Enum={"ScheduleBasedTriggerContext"}
type ScheduleBasedTriggerContext_ObjectType string

const ScheduleBasedTriggerContext_ObjectType_ScheduleBasedTriggerContext = ScheduleBasedTriggerContext_ObjectType("ScheduleBasedTriggerContext")

// Mapping from string to ScheduleBasedTriggerContext_ObjectType
var scheduleBasedTriggerContext_ObjectType_Values = map[string]ScheduleBasedTriggerContext_ObjectType{
	"schedulebasedtriggercontext": ScheduleBasedTriggerContext_ObjectType_ScheduleBasedTriggerContext,
}

// Tagging criteria
type TaggingCriteria_ARM struct {
	// Criteria: Criteria which decides whether the tag can be applied to a triggered backup.
	Criteria []BackupCriteria_ARM `json:"criteria,omitempty"`

	// IsDefault: Specifies if tag is default.
	IsDefault *bool `json:"isDefault,omitempty"`

	// TagInfo: Retention tag information
	TagInfo *RetentionTag_ARM `json:"tagInfo,omitempty"`

	// TaggingPriority: Retention Tag priority.
	TaggingPriority *int `json:"taggingPriority,omitempty"`
}

// +kubebuilder:validation:Enum={"AbsoluteDeleteOption"}
type AbsoluteDeleteOption_ObjectType string

const AbsoluteDeleteOption_ObjectType_AbsoluteDeleteOption = AbsoluteDeleteOption_ObjectType("AbsoluteDeleteOption")

// Mapping from string to AbsoluteDeleteOption_ObjectType
var absoluteDeleteOption_ObjectType_Values = map[string]AbsoluteDeleteOption_ObjectType{
	"absolutedeleteoption": AbsoluteDeleteOption_ObjectType_AbsoluteDeleteOption,
}

type BackupCriteria_ARM struct {
	// ScheduleBasedBackupCriteria: Mutually exclusive with all other properties
	ScheduleBasedBackupCriteria *ScheduleBasedBackupCriteria_ARM `json:"scheduleBasedBackupCriteria,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupCriteria_ARM represents a discriminated union (JSON OneOf)
func (criteria BackupCriteria_ARM) MarshalJSON() ([]byte, error) {
	if criteria.ScheduleBasedBackupCriteria != nil {
		return json.Marshal(criteria.ScheduleBasedBackupCriteria)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupCriteria_ARM
func (criteria *BackupCriteria_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "ScheduleBasedBackupCriteria" {
		criteria.ScheduleBasedBackupCriteria = &ScheduleBasedBackupCriteria_ARM{}
		return json.Unmarshal(data, criteria.ScheduleBasedBackupCriteria)
	}

	// No error
	return nil
}

type CopyOnExpiryOption_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType CopyOnExpiryOption_ObjectType `json:"objectType,omitempty"`
}

type CustomCopyOption_ARM struct {
	// Duration: Data copied after given timespan
	Duration *string `json:"duration,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType CustomCopyOption_ObjectType `json:"objectType,omitempty"`
}

type ImmediateCopyOption_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType ImmediateCopyOption_ObjectType `json:"objectType,omitempty"`
}

// Retention tag
type RetentionTag_ARM struct {
	// TagName: Retention Tag Name to relate it to retention rule.
	TagName *string `json:"tagName,omitempty"`
}

// +kubebuilder:validation:Enum={"CopyOnExpiryOption"}
type CopyOnExpiryOption_ObjectType string

const CopyOnExpiryOption_ObjectType_CopyOnExpiryOption = CopyOnExpiryOption_ObjectType("CopyOnExpiryOption")

// Mapping from string to CopyOnExpiryOption_ObjectType
var copyOnExpiryOption_ObjectType_Values = map[string]CopyOnExpiryOption_ObjectType{
	"copyonexpiryoption": CopyOnExpiryOption_ObjectType_CopyOnExpiryOption,
}

// +kubebuilder:validation:Enum={"CustomCopyOption"}
type CustomCopyOption_ObjectType string

const CustomCopyOption_ObjectType_CustomCopyOption = CustomCopyOption_ObjectType("CustomCopyOption")

// Mapping from string to CustomCopyOption_ObjectType
var customCopyOption_ObjectType_Values = map[string]CustomCopyOption_ObjectType{
	"customcopyoption": CustomCopyOption_ObjectType_CustomCopyOption,
}

// +kubebuilder:validation:Enum={"ImmediateCopyOption"}
type ImmediateCopyOption_ObjectType string

const ImmediateCopyOption_ObjectType_ImmediateCopyOption = ImmediateCopyOption_ObjectType("ImmediateCopyOption")

// Mapping from string to ImmediateCopyOption_ObjectType
var immediateCopyOption_ObjectType_Values = map[string]ImmediateCopyOption_ObjectType{
	"immediatecopyoption": ImmediateCopyOption_ObjectType_ImmediateCopyOption,
}

type ScheduleBasedBackupCriteria_ARM struct {
	// AbsoluteCriteria: it contains absolute values like "AllBackup" / "FirstOfDay" / "FirstOfWeek" / "FirstOfMonth"
	// and should be part of AbsoluteMarker enum
	AbsoluteCriteria []ScheduleBasedBackupCriteria_AbsoluteCriteria `json:"absoluteCriteria,omitempty"`

	// DaysOfMonth: This is day of the month from 1 to 28 other wise last of month
	DaysOfMonth []Day_ARM `json:"daysOfMonth,omitempty"`

	// DaysOfTheWeek: It should be Sunday/Monday/T..../Saturday
	DaysOfTheWeek []ScheduleBasedBackupCriteria_DaysOfTheWeek `json:"daysOfTheWeek,omitempty"`

	// MonthsOfYear: It should be January/February/....../December
	MonthsOfYear []ScheduleBasedBackupCriteria_MonthsOfYear `json:"monthsOfYear,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType ScheduleBasedBackupCriteria_ObjectType `json:"objectType,omitempty"`

	// ScheduleTimes: List of schedule times for backup
	ScheduleTimes []string `json:"scheduleTimes,omitempty"`

	// WeeksOfTheMonth: It should be First/Second/Third/Fourth/Last
	WeeksOfTheMonth []ScheduleBasedBackupCriteria_WeeksOfTheMonth `json:"weeksOfTheMonth,omitempty"`
}

// Day of the week
type Day_ARM struct {
	// Date: Date of the month
	Date *int `json:"date,omitempty"`

	// IsLast: Whether Date is last date of month
	IsLast *bool `json:"isLast,omitempty"`
}

// +kubebuilder:validation:Enum={"AllBackup","FirstOfDay","FirstOfMonth","FirstOfWeek","FirstOfYear"}
type ScheduleBasedBackupCriteria_AbsoluteCriteria string

const (
	ScheduleBasedBackupCriteria_AbsoluteCriteria_AllBackup    = ScheduleBasedBackupCriteria_AbsoluteCriteria("AllBackup")
	ScheduleBasedBackupCriteria_AbsoluteCriteria_FirstOfDay   = ScheduleBasedBackupCriteria_AbsoluteCriteria("FirstOfDay")
	ScheduleBasedBackupCriteria_AbsoluteCriteria_FirstOfMonth = ScheduleBasedBackupCriteria_AbsoluteCriteria("FirstOfMonth")
	ScheduleBasedBackupCriteria_AbsoluteCriteria_FirstOfWeek  = ScheduleBasedBackupCriteria_AbsoluteCriteria("FirstOfWeek")
	ScheduleBasedBackupCriteria_AbsoluteCriteria_FirstOfYear  = ScheduleBasedBackupCriteria_AbsoluteCriteria("FirstOfYear")
)

// Mapping from string to ScheduleBasedBackupCriteria_AbsoluteCriteria
var scheduleBasedBackupCriteria_AbsoluteCriteria_Values = map[string]ScheduleBasedBackupCriteria_AbsoluteCriteria{
	"allbackup":    ScheduleBasedBackupCriteria_AbsoluteCriteria_AllBackup,
	"firstofday":   ScheduleBasedBackupCriteria_AbsoluteCriteria_FirstOfDay,
	"firstofmonth": ScheduleBasedBackupCriteria_AbsoluteCriteria_FirstOfMonth,
	"firstofweek":  ScheduleBasedBackupCriteria_AbsoluteCriteria_FirstOfWeek,
	"firstofyear":  ScheduleBasedBackupCriteria_AbsoluteCriteria_FirstOfYear,
}

// +kubebuilder:validation:Enum={"Friday","Monday","Saturday","Sunday","Thursday","Tuesday","Wednesday"}
type ScheduleBasedBackupCriteria_DaysOfTheWeek string

const (
	ScheduleBasedBackupCriteria_DaysOfTheWeek_Friday    = ScheduleBasedBackupCriteria_DaysOfTheWeek("Friday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_Monday    = ScheduleBasedBackupCriteria_DaysOfTheWeek("Monday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_Saturday  = ScheduleBasedBackupCriteria_DaysOfTheWeek("Saturday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_Sunday    = ScheduleBasedBackupCriteria_DaysOfTheWeek("Sunday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_Thursday  = ScheduleBasedBackupCriteria_DaysOfTheWeek("Thursday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_Tuesday   = ScheduleBasedBackupCriteria_DaysOfTheWeek("Tuesday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_Wednesday = ScheduleBasedBackupCriteria_DaysOfTheWeek("Wednesday")
)

// Mapping from string to ScheduleBasedBackupCriteria_DaysOfTheWeek
var scheduleBasedBackupCriteria_DaysOfTheWeek_Values = map[string]ScheduleBasedBackupCriteria_DaysOfTheWeek{
	"friday":    ScheduleBasedBackupCriteria_DaysOfTheWeek_Friday,
	"monday":    ScheduleBasedBackupCriteria_DaysOfTheWeek_Monday,
	"saturday":  ScheduleBasedBackupCriteria_DaysOfTheWeek_Saturday,
	"sunday":    ScheduleBasedBackupCriteria_DaysOfTheWeek_Sunday,
	"thursday":  ScheduleBasedBackupCriteria_DaysOfTheWeek_Thursday,
	"tuesday":   ScheduleBasedBackupCriteria_DaysOfTheWeek_Tuesday,
	"wednesday": ScheduleBasedBackupCriteria_DaysOfTheWeek_Wednesday,
}

// +kubebuilder:validation:Enum={"April","August","December","February","January","July","June","March","May","November","October","September"}
type ScheduleBasedBackupCriteria_MonthsOfYear string

const (
	ScheduleBasedBackupCriteria_MonthsOfYear_April     = ScheduleBasedBackupCriteria_MonthsOfYear("April")
	ScheduleBasedBackupCriteria_MonthsOfYear_August    = ScheduleBasedBackupCriteria_MonthsOfYear("August")
	ScheduleBasedBackupCriteria_MonthsOfYear_December  = ScheduleBasedBackupCriteria_MonthsOfYear("December")
	ScheduleBasedBackupCriteria_MonthsOfYear_February  = ScheduleBasedBackupCriteria_MonthsOfYear("February")
	ScheduleBasedBackupCriteria_MonthsOfYear_January   = ScheduleBasedBackupCriteria_MonthsOfYear("January")
	ScheduleBasedBackupCriteria_MonthsOfYear_July      = ScheduleBasedBackupCriteria_MonthsOfYear("July")
	ScheduleBasedBackupCriteria_MonthsOfYear_June      = ScheduleBasedBackupCriteria_MonthsOfYear("June")
	ScheduleBasedBackupCriteria_MonthsOfYear_March     = ScheduleBasedBackupCriteria_MonthsOfYear("March")
	ScheduleBasedBackupCriteria_MonthsOfYear_May       = ScheduleBasedBackupCriteria_MonthsOfYear("May")
	ScheduleBasedBackupCriteria_MonthsOfYear_November  = ScheduleBasedBackupCriteria_MonthsOfYear("November")
	ScheduleBasedBackupCriteria_MonthsOfYear_October   = ScheduleBasedBackupCriteria_MonthsOfYear("October")
	ScheduleBasedBackupCriteria_MonthsOfYear_September = ScheduleBasedBackupCriteria_MonthsOfYear("September")
)

// Mapping from string to ScheduleBasedBackupCriteria_MonthsOfYear
var scheduleBasedBackupCriteria_MonthsOfYear_Values = map[string]ScheduleBasedBackupCriteria_MonthsOfYear{
	"april":     ScheduleBasedBackupCriteria_MonthsOfYear_April,
	"august":    ScheduleBasedBackupCriteria_MonthsOfYear_August,
	"december":  ScheduleBasedBackupCriteria_MonthsOfYear_December,
	"february":  ScheduleBasedBackupCriteria_MonthsOfYear_February,
	"january":   ScheduleBasedBackupCriteria_MonthsOfYear_January,
	"july":      ScheduleBasedBackupCriteria_MonthsOfYear_July,
	"june":      ScheduleBasedBackupCriteria_MonthsOfYear_June,
	"march":     ScheduleBasedBackupCriteria_MonthsOfYear_March,
	"may":       ScheduleBasedBackupCriteria_MonthsOfYear_May,
	"november":  ScheduleBasedBackupCriteria_MonthsOfYear_November,
	"october":   ScheduleBasedBackupCriteria_MonthsOfYear_October,
	"september": ScheduleBasedBackupCriteria_MonthsOfYear_September,
}

// +kubebuilder:validation:Enum={"ScheduleBasedBackupCriteria"}
type ScheduleBasedBackupCriteria_ObjectType string

const ScheduleBasedBackupCriteria_ObjectType_ScheduleBasedBackupCriteria = ScheduleBasedBackupCriteria_ObjectType("ScheduleBasedBackupCriteria")

// Mapping from string to ScheduleBasedBackupCriteria_ObjectType
var scheduleBasedBackupCriteria_ObjectType_Values = map[string]ScheduleBasedBackupCriteria_ObjectType{
	"schedulebasedbackupcriteria": ScheduleBasedBackupCriteria_ObjectType_ScheduleBasedBackupCriteria,
}

// +kubebuilder:validation:Enum={"First","Fourth","Last","Second","Third"}
type ScheduleBasedBackupCriteria_WeeksOfTheMonth string

const (
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_First  = ScheduleBasedBackupCriteria_WeeksOfTheMonth("First")
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_Fourth = ScheduleBasedBackupCriteria_WeeksOfTheMonth("Fourth")
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_Last   = ScheduleBasedBackupCriteria_WeeksOfTheMonth("Last")
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_Second = ScheduleBasedBackupCriteria_WeeksOfTheMonth("Second")
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_Third  = ScheduleBasedBackupCriteria_WeeksOfTheMonth("Third")
)

// Mapping from string to ScheduleBasedBackupCriteria_WeeksOfTheMonth
var scheduleBasedBackupCriteria_WeeksOfTheMonth_Values = map[string]ScheduleBasedBackupCriteria_WeeksOfTheMonth{
	"first":  ScheduleBasedBackupCriteria_WeeksOfTheMonth_First,
	"fourth": ScheduleBasedBackupCriteria_WeeksOfTheMonth_Fourth,
	"last":   ScheduleBasedBackupCriteria_WeeksOfTheMonth_Last,
	"second": ScheduleBasedBackupCriteria_WeeksOfTheMonth_Second,
	"third":  ScheduleBasedBackupCriteria_WeeksOfTheMonth_Third,
}
