// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20210515

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

func Test_AutoUpgradePolicyResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoUpgradePolicyResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoUpgradePolicyResource_ARM, AutoUpgradePolicyResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoUpgradePolicyResource_ARM runs a test to see if a specific instance of AutoUpgradePolicyResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoUpgradePolicyResource_ARM(subject AutoUpgradePolicyResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoUpgradePolicyResource_ARM
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

// Generator of AutoUpgradePolicyResource_ARM instances for property testing - lazily instantiated by
// AutoUpgradePolicyResource_ARMGenerator()
var autoUpgradePolicyResource_ARMGenerator gopter.Gen

// AutoUpgradePolicyResource_ARMGenerator returns a generator of AutoUpgradePolicyResource_ARM instances for property testing.
func AutoUpgradePolicyResource_ARMGenerator() gopter.Gen {
	if autoUpgradePolicyResource_ARMGenerator != nil {
		return autoUpgradePolicyResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource_ARM(generators)
	autoUpgradePolicyResource_ARMGenerator = gen.Struct(reflect.TypeOf(AutoUpgradePolicyResource_ARM{}), generators)

	return autoUpgradePolicyResource_ARMGenerator
}

// AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoUpgradePolicyResource_ARM(gens map[string]gopter.Gen) {
	gens["ThroughputPolicy"] = gen.PtrOf(ThroughputPolicyResource_ARMGenerator())
}

func Test_AutoscaleSettingsResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of AutoscaleSettingsResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForAutoscaleSettingsResource_ARM, AutoscaleSettingsResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForAutoscaleSettingsResource_ARM runs a test to see if a specific instance of AutoscaleSettingsResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForAutoscaleSettingsResource_ARM(subject AutoscaleSettingsResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual AutoscaleSettingsResource_ARM
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

// Generator of AutoscaleSettingsResource_ARM instances for property testing - lazily instantiated by
// AutoscaleSettingsResource_ARMGenerator()
var autoscaleSettingsResource_ARMGenerator gopter.Gen

// AutoscaleSettingsResource_ARMGenerator returns a generator of AutoscaleSettingsResource_ARM instances for property testing.
// We first initialize autoscaleSettingsResource_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func AutoscaleSettingsResource_ARMGenerator() gopter.Gen {
	if autoscaleSettingsResource_ARMGenerator != nil {
		return autoscaleSettingsResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_ARM(generators)
	autoscaleSettingsResource_ARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_ARM(generators)
	AddRelatedPropertyGeneratorsForAutoscaleSettingsResource_ARM(generators)
	autoscaleSettingsResource_ARMGenerator = gen.Struct(reflect.TypeOf(AutoscaleSettingsResource_ARM{}), generators)

	return autoscaleSettingsResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForAutoscaleSettingsResource_ARM(gens map[string]gopter.Gen) {
	gens["MaxThroughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForAutoscaleSettingsResource_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForAutoscaleSettingsResource_ARM(gens map[string]gopter.Gen) {
	gens["AutoUpgradePolicy"] = gen.PtrOf(AutoUpgradePolicyResource_ARMGenerator())
}

func Test_MongodbDatabaseCollectionThroughputSetting_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of MongodbDatabaseCollectionThroughputSetting_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting_Spec_ARM, MongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting_Spec_ARM runs a test to see if a specific instance of MongodbDatabaseCollectionThroughputSetting_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForMongodbDatabaseCollectionThroughputSetting_Spec_ARM(subject MongodbDatabaseCollectionThroughputSetting_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual MongodbDatabaseCollectionThroughputSetting_Spec_ARM
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

// Generator of MongodbDatabaseCollectionThroughputSetting_Spec_ARM instances for property testing - lazily instantiated
// by MongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator()
var mongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator gopter.Gen

// MongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator returns a generator of MongodbDatabaseCollectionThroughputSetting_Spec_ARM instances for property testing.
// We first initialize mongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func MongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator() gopter.Gen {
	if mongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator != nil {
		return mongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec_ARM(generators)
	mongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollectionThroughputSetting_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec_ARM(generators)
	mongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(MongodbDatabaseCollectionThroughputSetting_Spec_ARM{}), generators)

	return mongodbDatabaseCollectionThroughputSetting_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.AlphaString()
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForMongodbDatabaseCollectionThroughputSetting_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(ThroughputSettingsUpdateProperties_ARMGenerator())
}

func Test_ThroughputPolicyResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputPolicyResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputPolicyResource_ARM, ThroughputPolicyResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputPolicyResource_ARM runs a test to see if a specific instance of ThroughputPolicyResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputPolicyResource_ARM(subject ThroughputPolicyResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputPolicyResource_ARM
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

// Generator of ThroughputPolicyResource_ARM instances for property testing - lazily instantiated by
// ThroughputPolicyResource_ARMGenerator()
var throughputPolicyResource_ARMGenerator gopter.Gen

// ThroughputPolicyResource_ARMGenerator returns a generator of ThroughputPolicyResource_ARM instances for property testing.
func ThroughputPolicyResource_ARMGenerator() gopter.Gen {
	if throughputPolicyResource_ARMGenerator != nil {
		return throughputPolicyResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputPolicyResource_ARM(generators)
	throughputPolicyResource_ARMGenerator = gen.Struct(reflect.TypeOf(ThroughputPolicyResource_ARM{}), generators)

	return throughputPolicyResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputPolicyResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputPolicyResource_ARM(gens map[string]gopter.Gen) {
	gens["IncrementPercent"] = gen.PtrOf(gen.Int())
	gens["IsEnabled"] = gen.PtrOf(gen.Bool())
}

func Test_ThroughputSettingsResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsResource_ARM, ThroughputSettingsResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsResource_ARM runs a test to see if a specific instance of ThroughputSettingsResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsResource_ARM(subject ThroughputSettingsResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsResource_ARM
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

// Generator of ThroughputSettingsResource_ARM instances for property testing - lazily instantiated by
// ThroughputSettingsResource_ARMGenerator()
var throughputSettingsResource_ARMGenerator gopter.Gen

// ThroughputSettingsResource_ARMGenerator returns a generator of ThroughputSettingsResource_ARM instances for property testing.
// We first initialize throughputSettingsResource_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func ThroughputSettingsResource_ARMGenerator() gopter.Gen {
	if throughputSettingsResource_ARMGenerator != nil {
		return throughputSettingsResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResource_ARM(generators)
	throughputSettingsResource_ARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResource_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForThroughputSettingsResource_ARM(generators)
	AddRelatedPropertyGeneratorsForThroughputSettingsResource_ARM(generators)
	throughputSettingsResource_ARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsResource_ARM{}), generators)

	return throughputSettingsResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForThroughputSettingsResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForThroughputSettingsResource_ARM(gens map[string]gopter.Gen) {
	gens["Throughput"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForThroughputSettingsResource_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsResource_ARM(gens map[string]gopter.Gen) {
	gens["AutoscaleSettings"] = gen.PtrOf(AutoscaleSettingsResource_ARMGenerator())
}

func Test_ThroughputSettingsUpdateProperties_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of ThroughputSettingsUpdateProperties_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForThroughputSettingsUpdateProperties_ARM, ThroughputSettingsUpdateProperties_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForThroughputSettingsUpdateProperties_ARM runs a test to see if a specific instance of ThroughputSettingsUpdateProperties_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForThroughputSettingsUpdateProperties_ARM(subject ThroughputSettingsUpdateProperties_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual ThroughputSettingsUpdateProperties_ARM
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

// Generator of ThroughputSettingsUpdateProperties_ARM instances for property testing - lazily instantiated by
// ThroughputSettingsUpdateProperties_ARMGenerator()
var throughputSettingsUpdateProperties_ARMGenerator gopter.Gen

// ThroughputSettingsUpdateProperties_ARMGenerator returns a generator of ThroughputSettingsUpdateProperties_ARM instances for property testing.
func ThroughputSettingsUpdateProperties_ARMGenerator() gopter.Gen {
	if throughputSettingsUpdateProperties_ARMGenerator != nil {
		return throughputSettingsUpdateProperties_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForThroughputSettingsUpdateProperties_ARM(generators)
	throughputSettingsUpdateProperties_ARMGenerator = gen.Struct(reflect.TypeOf(ThroughputSettingsUpdateProperties_ARM{}), generators)

	return throughputSettingsUpdateProperties_ARMGenerator
}

// AddRelatedPropertyGeneratorsForThroughputSettingsUpdateProperties_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForThroughputSettingsUpdateProperties_ARM(gens map[string]gopter.Gen) {
	gens["Resource"] = gen.PtrOf(ThroughputSettingsResource_ARMGenerator())
}