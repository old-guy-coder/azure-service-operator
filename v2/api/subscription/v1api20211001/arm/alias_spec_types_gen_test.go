// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

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

func Test_Alias_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Alias_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAlias_Spec, Alias_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAlias_Spec runs a test to see if a specific instance of Alias_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForAlias_Spec(subject Alias_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Alias_Spec
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

// Generator of Alias_Spec instances for property testing - lazily instantiated by Alias_SpecGenerator()
var alias_SpecGenerator gopter.Gen

// Alias_SpecGenerator returns a generator of Alias_Spec instances for property testing.
// We first initialize alias_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Alias_SpecGenerator() gopter.Gen {
	if alias_SpecGenerator != nil {
		return alias_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_Spec(generators)
	alias_SpecGenerator = gen.Struct(reflect.TypeOf(Alias_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAlias_Spec(generators)
	AddRelatedPropertyGeneratorsForAlias_Spec(generators)
	alias_SpecGenerator = gen.Struct(reflect.TypeOf(Alias_Spec{}), generators)

	return alias_SpecGenerator
}

// AddIndependentPropertyGeneratorsForAlias_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAlias_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForAlias_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAlias_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(PutAliasRequestPropertiesGenerator())
}

func Test_PutAliasRequestAdditionalProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestAdditionalProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestAdditionalProperties, PutAliasRequestAdditionalPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestAdditionalProperties runs a test to see if a specific instance of PutAliasRequestAdditionalProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestAdditionalProperties(subject PutAliasRequestAdditionalProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestAdditionalProperties
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

// Generator of PutAliasRequestAdditionalProperties instances for property testing - lazily instantiated by
// PutAliasRequestAdditionalPropertiesGenerator()
var putAliasRequestAdditionalPropertiesGenerator gopter.Gen

// PutAliasRequestAdditionalPropertiesGenerator returns a generator of PutAliasRequestAdditionalProperties instances for property testing.
func PutAliasRequestAdditionalPropertiesGenerator() gopter.Gen {
	if putAliasRequestAdditionalPropertiesGenerator != nil {
		return putAliasRequestAdditionalPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties(generators)
	putAliasRequestAdditionalPropertiesGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestAdditionalProperties{}), generators)

	return putAliasRequestAdditionalPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestAdditionalProperties(gens map[string]gopter.Gen) {
	gens["ManagementGroupId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionOwnerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionTenantId"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

func Test_PutAliasRequestProperties_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of PutAliasRequestProperties via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPutAliasRequestProperties, PutAliasRequestPropertiesGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPutAliasRequestProperties runs a test to see if a specific instance of PutAliasRequestProperties round trips to JSON and back losslessly
func RunJSONSerializationTestForPutAliasRequestProperties(subject PutAliasRequestProperties) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual PutAliasRequestProperties
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

// Generator of PutAliasRequestProperties instances for property testing - lazily instantiated by
// PutAliasRequestPropertiesGenerator()
var putAliasRequestPropertiesGenerator gopter.Gen

// PutAliasRequestPropertiesGenerator returns a generator of PutAliasRequestProperties instances for property testing.
// We first initialize putAliasRequestPropertiesGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func PutAliasRequestPropertiesGenerator() gopter.Gen {
	if putAliasRequestPropertiesGenerator != nil {
		return putAliasRequestPropertiesGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestProperties(generators)
	putAliasRequestPropertiesGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestProperties{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPutAliasRequestProperties(generators)
	AddRelatedPropertyGeneratorsForPutAliasRequestProperties(generators)
	putAliasRequestPropertiesGenerator = gen.Struct(reflect.TypeOf(PutAliasRequestProperties{}), generators)

	return putAliasRequestPropertiesGenerator
}

// AddIndependentPropertyGeneratorsForPutAliasRequestProperties is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPutAliasRequestProperties(gens map[string]gopter.Gen) {
	gens["BillingScope"] = gen.PtrOf(gen.AlphaString())
	gens["DisplayName"] = gen.PtrOf(gen.AlphaString())
	gens["ResellerId"] = gen.PtrOf(gen.AlphaString())
	gens["SubscriptionId"] = gen.PtrOf(gen.AlphaString())
	gens["Workload"] = gen.PtrOf(gen.OneConstOf(Workload_DevTest, Workload_Production))
}

// AddRelatedPropertyGeneratorsForPutAliasRequestProperties is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForPutAliasRequestProperties(gens map[string]gopter.Gen) {
	gens["AdditionalProperties"] = gen.PtrOf(PutAliasRequestAdditionalPropertiesGenerator())
}