// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220901

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

func Test_StorageAccountsTableService_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccountsTableService_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccountsTableService_STATUS_ARM, StorageAccountsTableService_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccountsTableService_STATUS_ARM runs a test to see if a specific instance of StorageAccountsTableService_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccountsTableService_STATUS_ARM(subject StorageAccountsTableService_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccountsTableService_STATUS_ARM
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

// Generator of StorageAccountsTableService_STATUS_ARM instances for property testing - lazily instantiated by
// StorageAccountsTableService_STATUS_ARMGenerator()
var storageAccountsTableService_STATUS_ARMGenerator gopter.Gen

// StorageAccountsTableService_STATUS_ARMGenerator returns a generator of StorageAccountsTableService_STATUS_ARM instances for property testing.
// We first initialize storageAccountsTableService_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func StorageAccountsTableService_STATUS_ARMGenerator() gopter.Gen {
	if storageAccountsTableService_STATUS_ARMGenerator != nil {
		return storageAccountsTableService_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableService_STATUS_ARM(generators)
	storageAccountsTableService_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableService_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForStorageAccountsTableService_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForStorageAccountsTableService_STATUS_ARM(generators)
	storageAccountsTableService_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccountsTableService_STATUS_ARM{}), generators)

	return storageAccountsTableService_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForStorageAccountsTableService_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForStorageAccountsTableService_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForStorageAccountsTableService_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccountsTableService_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(StorageAccounts_TableService_Properties_STATUS_ARMGenerator())
}

func Test_StorageAccounts_TableService_Properties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of StorageAccounts_TableService_Properties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForStorageAccounts_TableService_Properties_STATUS_ARM, StorageAccounts_TableService_Properties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForStorageAccounts_TableService_Properties_STATUS_ARM runs a test to see if a specific instance of StorageAccounts_TableService_Properties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForStorageAccounts_TableService_Properties_STATUS_ARM(subject StorageAccounts_TableService_Properties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual StorageAccounts_TableService_Properties_STATUS_ARM
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

// Generator of StorageAccounts_TableService_Properties_STATUS_ARM instances for property testing - lazily instantiated
// by StorageAccounts_TableService_Properties_STATUS_ARMGenerator()
var storageAccounts_TableService_Properties_STATUS_ARMGenerator gopter.Gen

// StorageAccounts_TableService_Properties_STATUS_ARMGenerator returns a generator of StorageAccounts_TableService_Properties_STATUS_ARM instances for property testing.
func StorageAccounts_TableService_Properties_STATUS_ARMGenerator() gopter.Gen {
	if storageAccounts_TableService_Properties_STATUS_ARMGenerator != nil {
		return storageAccounts_TableService_Properties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForStorageAccounts_TableService_Properties_STATUS_ARM(generators)
	storageAccounts_TableService_Properties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(StorageAccounts_TableService_Properties_STATUS_ARM{}), generators)

	return storageAccounts_TableService_Properties_STATUS_ARMGenerator
}

// AddRelatedPropertyGeneratorsForStorageAccounts_TableService_Properties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForStorageAccounts_TableService_Properties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Cors"] = gen.PtrOf(CorsRules_STATUS_ARMGenerator())
}
