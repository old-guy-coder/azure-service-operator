// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package storage

import (
	"encoding/json"
	v20230401s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230401/storage"
	v20230801s "github.com/Azure/azure-service-operator/v2/api/cache/v1api20230801/storage"
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

func Test_RedisPatchSchedule_WhenConvertedToHub_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	parameters.MinSuccessfulTests = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisPatchSchedule to hub returns original",
		prop.ForAll(RunResourceConversionTestForRedisPatchSchedule, RedisPatchScheduleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunResourceConversionTestForRedisPatchSchedule tests if a specific instance of RedisPatchSchedule round trips to the hub storage version and back losslessly
func RunResourceConversionTestForRedisPatchSchedule(subject RedisPatchSchedule) string {
	// Copy subject to make sure conversion doesn't modify it
	copied := subject.DeepCopy()

	// Convert to our hub version
	var hub v20230801s.RedisPatchSchedule
	err := copied.ConvertTo(&hub)
	if err != nil {
		return err.Error()
	}

	// Convert from our hub version
	var actual RedisPatchSchedule
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

func Test_RedisPatchSchedule_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisPatchSchedule to RedisPatchSchedule via AssignProperties_To_RedisPatchSchedule & AssignProperties_From_RedisPatchSchedule returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisPatchSchedule, RedisPatchScheduleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisPatchSchedule tests if a specific instance of RedisPatchSchedule can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRedisPatchSchedule(subject RedisPatchSchedule) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230401s.RedisPatchSchedule
	err := copied.AssignProperties_To_RedisPatchSchedule(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisPatchSchedule
	err = actual.AssignProperties_From_RedisPatchSchedule(&other)
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

func Test_RedisPatchSchedule_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisPatchSchedule via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchSchedule, RedisPatchScheduleGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchSchedule runs a test to see if a specific instance of RedisPatchSchedule round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchSchedule(subject RedisPatchSchedule) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisPatchSchedule
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

// Generator of RedisPatchSchedule instances for property testing - lazily instantiated by RedisPatchScheduleGenerator()
var redisPatchScheduleGenerator gopter.Gen

// RedisPatchScheduleGenerator returns a generator of RedisPatchSchedule instances for property testing.
func RedisPatchScheduleGenerator() gopter.Gen {
	if redisPatchScheduleGenerator != nil {
		return redisPatchScheduleGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForRedisPatchSchedule(generators)
	redisPatchScheduleGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule{}), generators)

	return redisPatchScheduleGenerator
}

// AddRelatedPropertyGeneratorsForRedisPatchSchedule is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchSchedule(gens map[string]gopter.Gen) {
	gens["Spec"] = RedisPatchSchedule_SpecGenerator()
	gens["Status"] = RedisPatchSchedule_STATUSGenerator()
}

func Test_RedisPatchSchedule_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisPatchSchedule_STATUS to RedisPatchSchedule_STATUS via AssignProperties_To_RedisPatchSchedule_STATUS & AssignProperties_From_RedisPatchSchedule_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisPatchSchedule_STATUS, RedisPatchSchedule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisPatchSchedule_STATUS tests if a specific instance of RedisPatchSchedule_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRedisPatchSchedule_STATUS(subject RedisPatchSchedule_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230401s.RedisPatchSchedule_STATUS
	err := copied.AssignProperties_To_RedisPatchSchedule_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisPatchSchedule_STATUS
	err = actual.AssignProperties_From_RedisPatchSchedule_STATUS(&other)
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

func Test_RedisPatchSchedule_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisPatchSchedule_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchSchedule_STATUS, RedisPatchSchedule_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchSchedule_STATUS runs a test to see if a specific instance of RedisPatchSchedule_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchSchedule_STATUS(subject RedisPatchSchedule_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisPatchSchedule_STATUS
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

// Generator of RedisPatchSchedule_STATUS instances for property testing - lazily instantiated by
// RedisPatchSchedule_STATUSGenerator()
var redisPatchSchedule_STATUSGenerator gopter.Gen

// RedisPatchSchedule_STATUSGenerator returns a generator of RedisPatchSchedule_STATUS instances for property testing.
// We first initialize redisPatchSchedule_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPatchSchedule_STATUSGenerator() gopter.Gen {
	if redisPatchSchedule_STATUSGenerator != nil {
		return redisPatchSchedule_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedule_STATUS(generators)
	redisPatchSchedule_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedule_STATUS(generators)
	AddRelatedPropertyGeneratorsForRedisPatchSchedule_STATUS(generators)
	redisPatchSchedule_STATUSGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_STATUS{}), generators)

	return redisPatchSchedule_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForRedisPatchSchedule_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPatchSchedule_STATUS(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisPatchSchedule_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchSchedule_STATUS(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntry_STATUSGenerator())
}

func Test_RedisPatchSchedule_Spec_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from RedisPatchSchedule_Spec to RedisPatchSchedule_Spec via AssignProperties_To_RedisPatchSchedule_Spec & AssignProperties_From_RedisPatchSchedule_Spec returns original",
		prop.ForAll(RunPropertyAssignmentTestForRedisPatchSchedule_Spec, RedisPatchSchedule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForRedisPatchSchedule_Spec tests if a specific instance of RedisPatchSchedule_Spec can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForRedisPatchSchedule_Spec(subject RedisPatchSchedule_Spec) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230401s.RedisPatchSchedule_Spec
	err := copied.AssignProperties_To_RedisPatchSchedule_Spec(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual RedisPatchSchedule_Spec
	err = actual.AssignProperties_From_RedisPatchSchedule_Spec(&other)
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

func Test_RedisPatchSchedule_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisPatchSchedule_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisPatchSchedule_Spec, RedisPatchSchedule_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisPatchSchedule_Spec runs a test to see if a specific instance of RedisPatchSchedule_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisPatchSchedule_Spec(subject RedisPatchSchedule_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisPatchSchedule_Spec
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

// Generator of RedisPatchSchedule_Spec instances for property testing - lazily instantiated by
// RedisPatchSchedule_SpecGenerator()
var redisPatchSchedule_SpecGenerator gopter.Gen

// RedisPatchSchedule_SpecGenerator returns a generator of RedisPatchSchedule_Spec instances for property testing.
// We first initialize redisPatchSchedule_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisPatchSchedule_SpecGenerator() gopter.Gen {
	if redisPatchSchedule_SpecGenerator != nil {
		return redisPatchSchedule_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedule_Spec(generators)
	redisPatchSchedule_SpecGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisPatchSchedule_Spec(generators)
	AddRelatedPropertyGeneratorsForRedisPatchSchedule_Spec(generators)
	redisPatchSchedule_SpecGenerator = gen.Struct(reflect.TypeOf(RedisPatchSchedule_Spec{}), generators)

	return redisPatchSchedule_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRedisPatchSchedule_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisPatchSchedule_Spec(gens map[string]gopter.Gen) {
	gens["OriginalVersion"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRedisPatchSchedule_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisPatchSchedule_Spec(gens map[string]gopter.Gen) {
	gens["ScheduleEntries"] = gen.SliceOf(ScheduleEntryGenerator())
}

func Test_ScheduleEntry_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ScheduleEntry to ScheduleEntry via AssignProperties_To_ScheduleEntry & AssignProperties_From_ScheduleEntry returns original",
		prop.ForAll(RunPropertyAssignmentTestForScheduleEntry, ScheduleEntryGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForScheduleEntry tests if a specific instance of ScheduleEntry can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForScheduleEntry(subject ScheduleEntry) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230401s.ScheduleEntry
	err := copied.AssignProperties_To_ScheduleEntry(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ScheduleEntry
	err = actual.AssignProperties_From_ScheduleEntry(&other)
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

func Test_ScheduleEntry_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntry via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntry, ScheduleEntryGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntry runs a test to see if a specific instance of ScheduleEntry round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntry(subject ScheduleEntry) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntry
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

// Generator of ScheduleEntry instances for property testing - lazily instantiated by ScheduleEntryGenerator()
var scheduleEntryGenerator gopter.Gen

// ScheduleEntryGenerator returns a generator of ScheduleEntry instances for property testing.
func ScheduleEntryGenerator() gopter.Gen {
	if scheduleEntryGenerator != nil {
		return scheduleEntryGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduleEntry(generators)
	scheduleEntryGenerator = gen.Struct(reflect.TypeOf(ScheduleEntry{}), generators)

	return scheduleEntryGenerator
}

// AddIndependentPropertyGeneratorsForScheduleEntry is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduleEntry(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.AlphaString())
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}

func Test_ScheduleEntry_STATUS_WhenPropertiesConverted_RoundTripsWithoutLoss(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MaxSize = 10
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip from ScheduleEntry_STATUS to ScheduleEntry_STATUS via AssignProperties_To_ScheduleEntry_STATUS & AssignProperties_From_ScheduleEntry_STATUS returns original",
		prop.ForAll(RunPropertyAssignmentTestForScheduleEntry_STATUS, ScheduleEntry_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(false, 240, os.Stdout))
}

// RunPropertyAssignmentTestForScheduleEntry_STATUS tests if a specific instance of ScheduleEntry_STATUS can be assigned to storage and back losslessly
func RunPropertyAssignmentTestForScheduleEntry_STATUS(subject ScheduleEntry_STATUS) string {
	// Copy subject to make sure assignment doesn't modify it
	copied := subject.DeepCopy()

	// Use AssignPropertiesTo() for the first stage of conversion
	var other v20230401s.ScheduleEntry_STATUS
	err := copied.AssignProperties_To_ScheduleEntry_STATUS(&other)
	if err != nil {
		return err.Error()
	}

	// Use AssignPropertiesFrom() to convert back to our original type
	var actual ScheduleEntry_STATUS
	err = actual.AssignProperties_From_ScheduleEntry_STATUS(&other)
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

func Test_ScheduleEntry_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ScheduleEntry_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForScheduleEntry_STATUS, ScheduleEntry_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForScheduleEntry_STATUS runs a test to see if a specific instance of ScheduleEntry_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForScheduleEntry_STATUS(subject ScheduleEntry_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ScheduleEntry_STATUS
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

// Generator of ScheduleEntry_STATUS instances for property testing - lazily instantiated by
// ScheduleEntry_STATUSGenerator()
var scheduleEntry_STATUSGenerator gopter.Gen

// ScheduleEntry_STATUSGenerator returns a generator of ScheduleEntry_STATUS instances for property testing.
func ScheduleEntry_STATUSGenerator() gopter.Gen {
	if scheduleEntry_STATUSGenerator != nil {
		return scheduleEntry_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForScheduleEntry_STATUS(generators)
	scheduleEntry_STATUSGenerator = gen.Struct(reflect.TypeOf(ScheduleEntry_STATUS{}), generators)

	return scheduleEntry_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForScheduleEntry_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForScheduleEntry_STATUS(gens map[string]gopter.Gen) {
	gens["DayOfWeek"] = gen.PtrOf(gen.AlphaString())
	gens["MaintenanceWindow"] = gen.PtrOf(gen.AlphaString())
	gens["StartHourUtc"] = gen.PtrOf(gen.Int())
}
