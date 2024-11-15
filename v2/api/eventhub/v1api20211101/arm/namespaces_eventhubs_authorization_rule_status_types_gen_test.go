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

func Test_NamespacesEventhubsAuthorizationRule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of NamespacesEventhubsAuthorizationRule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule_STATUS, NamespacesEventhubsAuthorizationRule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule_STATUS runs a test to see if a specific instance of NamespacesEventhubsAuthorizationRule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespacesEventhubsAuthorizationRule_STATUS(subject NamespacesEventhubsAuthorizationRule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual NamespacesEventhubsAuthorizationRule_STATUS
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

// Generator of NamespacesEventhubsAuthorizationRule_STATUS instances for property testing - lazily instantiated by
// NamespacesEventhubsAuthorizationRule_STATUSGenerator()
var namespacesEventhubsAuthorizationRule_STATUSGenerator gopter.Gen

// NamespacesEventhubsAuthorizationRule_STATUSGenerator returns a generator of NamespacesEventhubsAuthorizationRule_STATUS instances for property testing.
// We first initialize namespacesEventhubsAuthorizationRule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func NamespacesEventhubsAuthorizationRule_STATUSGenerator() gopter.Gen {
	if namespacesEventhubsAuthorizationRule_STATUSGenerator != nil {
		return namespacesEventhubsAuthorizationRule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_STATUS(generators)
	namespacesEventhubsAuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsAuthorizationRule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_STATUS(generators)
	AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_STATUS(generators)
	namespacesEventhubsAuthorizationRule_STATUSGenerator = gen.Struct(reflect.TypeOf(NamespacesEventhubsAuthorizationRule_STATUS{}), generators)

	return namespacesEventhubsAuthorizationRule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForNamespacesEventhubsAuthorizationRule_STATUS(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(Namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUSGenerator())
}

func Test_Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_Properties_STATUS, Namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_Properties_STATUS runs a test to see if a specific instance of Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForNamespaces_Eventhubs_AuthorizationRule_Properties_STATUS(subject Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS
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

// Generator of Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS instances for property testing - lazily
// instantiated by Namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator()
var namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator gopter.Gen

// Namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator returns a generator of Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS instances for property testing.
func Namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator() gopter.Gen {
	if namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator != nil {
		return namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_Properties_STATUS(generators)
	namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator = gen.Struct(reflect.TypeOf(Namespaces_Eventhubs_AuthorizationRule_Properties_STATUS{}), generators)

	return namespaces_Eventhubs_AuthorizationRule_Properties_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_Properties_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForNamespaces_Eventhubs_AuthorizationRule_Properties_STATUS(gens map[string]gopter.Gen) {
	gens["Rights"] = gen.SliceOf(gen.OneConstOf(Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Listen, Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Manage, Namespaces_Eventhubs_AuthorizationRule_Properties_Rights_STATUS_Send))
}
