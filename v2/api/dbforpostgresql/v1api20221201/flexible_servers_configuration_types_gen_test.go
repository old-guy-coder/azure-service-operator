// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20221201

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/dbforpostgresql/v1api20221201/storage"
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

func Test_FlexibleServersConfiguration_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersConfiguration to hub returns original",
		prop.ForAll(RunResourceConversionTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForFlexibleServersConfiguration tests if a specific instance of FlexibleServersConfiguration round trips to the hub storage version and back losslessly
func RunResourceConversionTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.FlexibleServersConfiguration
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual FlexibleServersConfiguration
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

func Test_FlexibleServersConfiguration_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersConfiguration to FlexibleServersConfiguration via AssignProperties_To_FlexibleServersConfiguration & AssignProperties_From_FlexibleServersConfiguration returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersConfiguration tests if a specific instance of FlexibleServersConfiguration can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServersConfiguration
	err := copied.AssignProperties_To_FlexibleServersConfiguration(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersConfiguration
	err = actual.AssignProperties_From_FlexibleServersConfiguration(&other)
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

func Test_FlexibleServersConfiguration_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfiguration via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfiguration, FlexibleServersConfigurationGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfiguration runs a test to see if a specific instance of FlexibleServersConfiguration round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfiguration(subject FlexibleServersConfiguration) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfiguration
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

// Generator of FlexibleServersConfiguration instances for property testing - lazily instantiated by
// FlexibleServersConfigurationGenerator()
var flexibleServersConfigurationGenerator gopter.Gen

// FlexibleServersConfigurationGenerator returns a generator of FlexibleServersConfiguration instances for property testing.
func FlexibleServersConfigurationGenerator() gopter.Gen {
	if flexibleServersConfigurationGenerator != nil {
		return flexibleServersConfigurationGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(generators)
	flexibleServersConfigurationGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration{}), generators)

	return flexibleServersConfigurationGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersConfiguration is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersConfiguration(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersConfiguration_SpecGenerator()
	gens["Status"] = FlexibleServersConfiguration_STATUSGenerator()
}

func Test_FlexibleServersConfiguration_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersConfiguration_STATUS to FlexibleServersConfiguration_STATUS via AssignProperties_To_FlexibleServersConfiguration_STATUS & AssignProperties_From_FlexibleServersConfiguration_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersConfiguration_STATUS, FlexibleServersConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersConfiguration_STATUS tests if a specific instance of FlexibleServersConfiguration_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersConfiguration_STATUS(subject FlexibleServersConfiguration_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServersConfiguration_STATUS
	err := copied.AssignProperties_To_FlexibleServersConfiguration_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersConfiguration_STATUS
	err = actual.AssignProperties_From_FlexibleServersConfiguration_STATUS(&other)
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

func Test_FlexibleServersConfiguration_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfiguration_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfiguration_STATUS, FlexibleServersConfiguration_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfiguration_STATUS runs a test to see if a specific instance of FlexibleServersConfiguration_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfiguration_STATUS(subject FlexibleServersConfiguration_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfiguration_STATUS
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

// Generator of FlexibleServersConfiguration_STATUS instances for property testing - lazily instantiated by
// FlexibleServersConfiguration_STATUSGenerator()
var flexibleServersConfiguration_STATUSGenerator gopter.Gen

// FlexibleServersConfiguration_STATUSGenerator returns a generator of FlexibleServersConfiguration_STATUS instances for property testing.
// We first initialize flexibleServersConfiguration_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersConfiguration_STATUSGenerator() gopter.Gen {
	if flexibleServersConfiguration_STATUSGenerator != nil {
		return flexibleServersConfiguration_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS(generators)
	flexibleServersConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersConfiguration_STATUS(generators)
	flexibleServersConfiguration_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration_STATUS{}), generators)

	return flexibleServersConfiguration_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["AllowedValues"] = gen.PtrOf(gen.AlphaString())
	gens["DataType"] = gen.PtrOf(gen.OneConstOf(
		ConfigurationProperties_DataType_STATUS_Boolean,
		ConfigurationProperties_DataType_STATUS_Enumeration,
		ConfigurationProperties_DataType_STATUS_Integer,
		ConfigurationProperties_DataType_STATUS_Numeric))
	gens["DefaultValue"] = gen.PtrOf(gen.AlphaString())
	gens["Description"] = gen.PtrOf(gen.AlphaString())
	gens["DocumentationLink"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["IsConfigPendingRestart"] = gen.PtrOf(gen.Bool())
	gens["IsDynamicConfig"] = gen.PtrOf(gen.Bool())
	gens["IsReadOnly"] = gen.PtrOf(gen.Bool())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
	gens["Unit"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersConfiguration_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersConfiguration_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_FlexibleServersConfiguration_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from FlexibleServersConfiguration_Spec to FlexibleServersConfiguration_Spec via AssignProperties_To_FlexibleServersConfiguration_Spec & AssignProperties_From_FlexibleServersConfiguration_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForFlexibleServersConfiguration_Spec, FlexibleServersConfiguration_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForFlexibleServersConfiguration_Spec tests if a specific instance of FlexibleServersConfiguration_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForFlexibleServersConfiguration_Spec(subject FlexibleServersConfiguration_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.FlexibleServersConfiguration_Spec
	err := copied.AssignProperties_To_FlexibleServersConfiguration_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual FlexibleServersConfiguration_Spec
	err = actual.AssignProperties_From_FlexibleServersConfiguration_Spec(&other)
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

func Test_FlexibleServersConfiguration_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersConfiguration_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersConfiguration_Spec, FlexibleServersConfiguration_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersConfiguration_Spec runs a test to see if a specific instance of FlexibleServersConfiguration_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersConfiguration_Spec(subject FlexibleServersConfiguration_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersConfiguration_Spec
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

// Generator of FlexibleServersConfiguration_Spec instances for property testing - lazily instantiated by
// FlexibleServersConfiguration_SpecGenerator()
var flexibleServersConfiguration_SpecGenerator gopter.Gen

// FlexibleServersConfiguration_SpecGenerator returns a generator of FlexibleServersConfiguration_Spec instances for property testing.
func FlexibleServersConfiguration_SpecGenerator() gopter.Gen {
	if flexibleServersConfiguration_SpecGenerator != nil {
		return flexibleServersConfiguration_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_Spec(generators)
	flexibleServersConfiguration_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersConfiguration_Spec{}), generators)

	return flexibleServersConfiguration_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersConfiguration_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Source"] = gen.PtrOf(gen.AlphaString())
	gens["Value"] = gen.PtrOf(gen.AlphaString())
}
