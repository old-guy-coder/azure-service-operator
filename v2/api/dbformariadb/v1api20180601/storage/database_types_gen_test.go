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

func Test_Database_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase, DatabaseGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase runs a test to see if a specific instance of Database round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase(subject Database) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database
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

// Generator of Database instances for property testing - lazily instantiated by DatabaseGenerator()
var databaseGenerator gopter.Gen

// DatabaseGenerator returns a generator of Database instances for property testing.
func DatabaseGenerator() gopter.Gen {
	if databaseGenerator != nil {
		return databaseGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDatabase(generators)
	databaseGenerator = gen.Struct(reflect.TypeOf(Database{}), generators)

	return databaseGenerator
}

// AddRelatedPropertyGeneratorsForDatabase is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase(gens map[string]gopter.Gen) {
	gens["Spec"] = Database_SpecGenerator()
	gens["Status"] = Database_STATUSGenerator()
}

func Test_DatabaseOperatorSpec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseOperatorSpec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseOperatorSpec, DatabaseOperatorSpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseOperatorSpec runs a test to see if a specific instance of DatabaseOperatorSpec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseOperatorSpec(subject DatabaseOperatorSpec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseOperatorSpec
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

// Generator of DatabaseOperatorSpec instances for property testing - lazily instantiated by
// DatabaseOperatorSpecGenerator()
var databaseOperatorSpecGenerator gopter.Gen

// DatabaseOperatorSpecGenerator returns a generator of DatabaseOperatorSpec instances for property testing.
func DatabaseOperatorSpecGenerator() gopter.Gen {
	if databaseOperatorSpecGenerator != nil {
		return databaseOperatorSpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	databaseOperatorSpecGenerator = gen.Struct(reflect.TypeOf(DatabaseOperatorSpec{}), generators)

	return databaseOperatorSpecGenerator
}

func Test_Database_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase_STATUS, Database_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase_STATUS runs a test to see if a specific instance of Database_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase_STATUS(subject Database_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_STATUS
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

// Generator of Database_STATUS instances for property testing - lazily instantiated by Database_STATUSGenerator()
var database_STATUSGenerator gopter.Gen

// Database_STATUSGenerator returns a generator of Database_STATUS instances for property testing.
func Database_STATUSGenerator() gopter.Gen {
	if database_STATUSGenerator != nil {
		return database_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabase_STATUS(generators)
	database_STATUSGenerator = gen.Struct(reflect.TypeOf(Database_STATUS{}), generators)

	return database_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDatabase_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabase_STATUS(gens map[string]gopter.Gen) {
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

func Test_Database_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Database_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabase_Spec, Database_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabase_Spec runs a test to see if a specific instance of Database_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabase_Spec(subject Database_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Database_Spec
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

// Generator of Database_Spec instances for property testing - lazily instantiated by Database_SpecGenerator()
var database_SpecGenerator gopter.Gen

// Database_SpecGenerator returns a generator of Database_Spec instances for property testing.
// We first initialize database_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func Database_SpecGenerator() gopter.Gen {
	if database_SpecGenerator != nil {
		return database_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabase_Spec(generators)
	database_SpecGenerator = gen.Struct(reflect.TypeOf(Database_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabase_Spec(generators)
	AddRelatedPropertyGeneratorsForDatabase_Spec(generators)
	database_SpecGenerator = gen.Struct(reflect.TypeOf(Database_Spec{}), generators)

	return database_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDatabase_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabase_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Charset"] = gen.PtrOf(gen.AlphaString())
	gens["Collation"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForDatabase_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabase_Spec(gens map[string]gopter.Gen) {
	gens["OperatorSpec"] = gen.PtrOf(DatabaseOperatorSpecGenerator())
}
