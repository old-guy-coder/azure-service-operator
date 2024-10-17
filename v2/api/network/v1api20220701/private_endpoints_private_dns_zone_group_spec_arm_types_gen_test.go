// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

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

func Test_PrivateDnsZoneConfig_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZoneConfig_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZoneConfig_ARM, PrivateDnsZoneConfig_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZoneConfig_ARM runs a test to see if a specific instance of PrivateDnsZoneConfig_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZoneConfig_ARM(subject PrivateDnsZoneConfig_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZoneConfig_ARM
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

// Generator of PrivateDnsZoneConfig_ARM instances for property testing - lazily instantiated by
// PrivateDnsZoneConfig_ARMGenerator()
var privateDnsZoneConfig_ARMGenerator gopter.Gen

// PrivateDnsZoneConfig_ARMGenerator returns a generator of PrivateDnsZoneConfig_ARM instances for property testing.
// We first initialize privateDnsZoneConfig_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateDnsZoneConfig_ARMGenerator() gopter.Gen {
	if privateDnsZoneConfig_ARMGenerator != nil {
		return privateDnsZoneConfig_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_ARM(generators)
	privateDnsZoneConfig_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneConfig_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig_ARM(generators)
	privateDnsZoneConfig_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneConfig_ARM{}), generators)

	return privateDnsZoneConfig_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZoneConfig_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZoneConfig_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateDnsZonePropertiesFormat_ARMGenerator())
}

func Test_PrivateDnsZoneGroupPropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZoneGroupPropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZoneGroupPropertiesFormat_ARM, PrivateDnsZoneGroupPropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZoneGroupPropertiesFormat_ARM runs a test to see if a specific instance of PrivateDnsZoneGroupPropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZoneGroupPropertiesFormat_ARM(subject PrivateDnsZoneGroupPropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZoneGroupPropertiesFormat_ARM
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

// Generator of PrivateDnsZoneGroupPropertiesFormat_ARM instances for property testing - lazily instantiated by
// PrivateDnsZoneGroupPropertiesFormat_ARMGenerator()
var privateDnsZoneGroupPropertiesFormat_ARMGenerator gopter.Gen

// PrivateDnsZoneGroupPropertiesFormat_ARMGenerator returns a generator of PrivateDnsZoneGroupPropertiesFormat_ARM instances for property testing.
func PrivateDnsZoneGroupPropertiesFormat_ARMGenerator() gopter.Gen {
	if privateDnsZoneGroupPropertiesFormat_ARMGenerator != nil {
		return privateDnsZoneGroupPropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_ARM(generators)
	privateDnsZoneGroupPropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZoneGroupPropertiesFormat_ARM{}), generators)

	return privateDnsZoneGroupPropertiesFormat_ARMGenerator
}

// AddRelatedPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateDnsZoneGroupPropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["PrivateDnsZoneConfigs"] = gen.SliceOf(PrivateDnsZoneConfig_ARMGenerator())
}

func Test_PrivateDnsZonePropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateDnsZonePropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateDnsZonePropertiesFormat_ARM, PrivateDnsZonePropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateDnsZonePropertiesFormat_ARM runs a test to see if a specific instance of PrivateDnsZonePropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateDnsZonePropertiesFormat_ARM(subject PrivateDnsZonePropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateDnsZonePropertiesFormat_ARM
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

// Generator of PrivateDnsZonePropertiesFormat_ARM instances for property testing - lazily instantiated by
// PrivateDnsZonePropertiesFormat_ARMGenerator()
var privateDnsZonePropertiesFormat_ARMGenerator gopter.Gen

// PrivateDnsZonePropertiesFormat_ARMGenerator returns a generator of PrivateDnsZonePropertiesFormat_ARM instances for property testing.
func PrivateDnsZonePropertiesFormat_ARMGenerator() gopter.Gen {
	if privateDnsZonePropertiesFormat_ARMGenerator != nil {
		return privateDnsZonePropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat_ARM(generators)
	privateDnsZonePropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateDnsZonePropertiesFormat_ARM{}), generators)

	return privateDnsZonePropertiesFormat_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateDnsZonePropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["PrivateDnsZoneId"] = gen.PtrOf(gen.AlphaString())
}

func Test_PrivateEndpointsPrivateDnsZoneGroup_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PrivateEndpointsPrivateDnsZoneGroup_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM, PrivateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM runs a test to see if a specific instance of PrivateEndpointsPrivateDnsZoneGroup_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM(subject PrivateEndpointsPrivateDnsZoneGroup_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PrivateEndpointsPrivateDnsZoneGroup_Spec_ARM
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

// Generator of PrivateEndpointsPrivateDnsZoneGroup_Spec_ARM instances for property testing - lazily instantiated by
// PrivateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator()
var privateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator gopter.Gen

// PrivateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator returns a generator of PrivateEndpointsPrivateDnsZoneGroup_Spec_ARM instances for property testing.
// We first initialize privateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PrivateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator() gopter.Gen {
	if privateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator != nil {
		return privateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM(generators)
	privateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointsPrivateDnsZoneGroup_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM(generators)
	privateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(PrivateEndpointsPrivateDnsZoneGroup_Spec_ARM{}), generators)

	return privateEndpointsPrivateDnsZoneGroup_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPrivateEndpointsPrivateDnsZoneGroup_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PrivateDnsZoneGroupPropertiesFormat_ARMGenerator())
}
