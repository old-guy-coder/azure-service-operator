// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20211101

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

func Test_NamespacesEventhubsAuthorizationRule_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsAuthorizationRule_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule_Spec_ARM, NamespacesEventhubsAuthorizationRule_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule_Spec_ARM runs a test to see if a specific instance of NamespacesEventhubsAuthorizationRule_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule_Spec_ARM(subject NamespacesEventhubsAuthorizationRule_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsAuthorizationRule_Spec_ARM
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

// Generator of NamespacesEventhubsAuthorizationRule_Spec_ARM instances for property testing - lazily instantiated by
// NamespacesEventhubsAuthorizationRule_Spec_ARMGenerator()
var namespacesEventhubsAuthorizationRule_Spec_ARMGenerator gopter.Gen

// NamespacesEventhubsAuthorizationRule_Spec_ARMGenerator returns a generator of NamespacesEventhubsAuthorizationRule_Spec_ARM instances for property testing.
// We first initialize namespacesEventhubsAuthorizationRule_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsAuthorizationRule_Spec_ARMGenerator() gopter.Gen {
	if namespacesEventhubsAuthorizationRule_Spec_ARMGenerator != nil {
		return namespacesEventhubsAuthorizationRule_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_Spec_ARM(generators)
	namespacesEventhubsAuthorizationRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsAuthorizationRule_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_Spec_ARM(generators)
	namespacesEventhubsAuthorizationRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsAuthorizationRule_Spec_ARM{}), generators)

	return namespacesEventhubsAuthorizationRule_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator())
}

func Test_Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM, Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM runs a test to see if a specific instance of Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM(subject Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM
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

// Generator of Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM instances for property testing - lazily
// instantiated by Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator()
var namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator gopter.Gen

// Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator returns a generator of Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM instances for property testing.
func Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator() gopter.Gen {
	if namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator != nil {
		return namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM(generators)
	namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM{}), generators)

	return namespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_Properties_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec_ARM_Listen, Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec_ARM_Manage, Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_Spec_ARM_Send))
}
