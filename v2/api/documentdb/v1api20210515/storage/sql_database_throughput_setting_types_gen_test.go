// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	storage "github.com/Azure/azure-service-operator/v2/api/documentdb/v1api20231115/storage"
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

func Test_SqlDatabaseThroughputSetting_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseThroughputSetting to hub returns original",
		prop.ForAll(RunResourceConversionTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForSqlDatabaseThroughputSetting tests if a specific instance of SqlDatabaseThroughputSetting round trips to the hub storage version and back losslessly
func RunResourceConversionTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub storage.SqlDatabaseThroughputSetting
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual SqlDatabaseThroughputSetting
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

func Test_SqlDatabaseThroughputSetting_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseThroughputSetting to SqlDatabaseThroughputSetting via AssignProperties_To_SqlDatabaseThroughputSetting & AssignProperties_From_SqlDatabaseThroughputSetting returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseThroughputSetting tests if a specific instance of SqlDatabaseThroughputSetting can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlDatabaseThroughputSetting
	err := copied.AssignProperties_To_SqlDatabaseThroughputSetting(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseThroughputSetting
	err = actual.AssignProperties_From_SqlDatabaseThroughputSetting(&other)
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

func Test_SqlDatabaseThroughputSetting_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseThroughputSetting via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseThroughputSetting, SqlDatabaseThroughputSettingGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseThroughputSetting runs a test to see if a specific instance of SqlDatabaseThroughputSetting round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseThroughputSetting(subject SqlDatabaseThroughputSetting) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseThroughputSetting
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

// Generator of SqlDatabaseThroughputSetting instances for property testing - lazily instantiated by
// SqlDatabaseThroughputSettingGenerator()
var sqlDatabaseThroughputSettingGenerator gopter.Gen

// SqlDatabaseThroughputSettingGenerator returns a generator of SqlDatabaseThroughputSetting instances for property testing.
func SqlDatabaseThroughputSettingGenerator() gopter.Gen {
	if sqlDatabaseThroughputSettingGenerator != nil {
		return sqlDatabaseThroughputSettingGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(generators)
	sqlDatabaseThroughputSettingGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseThroughputSetting{}), generators)

	return sqlDatabaseThroughputSettingGenerator
}

// AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting(gens map[string]gopter.Gen) {
	gens["Spec"] = SqlDatabaseThroughputSetting_SpecGenerator()
	gens["Status"] = SqlDatabaseThroughputSetting_STATUSGenerator()
}

func Test_SqlDatabaseThroughputSetting_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseThroughputSetting_STATUS to SqlDatabaseThroughputSetting_STATUS via AssignProperties_To_SqlDatabaseThroughputSetting_STATUS & AssignProperties_From_SqlDatabaseThroughputSetting_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseThroughputSetting_STATUS, SqlDatabaseThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseThroughputSetting_STATUS tests if a specific instance of SqlDatabaseThroughputSetting_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseThroughputSetting_STATUS(subject SqlDatabaseThroughputSetting_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlDatabaseThroughputSetting_STATUS
	err := copied.AssignProperties_To_SqlDatabaseThroughputSetting_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseThroughputSetting_STATUS
	err = actual.AssignProperties_From_SqlDatabaseThroughputSetting_STATUS(&other)
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

func Test_SqlDatabaseThroughputSetting_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseThroughputSetting_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseThroughputSetting_STATUS, SqlDatabaseThroughputSetting_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseThroughputSetting_STATUS runs a test to see if a specific instance of SqlDatabaseThroughputSetting_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseThroughputSetting_STATUS(subject SqlDatabaseThroughputSetting_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseThroughputSetting_STATUS
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

// Generator of SqlDatabaseThroughputSetting_STATUS instances for property testing - lazily instantiated by
// SqlDatabaseThroughputSetting_STATUSGenerator()
var sqlDatabaseThroughputSetting_STATUSGenerator gopter.Gen

// SqlDatabaseThroughputSetting_STATUSGenerator returns a generator of SqlDatabaseThroughputSetting_STATUS instances for property testing.
// We first initialize sqlDatabaseThroughputSetting_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseThroughputSetting_STATUSGenerator() gopter.Gen {
	if sqlDatabaseThroughputSetting_STATUSGenerator != nil {
		return sqlDatabaseThroughputSetting_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseThroughputSetting_STATUS(generators)
	sqlDatabaseThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseThroughputSetting_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseThroughputSetting_STATUS(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting_STATUS(generators)
	sqlDatabaseThroughputSetting_STATUSGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseThroughputSetting_STATUS{}), generators)

	return sqlDatabaseThroughputSetting_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseThroughputSetting_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting_STATUS(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsGetProperties_Resource_STATUSGenerator())
}

func Test_SqlDatabaseThroughputSetting_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from SqlDatabaseThroughputSetting_Spec to SqlDatabaseThroughputSetting_Spec via AssignProperties_To_SqlDatabaseThroughputSetting_Spec & AssignProperties_From_SqlDatabaseThroughputSetting_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForSqlDatabaseThroughputSetting_Spec, SqlDatabaseThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForSqlDatabaseThroughputSetting_Spec tests if a specific instance of SqlDatabaseThroughputSetting_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForSqlDatabaseThroughputSetting_Spec(subject SqlDatabaseThroughputSetting_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other storage.SqlDatabaseThroughputSetting_Spec
	err := copied.AssignProperties_To_SqlDatabaseThroughputSetting_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual SqlDatabaseThroughputSetting_Spec
	err = actual.AssignProperties_From_SqlDatabaseThroughputSetting_Spec(&other)
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

func Test_SqlDatabaseThroughputSetting_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SqlDatabaseThroughputSetting_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSqlDatabaseThroughputSetting_Spec, SqlDatabaseThroughputSetting_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSqlDatabaseThroughputSetting_Spec runs a test to see if a specific instance of SqlDatabaseThroughputSetting_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForSqlDatabaseThroughputSetting_Spec(subject SqlDatabaseThroughputSetting_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SqlDatabaseThroughputSetting_Spec
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

// Generator of SqlDatabaseThroughputSetting_Spec instances for property testing - lazily instantiated by
// SqlDatabaseThroughputSetting_SpecGenerator()
var sqlDatabaseThroughputSetting_SpecGenerator gopter.Gen

// SqlDatabaseThroughputSetting_SpecGenerator returns a generator of SqlDatabaseThroughputSetting_Spec instances for property testing.
// We first initialize sqlDatabaseThroughputSetting_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func SqlDatabaseThroughputSetting_SpecGenerator() gopter.Gen {
	if sqlDatabaseThroughputSetting_SpecGenerator != nil {
		return sqlDatabaseThroughputSetting_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseThroughputSetting_Spec(generators)
	sqlDatabaseThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseThroughputSetting_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSqlDatabaseThroughputSetting_Spec(generators)
	AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting_Spec(generators)
	sqlDatabaseThroughputSetting_SpecGenerator = gen.Struct(reflect.TypeOf(SqlDatabaseThroughputSetting_Spec{}), generators)

	return sqlDatabaseThroughputSetting_SpecGenerator
}

// AddIndependentPropertyGeneratorsForSqlDatabaseThroughputSetting_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSqlDatabaseThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForSqlDatabaseThroughputSetting_Spec(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResourceGenerator())
}
