// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

import (
	"encoding/json"
	v20220901s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20220901/storage"
	v20230101s "github.com/Azure/azure-service-operator/v2/api/storage/v1api20230101/storage"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/kr/pretty"
	"github.com/kylelemons/godebug/diff"
	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"os"
	"reflect"
	"testing"
)

func Test_Multichannel_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Multichannel to Multichannel via AssignProperties_To_Multichannel & AssignProperties_From_Multichannel returns original",
		prop.ForAll(RunPropertyAssignmentTestForMultichannel, MultichannelGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForMultichannel tests if a specific instance of Multichannel can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForMultichannel(subject Multichannel) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.Multichannel
	err := copied.AssignProperties_To_Multichannel(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Multichannel
	err = actual.AssignProperties_From_Multichannel(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Multichannel_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Multichannel via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMultichannel, MultichannelGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMultichannel runs a test to see if a specific instance of Multichannel round trips to JSON and back losslessly
func RunJSONSerializationTestForMultichannel(subject Multichannel) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Multichannel
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Multichannel instances for property testing - lazily instantiated by MultichannelGenerator()
var multichannelGenerator gopter.Gen

// MultichannelGenerator returns a generator of Multichannel instances for property testing.
func MultichannelGenerator() gopter.Gen {
	if multichannelGenerator != nil {
		return multichannelGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMultichannel(generators)
	multichannelGenerator = gen.Struct(reflect.TypeOf(Multichannel{}), generators)

	return multichannelGenerator
}

// AddIndependentPropertyGeneratorsForMultichannel is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMultichannel(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_Multichannel_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from Multichannel_STATUS to Multichannel_STATUS via AssignProperties_To_Multichannel_STATUS & AssignProperties_From_Multichannel_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForMultichannel_STATUS, Multichannel_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForMultichannel_STATUS tests if a specific instance of Multichannel_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForMultichannel_STATUS(subject Multichannel_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.Multichannel_STATUS
	err := copied.AssignProperties_To_Multichannel_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual Multichannel_STATUS
	err = actual.AssignProperties_From_Multichannel_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_Multichannel_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Multichannel_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMultichannel_STATUS, Multichannel_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMultichannel_STATUS runs a test to see if a specific instance of Multichannel_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForMultichannel_STATUS(subject Multichannel_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Multichannel_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of Multichannel_STATUS instances for property testing - lazily instantiated by
// Multichannel_STATUSGenerator()
var multichannel_STATUSGenerator gopter.Gen

// Multichannel_STATUSGenerator returns a generator of Multichannel_STATUS instances for property testing.
func Multichannel_STATUSGenerator() gopter.Gen {
	if multichannel_STATUSGenerator != nil {
		return multichannel_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMultichannel_STATUS(generators)
	multichannel_STATUSGenerator = gen.Struct(reflect.TypeOf(Multichannel_STATUS{}), generators)

	return multichannel_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForMultichannel_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMultichannel_STATUS(gens map[string]gopter.Gen) {
	gens["Enabled"] = gen.PtrOf(gen.Bool())
}

func Test_ProtocolSettings_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ProtocolSettings to ProtocolSettings via AssignProperties_To_ProtocolSettings & AssignProperties_From_ProtocolSettings returns original",
		prop.ForAll(RunPropertyAssignmentTestForProtocolSettings, ProtocolSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProtocolSettings tests if a specific instance of ProtocolSettings can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProtocolSettings(subject ProtocolSettings) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.ProtocolSettings
	err := copied.AssignProperties_To_ProtocolSettings(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ProtocolSettings
	err = actual.AssignProperties_From_ProtocolSettings(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ProtocolSettings_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProtocolSettings via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProtocolSettings, ProtocolSettingsGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProtocolSettings runs a test to see if a specific instance of ProtocolSettings round trips to JSON and back losslessly
func RunJSONSerializationTestForProtocolSettings(subject ProtocolSettings) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProtocolSettings
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ProtocolSettings instances for property testing - lazily instantiated by ProtocolSettingsGenerator()
var protocolSettingsGenerator gopter.Gen

// ProtocolSettingsGenerator returns a generator of ProtocolSettings instances for property testing.
func ProtocolSettingsGenerator() gopter.Gen {
	if protocolSettingsGenerator != nil {
		return protocolSettingsGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForProtocolSettings(generators)
	protocolSettingsGenerator = gen.Struct(reflect.TypeOf(ProtocolSettings{}), generators)

	return protocolSettingsGenerator
}

// AddRelatedPropertyGeneratorsForProtocolSettings is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProtocolSettings(gens map[string]gopter.Gen) {
	gens["Smb"] = gen.PtrOf(SmbSettingGenerator())
}

func Test_ProtocolSettings_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ProtocolSettings_STATUS to ProtocolSettings_STATUS via AssignProperties_To_ProtocolSettings_STATUS & AssignProperties_From_ProtocolSettings_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForProtocolSettings_STATUS, ProtocolSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForProtocolSettings_STATUS tests if a specific instance of ProtocolSettings_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForProtocolSettings_STATUS(subject ProtocolSettings_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.ProtocolSettings_STATUS
	err := copied.AssignProperties_To_ProtocolSettings_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ProtocolSettings_STATUS
	err = actual.AssignProperties_From_ProtocolSettings_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_ProtocolSettings_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ProtocolSettings_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForProtocolSettings_STATUS, ProtocolSettings_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForProtocolSettings_STATUS runs a test to see if a specific instance of ProtocolSettings_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForProtocolSettings_STATUS(subject ProtocolSettings_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ProtocolSettings_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of ProtocolSettings_STATUS instances for property testing - lazily instantiated by
// ProtocolSettings_STATUSGenerator()
var protocolSettings_STATUSGenerator gopter.Gen

// ProtocolSettings_STATUSGenerator returns a generator of ProtocolSettings_STATUS instances for property testing.
func ProtocolSettings_STATUSGenerator() gopter.Gen {
	if protocolSettings_STATUSGenerator != nil {
		return protocolSettings_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForProtocolSettings_STATUS(generators)
	protocolSettings_STATUSGenerator = gen.Struct(reflect.TypeOf(ProtocolSettings_STATUS{}), generators)

	return protocolSettings_STATUSGenerator
}

// AddRelatedPropertyGeneratorsForProtocolSettings_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForProtocolSettings_STATUS(gens map[string]gopter.Gen) {
	gens["Smb"] = gen.PtrOf(SmbSetting_STATUSGenerator())
}

func Test_SmbSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SmbSetting to SmbSetting via AssignProperties_To_SmbSetting & AssignProperties_From_SmbSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForSmbSetting, SmbSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSmbSetting tests if a specific instance of SmbSetting can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSmbSetting(subject SmbSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.SmbSetting
	err := copied.AssignProperties_To_SmbSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SmbSetting
	err = actual.AssignProperties_From_SmbSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SmbSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SmbSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSmbSetting, SmbSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSmbSetting runs a test to see if a specific instance of SmbSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSmbSetting(subject SmbSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SmbSetting
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SmbSetting instances for property testing - lazily instantiated by SmbSettingGenerator()
var smbSettingGenerator gopter.Gen

// SmbSettingGenerator returns a generator of SmbSetting instances for property testing.
// We first initialize smbSettingGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SmbSettingGenerator() gopter.Gen {
	if smbSettingGenerator != nil {
		return smbSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting(generators)
	smbSettingGenerator = gen.Struct(reflect.TypeOf(SmbSetting{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting(generators)
	AddRelatedPropertyGeneratorsForSmbSetting(generators)
	smbSettingGenerator = gen.Struct(reflect.TypeOf(SmbSetting{}), generators)

	return smbSettingGenerator
}

// AddIndependentPropertyGeneratorsForSmbSetting is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSmbSetting(gens map[string]gopter.Gen) {
	gens["AuthenticationMethods"] = gen.PtrOf(gen.AlphaString())
	gens["ChannelEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["KerberosTicketEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["Versions"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSmbSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSmbSetting(gens map[string]gopter.Gen) {
	gens["Multichannel"] = gen.PtrOf(MultichannelGenerator())
}

func Test_SmbSetting_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SmbSetting_STATUS to SmbSetting_STATUS via AssignProperties_To_SmbSetting_STATUS & AssignProperties_From_SmbSetting_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSmbSetting_STATUS, SmbSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSmbSetting_STATUS tests if a specific instance of SmbSetting_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSmbSetting_STATUS(subject SmbSetting_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.SmbSetting_STATUS
	err := copied.AssignProperties_To_SmbSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SmbSetting_STATUS
	err = actual.AssignProperties_From_SmbSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_SmbSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SmbSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSmbSetting_STATUS, SmbSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSmbSetting_STATUS runs a test to see if a specific instance of SmbSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSmbSetting_STATUS(subject SmbSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SmbSetting_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of SmbSetting_STATUS instances for property testing - lazily instantiated by SmbSetting_STATUSGenerator()
var smbSetting_STATUSGenerator gopter.Gen

// SmbSetting_STATUSGenerator returns a generator of SmbSetting_STATUS instances for property testing.
// We first initialize smbSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SmbSetting_STATUSGenerator() gopter.Gen {
	if smbSetting_STATUSGenerator != nil {
		return smbSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting_STATUS(generators)
	smbSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(SmbSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSmbSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForSmbSetting_STATUS(generators)
	smbSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(SmbSetting_STATUS{}), generators)

	return smbSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSmbSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSmbSetting_STATUS(gens map[string]gopter.Gen) {
	gens["AuthenticationMethods"] = gen.PtrOf(gen.AlphaString())
	gens["ChannelEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["KerberosTicketEncryption"] = gen.PtrOf(gen.AlphaString())
	gens["Versions"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSmbSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSmbSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Multichannel"] = gen.PtrOf(Multichannel_STATUSGenerator())
}

func Test_StorageAccountsFileService_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsFileService to hub returns original",
		prop.ForAll(RunResourceConversionTestForStorageAccountsFileService, StorageAccountsFileServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForStorageAccountsFileService tests if a specific instance of StorageAccountsFileService round trips to the hub storage version and back losslessly
func RunResourceConversionTestForStorageAccountsFileService(subject StorageAccountsFileService) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20230101s.StorageAccountsFileService
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual StorageAccountsFileService
	err = actual.ConvertFrom(&hub)
	if err != nil {
		return err.Error()
	}

	// Compare actual with what we started with
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsFileService_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsFileService to StorageAccountsFileService via AssignProperties_To_StorageAccountsFileService & AssignProperties_From_StorageAccountsFileService returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsFileService, StorageAccountsFileServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsFileService tests if a specific instance of StorageAccountsFileService can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsFileService(subject StorageAccountsFileService) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccountsFileService
	err := copied.AssignProperties_To_StorageAccountsFileService(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsFileService
	err = actual.AssignProperties_From_StorageAccountsFileService(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsFileService_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsFileService via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsFileService, StorageAccountsFileServiceGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsFileService runs a test to see if a specific instance of StorageAccountsFileService round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsFileService(subject StorageAccountsFileService) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsFileService
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of StorageAccountsFileService instances for property testing - lazily instantiated by
// StorageAccountsFileServiceGenerator()
var storageAccountsFileServiceGenerator gopter.Gen

// StorageAccountsFileServiceGenerator returns a generator of StorageAccountsFileService instances for property testing.
func StorageAccountsFileServiceGenerator() gopter.Gen {
	if storageAccountsFileServiceGenerator != nil {
		return storageAccountsFileServiceGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsFileService(generators)
	storageAccountsFileServiceGenerator = gen.Struct(reflect.TypeOf(StorageAccountsFileService{}), generators)

	return storageAccountsFileServiceGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsFileService is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsFileService(gens map[string]gopter.Gen) {
	gens["Spec"] = StorageAccountsFileService_SpecGenerator()
	gens["Status"] = StorageAccountsFileService_STATUSGenerator()
}

func Test_StorageAccountsFileService_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsFileService_STATUS to StorageAccountsFileService_STATUS via AssignProperties_To_StorageAccountsFileService_STATUS & AssignProperties_From_StorageAccountsFileService_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsFileService_STATUS, StorageAccountsFileService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsFileService_STATUS tests if a specific instance of StorageAccountsFileService_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsFileService_STATUS(subject StorageAccountsFileService_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccountsFileService_STATUS
	err := copied.AssignProperties_To_StorageAccountsFileService_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsFileService_STATUS
	err = actual.AssignProperties_From_StorageAccountsFileService_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsFileService_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsFileService_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsFileService_STATUS, StorageAccountsFileService_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsFileService_STATUS runs a test to see if a specific instance of StorageAccountsFileService_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsFileService_STATUS(subject StorageAccountsFileService_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsFileService_STATUS
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of StorageAccountsFileService_STATUS instances for property testing - lazily instantiated by
// StorageAccountsFileService_STATUSGenerator()
var storageAccountsFileService_STATUSGenerator gopter.Gen

// StorageAccountsFileService_STATUSGenerator returns a generator of StorageAccountsFileService_STATUS instances for property testing.
// We first initialize storageAccountsFileService_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsFileService_STATUSGenerator() gopter.Gen {
	if storageAccountsFileService_STATUSGenerator != nil {
		return storageAccountsFileService_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsFileService_STATUS(generators)
	storageAccountsFileService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsFileService_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsFileService_STATUS(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsFileService_STATUS(generators)
	storageAccountsFileService_STATUSGenerator = gen.Struct(reflect.TypeOf(StorageAccountsFileService_STATUS{}), generators)

	return storageAccountsFileService_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsFileService_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsFileService_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsFileService_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsFileService_STATUS(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRules_STATUSGenerator())
	gens["ProtocolSettings"] = gen.PtrOf(ProtocolSettings_STATUSGenerator())
	gens["ShareDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicy_STATUSGenerator())
	gens["Sku"] = gen.PtrOf(Sku_STATUSGenerator())
}

func Test_StorageAccountsFileService_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from StorageAccountsFileService_Spec to StorageAccountsFileService_Spec via AssignProperties_To_StorageAccountsFileService_Spec & AssignProperties_From_StorageAccountsFileService_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForStorageAccountsFileService_Spec, StorageAccountsFileService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForStorageAccountsFileService_Spec tests if a specific instance of StorageAccountsFileService_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForStorageAccountsFileService_Spec(subject StorageAccountsFileService_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20220901s.StorageAccountsFileService_Spec
	err := copied.AssignProperties_To_StorageAccountsFileService_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual StorageAccountsFileService_Spec
	err = actual.AssignProperties_From_StorageAccountsFileService_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Check for a match
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

func Test_StorageAccountsFileService_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsFileService_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsFileService_Spec, StorageAccountsFileService_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsFileService_Spec runs a test to see if a specific instance of StorageAccountsFileService_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsFileService_Spec(subject StorageAccountsFileService_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsFileService_Spec
	err = json.Unmarshal(bin, &actual)
	if err != nil {
		return err.Error()
	}

	// Check for outcome
	match := cmp.Equal(subject, actual, cmpopts.EquateEmpty())
	if !match {
		actualFmt := pretty.Sprint(actual)
		subjectFmt := pretty.Sprint(subject)
		result := diff.Diff(subjectFmt, actualFmt)
		return result
	}

	return ""
}

// Generator of StorageAccountsFileService_Spec instances for property testing - lazily instantiated by
// StorageAccountsFileService_SpecGenerator()
var storageAccountsFileService_SpecGenerator gopter.Gen

// StorageAccountsFileService_SpecGenerator returns a generator of StorageAccountsFileService_Spec instances for property testing.
func StorageAccountsFileService_SpecGenerator() gopter.Gen {
	if storageAccountsFileService_SpecGenerator != nil {
		return storageAccountsFileService_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccountsFileService_Spec(generators)
	storageAccountsFileService_SpecGenerator = gen.Struct(reflect.TypeOf(StorageAccountsFileService_Spec{}), generators)

	return storageAccountsFileService_SpecGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccountsFileService_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsFileService_Spec(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRulesGenerator())
	gens["ProtocolSettings"] = gen.PtrOf(ProtocolSettingsGenerator())
	gens["ShareDeleteRetentionPolicy"] = gen.PtrOf(DeleteRetentionPolicyGenerator())
}
