// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
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

func Test_FlexibleServersDatabase_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersDatabase via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersDatabase, FlexibleServersDatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersDatabase runs a test to see if a specific instance of FlexibleServersDatabase round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersDatabase(subject FlexibleServersDatabase) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersDatabase
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

// Generator of FlexibleServersDatabase instances for property testing - lazily instantiated by
// FlexibleServersDatabaseGenerator()
var flexibleServersDatabaseGenerator gopter.Gen

// FlexibleServersDatabaseGenerator returns a generator of FlexibleServersDatabase instances for property testing.
func FlexibleServersDatabaseGenerator() gopter.Gen {
	if flexibleServersDatabaseGenerator != nil {
		return flexibleServersDatabaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForFlexibleServersDatabase(generators)
	flexibleServersDatabaseGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase{}), generators)

	return flexibleServersDatabaseGenerator
}

// AddRelatedPropertyGeneratorsForFlexibleServersDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = FlexibleServersDatabase_SpecGenerator()
	gens["Status"] = FlexibleServersDatabase_STATUSGenerator()
}

func Test_FlexibleServersDatabase_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersDatabase_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersDatabase_STATUS, FlexibleServersDatabase_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersDatabase_STATUS runs a test to see if a specific instance of FlexibleServersDatabase_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersDatabase_STATUS(subject FlexibleServersDatabase_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersDatabase_STATUS
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

// Generator of FlexibleServersDatabase_STATUS instances for property testing - lazily instantiated by
// FlexibleServersDatabase_STATUSGenerator()
var flexibleServersDatabase_STATUSGenerator gopter.Gen

// FlexibleServersDatabase_STATUSGenerator returns a generator of FlexibleServersDatabase_STATUS instances for property testing.
// We first initialize flexibleServersDatabase_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func FlexibleServersDatabase_STATUSGenerator() gopter.Gen {
	if flexibleServersDatabase_STATUSGenerator != nil {
		return flexibleServersDatabase_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabase_STATUS(generators)
	flexibleServersDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabase_STATUS(generators)
	AddRelatedPropertyGeneratorsForFlexibleServersDatabase_STATUS(generators)
	flexibleServersDatabase_STATUSGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase_STATUS{}), generators)

	return flexibleServersDatabase_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForFlexibleServersDatabase_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForFlexibleServersDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_FlexibleServersDatabase_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of FlexibleServersDatabase_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForFlexibleServersDatabase_Spec, FlexibleServersDatabase_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForFlexibleServersDatabase_Spec runs a test to see if a specific instance of FlexibleServersDatabase_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForFlexibleServersDatabase_Spec(subject FlexibleServersDatabase_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual FlexibleServersDatabase_Spec
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

// Generator of FlexibleServersDatabase_Spec instances for property testing - lazily instantiated by
// FlexibleServersDatabase_SpecGenerator()
var flexibleServersDatabase_SpecGenerator gopter.Gen

// FlexibleServersDatabase_SpecGenerator returns a generator of FlexibleServersDatabase_Spec instances for property testing.
func FlexibleServersDatabase_SpecGenerator() gopter.Gen {
	if flexibleServersDatabase_SpecGenerator != nil {
		return flexibleServersDatabase_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForFlexibleServersDatabase_Spec(generators)
	flexibleServersDatabase_SpecGenerator = gen.Struct(reflect.TypeOf(FlexibleServersDatabase_Spec{}), generators)

	return flexibleServersDatabase_SpecGenerator
}

// AddIndependentPropertyGeneratorsForFlexibleServersDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForFlexibleServersDatabase_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}
