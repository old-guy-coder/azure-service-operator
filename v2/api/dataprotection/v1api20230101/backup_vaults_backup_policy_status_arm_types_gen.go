// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230101

import "encoding/json"

type BackupVaultsBackupPolicy_STATUS_ARM struct {
	// Id: Resource Id represents the complete path to the resource.
	Id *string `json:"id,omitempty"`

	// Name: Resource name associated with the resource.
	Name *string `json:"name,omitempty"`

	// Properties: BaseBackupPolicyResource properties
	Properties *BaseBackupPolicy_STATUS_ARM `json:"properties,omitempty"`

	// SystemData: Metadata pertaining to creation and last modification of the resource.
	SystemData *SystemData_STATUS_ARM `json:"systemData,omitempty"`

	// Type: Resource type represents the complete path of the form Namespace/ResourceType/ResourceType/...
	Type *string `json:"type,omitempty"`
}

type BaseBackupPolicy_STATUS_ARM struct {
	// BackupPolicy: Mutually exclusive with all other properties
	BackupPolicy *BackupPolicy_STATUS_ARM `json:"backupPolicy,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BaseBackupPolicy_STATUS_ARM represents a discriminated union (JSON OneOf)
func (policy BaseBackupPolicy_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if policy.BackupPolicy != nil {
		return json.Marshal(policy.BackupPolicy)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BaseBackupPolicy_STATUS_ARM
func (policy *BaseBackupPolicy_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "BackupPolicy" {
		policy.BackupPolicy = &BackupPolicy_STATUS_ARM{}
		return json.Unmarshal(data, policy.BackupPolicy)
	}

	// No error
	return nil
}

type BackupPolicy_STATUS_ARM struct {
	// DatasourceTypes: Type of datasource for the backup management
	DatasourceTypes []string                           `json:"datasourceTypes,omitempty"`
	ObjectType      BackupPolicy_ObjectType_STATUS_ARM `json:"objectType,omitempty"`

	// PolicyRules: Policy rule dictionary that contains rules for each backuptype i.e Full/Incremental/Logs etc
	PolicyRules []BasePolicyRule_STATUS_ARM `json:"policyRules,omitempty"`
}

type BackupPolicy_ObjectType_STATUS_ARM string

const BackupPolicy_ObjectType_STATUS_ARM_BackupPolicy = BackupPolicy_ObjectType_STATUS_ARM("BackupPolicy")

// Mapping from string to BackupPolicy_ObjectType_STATUS_ARM
var backupPolicy_ObjectType_STATUS_ARM_Values = map[string]BackupPolicy_ObjectType_STATUS_ARM{
	"backuppolicy": BackupPolicy_ObjectType_STATUS_ARM_BackupPolicy,
}

type BasePolicyRule_STATUS_ARM struct {
	// AzureBackup: Mutually exclusive with all other properties
	AzureBackup *AzureBackupRule_STATUS_ARM `json:"azureBackupRule,omitempty"`

	// AzureRetention: Mutually exclusive with all other properties
	AzureRetention *AzureRetentionRule_STATUS_ARM `json:"azureRetentionRule,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BasePolicyRule_STATUS_ARM represents a discriminated union (JSON OneOf)
func (rule BasePolicyRule_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if rule.AzureBackup != nil {
		return json.Marshal(rule.AzureBackup)
	}
	if rule.AzureRetention != nil {
		return json.Marshal(rule.AzureRetention)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BasePolicyRule_STATUS_ARM
func (rule *BasePolicyRule_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AzureBackupRule" {
		rule.AzureBackup = &AzureBackupRule_STATUS_ARM{}
		return json.Unmarshal(data, rule.AzureBackup)
	}
	if discriminator == "AzureRetentionRule" {
		rule.AzureRetention = &AzureRetentionRule_STATUS_ARM{}
		return json.Unmarshal(data, rule.AzureRetention)
	}

	// No error
	return nil
}

type AzureBackupRule_STATUS_ARM struct {
	BackupParameters *BackupParameters_STATUS_ARM `json:"backupParameters,omitempty"`

	// DataStore: DataStoreInfo base
	DataStore  *DataStoreInfoBase_STATUS_ARM         `json:"dataStore,omitempty"`
	Name       *string                               `json:"name,omitempty"`
	ObjectType AzureBackupRule_ObjectType_STATUS_ARM `json:"objectType,omitempty"`
	Trigger    *TriggerContext_STATUS_ARM            `json:"trigger,omitempty"`
}

type AzureRetentionRule_STATUS_ARM struct {
	IsDefault  *bool                                    `json:"isDefault,omitempty"`
	Lifecycles []SourceLifeCycle_STATUS_ARM             `json:"lifecycles,omitempty"`
	Name       *string                                  `json:"name,omitempty"`
	ObjectType AzureRetentionRule_ObjectType_STATUS_ARM `json:"objectType,omitempty"`
}

type AzureBackupRule_ObjectType_STATUS_ARM string

const AzureBackupRule_ObjectType_STATUS_ARM_AzureBackupRule = AzureBackupRule_ObjectType_STATUS_ARM("AzureBackupRule")

// Mapping from string to AzureBackupRule_ObjectType_STATUS_ARM
var azureBackupRule_ObjectType_STATUS_ARM_Values = map[string]AzureBackupRule_ObjectType_STATUS_ARM{
	"azurebackuprule": AzureBackupRule_ObjectType_STATUS_ARM_AzureBackupRule,
}

type AzureRetentionRule_ObjectType_STATUS_ARM string

const AzureRetentionRule_ObjectType_STATUS_ARM_AzureRetentionRule = AzureRetentionRule_ObjectType_STATUS_ARM("AzureRetentionRule")

// Mapping from string to AzureRetentionRule_ObjectType_STATUS_ARM
var azureRetentionRule_ObjectType_STATUS_ARM_Values = map[string]AzureRetentionRule_ObjectType_STATUS_ARM{
	"azureretentionrule": AzureRetentionRule_ObjectType_STATUS_ARM_AzureRetentionRule,
}

type BackupParameters_STATUS_ARM struct {
	// AzureBackupParams: Mutually exclusive with all other properties
	AzureBackupParams *AzureBackupParams_STATUS_ARM `json:"azureBackupParams,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupParameters_STATUS_ARM represents a discriminated union (JSON OneOf)
func (parameters BackupParameters_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if parameters.AzureBackupParams != nil {
		return json.Marshal(parameters.AzureBackupParams)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupParameters_STATUS_ARM
func (parameters *BackupParameters_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AzureBackupParams" {
		parameters.AzureBackupParams = &AzureBackupParams_STATUS_ARM{}
		return json.Unmarshal(data, parameters.AzureBackupParams)
	}

	// No error
	return nil
}

// DataStoreInfo base
type DataStoreInfoBase_STATUS_ARM struct {
	// DataStoreType: type of datastore; Operational/Vault/Archive
	DataStoreType *DataStoreInfoBase_DataStoreType_STATUS_ARM `json:"dataStoreType,omitempty"`

	// ObjectType: Type of Datasource object, used to initialize the right inherited type
	ObjectType *string `json:"objectType,omitempty"`
}

// Source LifeCycle
type SourceLifeCycle_STATUS_ARM struct {
	DeleteAfter *DeleteOption_STATUS_ARM `json:"deleteAfter,omitempty"`

	// SourceDataStore: DataStoreInfo base
	SourceDataStore             *DataStoreInfoBase_STATUS_ARM  `json:"sourceDataStore,omitempty"`
	TargetDataStoreCopySettings []TargetCopySetting_STATUS_ARM `json:"targetDataStoreCopySettings,omitempty"`
}

type TriggerContext_STATUS_ARM struct {
	// Adhoc: Mutually exclusive with all other properties
	Adhoc *AdhocBasedTriggerContext_STATUS_ARM `json:"adhocBasedTriggerContext,omitempty"`

	// Schedule: Mutually exclusive with all other properties
	Schedule *ScheduleBasedTriggerContext_STATUS_ARM `json:"scheduleBasedTriggerContext,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because TriggerContext_STATUS_ARM represents a discriminated union (JSON OneOf)
func (context TriggerContext_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if context.Adhoc != nil {
		return json.Marshal(context.Adhoc)
	}
	if context.Schedule != nil {
		return json.Marshal(context.Schedule)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the TriggerContext_STATUS_ARM
func (context *TriggerContext_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AdhocBasedTriggerContext" {
		context.Adhoc = &AdhocBasedTriggerContext_STATUS_ARM{}
		return json.Unmarshal(data, context.Adhoc)
	}
	if discriminator == "ScheduleBasedTriggerContext" {
		context.Schedule = &ScheduleBasedTriggerContext_STATUS_ARM{}
		return json.Unmarshal(data, context.Schedule)
	}

	// No error
	return nil
}

type AdhocBasedTriggerContext_STATUS_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType AdhocBasedTriggerContext_ObjectType_STATUS_ARM `json:"objectType,omitempty"`

	// TaggingCriteria: Tagging Criteria containing retention tag for adhoc backup.
	TaggingCriteria *AdhocBasedTaggingCriteria_STATUS_ARM `json:"taggingCriteria,omitempty"`
}

type AzureBackupParams_STATUS_ARM struct {
	// BackupType: BackupType ; Full/Incremental etc
	BackupType *string `json:"backupType,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType AzureBackupParams_ObjectType_STATUS_ARM `json:"objectType,omitempty"`
}

type DataStoreInfoBase_DataStoreType_STATUS_ARM string

const (
	DataStoreInfoBase_DataStoreType_STATUS_ARM_ArchiveStore     = DataStoreInfoBase_DataStoreType_STATUS_ARM("ArchiveStore")
	DataStoreInfoBase_DataStoreType_STATUS_ARM_OperationalStore = DataStoreInfoBase_DataStoreType_STATUS_ARM("OperationalStore")
	DataStoreInfoBase_DataStoreType_STATUS_ARM_VaultStore       = DataStoreInfoBase_DataStoreType_STATUS_ARM("VaultStore")
)

// Mapping from string to DataStoreInfoBase_DataStoreType_STATUS_ARM
var dataStoreInfoBase_DataStoreType_STATUS_ARM_Values = map[string]DataStoreInfoBase_DataStoreType_STATUS_ARM{
	"archivestore":     DataStoreInfoBase_DataStoreType_STATUS_ARM_ArchiveStore,
	"operationalstore": DataStoreInfoBase_DataStoreType_STATUS_ARM_OperationalStore,
	"vaultstore":       DataStoreInfoBase_DataStoreType_STATUS_ARM_VaultStore,
}

type DeleteOption_STATUS_ARM struct {
	// AbsoluteDeleteOption: Mutually exclusive with all other properties
	AbsoluteDeleteOption *AbsoluteDeleteOption_STATUS_ARM `json:"absoluteDeleteOption,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because DeleteOption_STATUS_ARM represents a discriminated union (JSON OneOf)
func (option DeleteOption_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if option.AbsoluteDeleteOption != nil {
		return json.Marshal(option.AbsoluteDeleteOption)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the DeleteOption_STATUS_ARM
func (option *DeleteOption_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "AbsoluteDeleteOption" {
		option.AbsoluteDeleteOption = &AbsoluteDeleteOption_STATUS_ARM{}
		return json.Unmarshal(data, option.AbsoluteDeleteOption)
	}

	// No error
	return nil
}

type ScheduleBasedTriggerContext_STATUS_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType ScheduleBasedTriggerContext_ObjectType_STATUS_ARM `json:"objectType,omitempty"`

	// Schedule: Schedule for this backup
	Schedule *BackupSchedule_STATUS_ARM `json:"schedule,omitempty"`

	// TaggingCriteria: List of tags that can be applicable for given schedule.
	TaggingCriteria []TaggingCriteria_STATUS_ARM `json:"taggingCriteria,omitempty"`
}

// Target copy settings
type TargetCopySetting_STATUS_ARM struct {
	// CopyAfter: It can be CustomCopyOption or ImmediateCopyOption.
	CopyAfter *CopyOption_STATUS_ARM `json:"copyAfter,omitempty"`

	// DataStore: Info of target datastore
	DataStore *DataStoreInfoBase_STATUS_ARM `json:"dataStore,omitempty"`
}

type AbsoluteDeleteOption_STATUS_ARM struct {
	// Duration: Duration of deletion after given timespan
	Duration *string `json:"duration,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType AbsoluteDeleteOption_ObjectType_STATUS_ARM `json:"objectType,omitempty"`
}

// Adhoc backup tagging criteria
type AdhocBasedTaggingCriteria_STATUS_ARM struct {
	// TagInfo: Retention tag information
	TagInfo *RetentionTag_STATUS_ARM `json:"tagInfo,omitempty"`
}

type AdhocBasedTriggerContext_ObjectType_STATUS_ARM string

const AdhocBasedTriggerContext_ObjectType_STATUS_ARM_AdhocBasedTriggerContext = AdhocBasedTriggerContext_ObjectType_STATUS_ARM("AdhocBasedTriggerContext")

// Mapping from string to AdhocBasedTriggerContext_ObjectType_STATUS_ARM
var adhocBasedTriggerContext_ObjectType_STATUS_ARM_Values = map[string]AdhocBasedTriggerContext_ObjectType_STATUS_ARM{
	"adhocbasedtriggercontext": AdhocBasedTriggerContext_ObjectType_STATUS_ARM_AdhocBasedTriggerContext,
}

type AzureBackupParams_ObjectType_STATUS_ARM string

const AzureBackupParams_ObjectType_STATUS_ARM_AzureBackupParams = AzureBackupParams_ObjectType_STATUS_ARM("AzureBackupParams")

// Mapping from string to AzureBackupParams_ObjectType_STATUS_ARM
var azureBackupParams_ObjectType_STATUS_ARM_Values = map[string]AzureBackupParams_ObjectType_STATUS_ARM{
	"azurebackupparams": AzureBackupParams_ObjectType_STATUS_ARM_AzureBackupParams,
}

// Schedule for backup
type BackupSchedule_STATUS_ARM struct {
	// RepeatingTimeIntervals: ISO 8601 repeating time interval format
	RepeatingTimeIntervals []string `json:"repeatingTimeIntervals,omitempty"`

	// TimeZone: Time zone for a schedule. Example: Pacific Standard Time
	TimeZone *string `json:"timeZone,omitempty"`
}

type CopyOption_STATUS_ARM struct {
	// CopyOnExpiry: Mutually exclusive with all other properties
	CopyOnExpiry *CopyOnExpiryOption_STATUS_ARM `json:"copyOnExpiryOption,omitempty"`

	// CustomCopy: Mutually exclusive with all other properties
	CustomCopy *CustomCopyOption_STATUS_ARM `json:"customCopyOption,omitempty"`

	// ImmediateCopy: Mutually exclusive with all other properties
	ImmediateCopy *ImmediateCopyOption_STATUS_ARM `json:"immediateCopyOption,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because CopyOption_STATUS_ARM represents a discriminated union (JSON OneOf)
func (option CopyOption_STATUS_ARM) MarshalJSON() ([]byte, error) {
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

// UnmarshalJSON unmarshals the CopyOption_STATUS_ARM
func (option *CopyOption_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "CopyOnExpiryOption" {
		option.CopyOnExpiry = &CopyOnExpiryOption_STATUS_ARM{}
		return json.Unmarshal(data, option.CopyOnExpiry)
	}
	if discriminator == "CustomCopyOption" {
		option.CustomCopy = &CustomCopyOption_STATUS_ARM{}
		return json.Unmarshal(data, option.CustomCopy)
	}
	if discriminator == "ImmediateCopyOption" {
		option.ImmediateCopy = &ImmediateCopyOption_STATUS_ARM{}
		return json.Unmarshal(data, option.ImmediateCopy)
	}

	// No error
	return nil
}

type ScheduleBasedTriggerContext_ObjectType_STATUS_ARM string

const ScheduleBasedTriggerContext_ObjectType_STATUS_ARM_ScheduleBasedTriggerContext = ScheduleBasedTriggerContext_ObjectType_STATUS_ARM("ScheduleBasedTriggerContext")

// Mapping from string to ScheduleBasedTriggerContext_ObjectType_STATUS_ARM
var scheduleBasedTriggerContext_ObjectType_STATUS_ARM_Values = map[string]ScheduleBasedTriggerContext_ObjectType_STATUS_ARM{
	"schedulebasedtriggercontext": ScheduleBasedTriggerContext_ObjectType_STATUS_ARM_ScheduleBasedTriggerContext,
}

// Tagging criteria
type TaggingCriteria_STATUS_ARM struct {
	// Criteria: Criteria which decides whether the tag can be applied to a triggered backup.
	Criteria []BackupCriteria_STATUS_ARM `json:"criteria,omitempty"`

	// IsDefault: Specifies if tag is default.
	IsDefault *bool `json:"isDefault,omitempty"`

	// TagInfo: Retention tag information
	TagInfo *RetentionTag_STATUS_ARM `json:"tagInfo,omitempty"`

	// TaggingPriority: Retention Tag priority.
	TaggingPriority *int `json:"taggingPriority,omitempty"`
}

type AbsoluteDeleteOption_ObjectType_STATUS_ARM string

const AbsoluteDeleteOption_ObjectType_STATUS_ARM_AbsoluteDeleteOption = AbsoluteDeleteOption_ObjectType_STATUS_ARM("AbsoluteDeleteOption")

// Mapping from string to AbsoluteDeleteOption_ObjectType_STATUS_ARM
var absoluteDeleteOption_ObjectType_STATUS_ARM_Values = map[string]AbsoluteDeleteOption_ObjectType_STATUS_ARM{
	"absolutedeleteoption": AbsoluteDeleteOption_ObjectType_STATUS_ARM_AbsoluteDeleteOption,
}

type BackupCriteria_STATUS_ARM struct {
	// ScheduleBasedBackupCriteria: Mutually exclusive with all other properties
	ScheduleBasedBackupCriteria *ScheduleBasedBackupCriteria_STATUS_ARM `json:"scheduleBasedBackupCriteria,omitempty"`
}

// MarshalJSON defers JSON marshaling to the first non-nil property, because BackupCriteria_STATUS_ARM represents a discriminated union (JSON OneOf)
func (criteria BackupCriteria_STATUS_ARM) MarshalJSON() ([]byte, error) {
	if criteria.ScheduleBasedBackupCriteria != nil {
		return json.Marshal(criteria.ScheduleBasedBackupCriteria)
	}
	return nil, nil
}

// UnmarshalJSON unmarshals the BackupCriteria_STATUS_ARM
func (criteria *BackupCriteria_STATUS_ARM) UnmarshalJSON(data []byte) error {
	var rawJson map[string]interface{}
	err := json.Unmarshal(data, &rawJson)
	if err != nil {
		return err
	}
	discriminator := rawJson["objectType"]
	if discriminator == "ScheduleBasedBackupCriteria" {
		criteria.ScheduleBasedBackupCriteria = &ScheduleBasedBackupCriteria_STATUS_ARM{}
		return json.Unmarshal(data, criteria.ScheduleBasedBackupCriteria)
	}

	// No error
	return nil
}

type CopyOnExpiryOption_STATUS_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType CopyOnExpiryOption_ObjectType_STATUS_ARM `json:"objectType,omitempty"`
}

type CustomCopyOption_STATUS_ARM struct {
	// Duration: Data copied after given timespan
	Duration *string `json:"duration,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType CustomCopyOption_ObjectType_STATUS_ARM `json:"objectType,omitempty"`
}

type ImmediateCopyOption_STATUS_ARM struct {
	// ObjectType: Type of the specific object - used for deserializing
	ObjectType ImmediateCopyOption_ObjectType_STATUS_ARM `json:"objectType,omitempty"`
}

// Retention tag
type RetentionTag_STATUS_ARM struct {
	// ETag: Retention Tag version.
	ETag *string `json:"eTag,omitempty"`

	// Id: Retention Tag version.
	Id *string `json:"id,omitempty"`

	// TagName: Retention Tag Name to relate it to retention rule.
	TagName *string `json:"tagName,omitempty"`
}

type CopyOnExpiryOption_ObjectType_STATUS_ARM string

const CopyOnExpiryOption_ObjectType_STATUS_ARM_CopyOnExpiryOption = CopyOnExpiryOption_ObjectType_STATUS_ARM("CopyOnExpiryOption")

// Mapping from string to CopyOnExpiryOption_ObjectType_STATUS_ARM
var copyOnExpiryOption_ObjectType_STATUS_ARM_Values = map[string]CopyOnExpiryOption_ObjectType_STATUS_ARM{
	"copyonexpiryoption": CopyOnExpiryOption_ObjectType_STATUS_ARM_CopyOnExpiryOption,
}

type CustomCopyOption_ObjectType_STATUS_ARM string

const CustomCopyOption_ObjectType_STATUS_ARM_CustomCopyOption = CustomCopyOption_ObjectType_STATUS_ARM("CustomCopyOption")

// Mapping from string to CustomCopyOption_ObjectType_STATUS_ARM
var customCopyOption_ObjectType_STATUS_ARM_Values = map[string]CustomCopyOption_ObjectType_STATUS_ARM{
	"customcopyoption": CustomCopyOption_ObjectType_STATUS_ARM_CustomCopyOption,
}

type ImmediateCopyOption_ObjectType_STATUS_ARM string

const ImmediateCopyOption_ObjectType_STATUS_ARM_ImmediateCopyOption = ImmediateCopyOption_ObjectType_STATUS_ARM("ImmediateCopyOption")

// Mapping from string to ImmediateCopyOption_ObjectType_STATUS_ARM
var immediateCopyOption_ObjectType_STATUS_ARM_Values = map[string]ImmediateCopyOption_ObjectType_STATUS_ARM{
	"immediatecopyoption": ImmediateCopyOption_ObjectType_STATUS_ARM_ImmediateCopyOption,
}

type ScheduleBasedBackupCriteria_STATUS_ARM struct {
	// AbsoluteCriteria: it contains absolute values like "AllBackup" / "FirstOfDay" / "FirstOfWeek" / "FirstOfMonth"
	// and should be part of AbsoluteMarker enum
	AbsoluteCriteria []ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM `json:"absoluteCriteria,omitempty"`

	// DaysOfMonth: This is day of the month from 1 to 28 other wise last of month
	DaysOfMonth []Day_STATUS_ARM `json:"daysOfMonth,omitempty"`

	// DaysOfTheWeek: It should be Sunday/Monday/T..../Saturday
	DaysOfTheWeek []ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM `json:"daysOfTheWeek,omitempty"`

	// MonthsOfYear: It should be January/February/....../December
	MonthsOfYear []ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM `json:"monthsOfYear,omitempty"`

	// ObjectType: Type of the specific object - used for deserializing
	ObjectType ScheduleBasedBackupCriteria_ObjectType_STATUS_ARM `json:"objectType,omitempty"`

	// ScheduleTimes: List of schedule times for backup
	ScheduleTimes []string `json:"scheduleTimes,omitempty"`

	// WeeksOfTheMonth: It should be First/Second/Third/Fourth/Last
	WeeksOfTheMonth []ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM `json:"weeksOfTheMonth,omitempty"`
}

// Day of the week
type Day_STATUS_ARM struct {
	// Date: Date of the month
	Date *int `json:"date,omitempty"`

	// IsLast: Whether Date is last date of month
	IsLast *bool `json:"isLast,omitempty"`
}

type ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM string

const (
	ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_AllBackup    = ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM("AllBackup")
	ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_FirstOfDay   = ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM("FirstOfDay")
	ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_FirstOfMonth = ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM("FirstOfMonth")
	ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_FirstOfWeek  = ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM("FirstOfWeek")
	ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_FirstOfYear  = ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM("FirstOfYear")
)

// Mapping from string to ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM
var scheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_Values = map[string]ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM{
	"allbackup":    ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_AllBackup,
	"firstofday":   ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_FirstOfDay,
	"firstofmonth": ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_FirstOfMonth,
	"firstofweek":  ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_FirstOfWeek,
	"firstofyear":  ScheduleBasedBackupCriteria_AbsoluteCriteria_STATUS_ARM_FirstOfYear,
}

type ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM string

const (
	ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Friday    = ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM("Friday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Monday    = ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM("Monday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Saturday  = ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM("Saturday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Sunday    = ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM("Sunday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Thursday  = ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM("Thursday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Tuesday   = ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM("Tuesday")
	ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Wednesday = ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM("Wednesday")
)

// Mapping from string to ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM
var scheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Values = map[string]ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM{
	"friday":    ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Friday,
	"monday":    ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Monday,
	"saturday":  ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Saturday,
	"sunday":    ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Sunday,
	"thursday":  ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Thursday,
	"tuesday":   ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Tuesday,
	"wednesday": ScheduleBasedBackupCriteria_DaysOfTheWeek_STATUS_ARM_Wednesday,
}

type ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM string

const (
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_April     = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("April")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_August    = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("August")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_December  = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("December")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_February  = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("February")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_January   = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("January")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_July      = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("July")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_June      = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("June")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_March     = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("March")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_May       = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("May")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_November  = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("November")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_October   = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("October")
	ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_September = ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM("September")
)

// Mapping from string to ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM
var scheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_Values = map[string]ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM{
	"april":     ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_April,
	"august":    ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_August,
	"december":  ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_December,
	"february":  ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_February,
	"january":   ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_January,
	"july":      ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_July,
	"june":      ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_June,
	"march":     ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_March,
	"may":       ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_May,
	"november":  ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_November,
	"october":   ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_October,
	"september": ScheduleBasedBackupCriteria_MonthsOfYear_STATUS_ARM_September,
}

type ScheduleBasedBackupCriteria_ObjectType_STATUS_ARM string

const ScheduleBasedBackupCriteria_ObjectType_STATUS_ARM_ScheduleBasedBackupCriteria = ScheduleBasedBackupCriteria_ObjectType_STATUS_ARM("ScheduleBasedBackupCriteria")

// Mapping from string to ScheduleBasedBackupCriteria_ObjectType_STATUS_ARM
var scheduleBasedBackupCriteria_ObjectType_STATUS_ARM_Values = map[string]ScheduleBasedBackupCriteria_ObjectType_STATUS_ARM{
	"schedulebasedbackupcriteria": ScheduleBasedBackupCriteria_ObjectType_STATUS_ARM_ScheduleBasedBackupCriteria,
}

type ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM string

const (
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_First  = ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM("First")
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_Fourth = ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM("Fourth")
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_Last   = ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM("Last")
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_Second = ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM("Second")
	ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_Third  = ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM("Third")
)

// Mapping from string to ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM
var scheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_Values = map[string]ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM{
	"first":  ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_First,
	"fourth": ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_Fourth,
	"last":   ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_Last,
	"second": ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_Second,
	"third":  ScheduleBasedBackupCriteria_WeeksOfTheMonth_STATUS_ARM_Third,
}
