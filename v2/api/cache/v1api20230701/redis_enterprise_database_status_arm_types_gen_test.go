// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20230701

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

func Test_DatabaseProperties_GeoReplication_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_GeoReplication_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties_GeoReplication_STATUS_ARM, DatabaseProperties_GeoReplication_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties_GeoReplication_STATUS_ARM runs a test to see if a specific instance of DatabaseProperties_GeoReplication_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties_GeoReplication_STATUS_ARM(subject DatabaseProperties_GeoReplication_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_GeoReplication_STATUS_ARM
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

// Generator of DatabaseProperties_GeoReplication_STATUS_ARM instances for property testing - lazily instantiated by
// DatabaseProperties_GeoReplication_STATUS_ARMGenerator()
var databaseProperties_GeoReplication_STATUS_ARMGenerator gopter.Gen

// DatabaseProperties_GeoReplication_STATUS_ARMGenerator returns a generator of DatabaseProperties_GeoReplication_STATUS_ARM instances for property testing.
// We first initialize databaseProperties_GeoReplication_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseProperties_GeoReplication_STATUS_ARMGenerator() gopter.Gen {
	if databaseProperties_GeoReplication_STATUS_ARMGenerator != nil {
		return databaseProperties_GeoReplication_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_GeoReplication_STATUS_ARM(generators)
	databaseProperties_GeoReplication_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_GeoReplication_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_GeoReplication_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseProperties_GeoReplication_STATUS_ARM(generators)
	databaseProperties_GeoReplication_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_GeoReplication_STATUS_ARM{}), generators)

	return databaseProperties_GeoReplication_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties_GeoReplication_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties_GeoReplication_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["GroupNickname"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDatabaseProperties_GeoReplication_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseProperties_GeoReplication_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["LinkedDatabases"] = gen.SliceOf(LinkedDatabase_STATUS_ARMGenerator())
}

func Test_DatabaseProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DatabaseProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDatabaseProperties_STATUS_ARM, DatabaseProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDatabaseProperties_STATUS_ARM runs a test to see if a specific instance of DatabaseProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDatabaseProperties_STATUS_ARM(subject DatabaseProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DatabaseProperties_STATUS_ARM
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

// Generator of DatabaseProperties_STATUS_ARM instances for property testing - lazily instantiated by
// DatabaseProperties_STATUS_ARMGenerator()
var databaseProperties_STATUS_ARMGenerator gopter.Gen

// DatabaseProperties_STATUS_ARMGenerator returns a generator of DatabaseProperties_STATUS_ARM instances for property testing.
// We first initialize databaseProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DatabaseProperties_STATUS_ARMGenerator() gopter.Gen {
	if databaseProperties_STATUS_ARMGenerator != nil {
		return databaseProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_STATUS_ARM(generators)
	databaseProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDatabaseProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDatabaseProperties_STATUS_ARM(generators)
	databaseProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DatabaseProperties_STATUS_ARM{}), generators)

	return databaseProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDatabaseProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDatabaseProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ClientProtocol"] = gen.PtrOf(gen.OneConstOf(DatabaseProperties_ClientProtocol_STATUS_ARM_Encrypted, DatabaseProperties_ClientProtocol_STATUS_ARM_Plaintext))
	gens["ClusteringPolicy"] = gen.PtrOf(gen.OneConstOf(DatabaseProperties_ClusteringPolicy_STATUS_ARM_EnterpriseCluster, DatabaseProperties_ClusteringPolicy_STATUS_ARM_OSSCluster))
	gens["EvictionPolicy"] = gen.PtrOf(gen.OneConstOf(
		DatabaseProperties_EvictionPolicy_STATUS_ARM_AllKeysLFU,
		DatabaseProperties_EvictionPolicy_STATUS_ARM_AllKeysLRU,
		DatabaseProperties_EvictionPolicy_STATUS_ARM_AllKeysRandom,
		DatabaseProperties_EvictionPolicy_STATUS_ARM_NoEviction,
		DatabaseProperties_EvictionPolicy_STATUS_ARM_VolatileLFU,
		DatabaseProperties_EvictionPolicy_STATUS_ARM_VolatileLRU,
		DatabaseProperties_EvictionPolicy_STATUS_ARM_VolatileRandom,
		DatabaseProperties_EvictionPolicy_STATUS_ARM_VolatileTTL))
	gens["Port"] = gen.PtrOf(gen.Int())
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		ProvisioningState_STATUS_ARM_Canceled,
		ProvisioningState_STATUS_ARM_Creating,
		ProvisioningState_STATUS_ARM_Deleting,
		ProvisioningState_STATUS_ARM_Failed,
		ProvisioningState_STATUS_ARM_Succeeded,
		ProvisioningState_STATUS_ARM_Updating))
	gens["ResourceState"] = gen.PtrOf(gen.OneConstOf(
		ResourceState_STATUS_ARM_CreateFailed,
		ResourceState_STATUS_ARM_Creating,
		ResourceState_STATUS_ARM_DeleteFailed,
		ResourceState_STATUS_ARM_Deleting,
		ResourceState_STATUS_ARM_DisableFailed,
		ResourceState_STATUS_ARM_Disabled,
		ResourceState_STATUS_ARM_Disabling,
		ResourceState_STATUS_ARM_EnableFailed,
		ResourceState_STATUS_ARM_Enabling,
		ResourceState_STATUS_ARM_Running,
		ResourceState_STATUS_ARM_UpdateFailed,
		ResourceState_STATUS_ARM_Updating))
}

// AddRelatedPropertyGeneratorsForDatabaseProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDatabaseProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["GeoReplication"] = gen.PtrOf(DatabaseProperties_GeoReplication_STATUS_ARMGenerator())
	gens["Modules"] = gen.SliceOf(Module_STATUS_ARMGenerator())
	gens["Persistence"] = gen.PtrOf(Persistence_STATUS_ARMGenerator())
}

func Test_LinkedDatabase_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LinkedDatabase_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLinkedDatabase_STATUS_ARM, LinkedDatabase_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLinkedDatabase_STATUS_ARM runs a test to see if a specific instance of LinkedDatabase_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLinkedDatabase_STATUS_ARM(subject LinkedDatabase_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LinkedDatabase_STATUS_ARM
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

// Generator of LinkedDatabase_STATUS_ARM instances for property testing - lazily instantiated by
// LinkedDatabase_STATUS_ARMGenerator()
var linkedDatabase_STATUS_ARMGenerator gopter.Gen

// LinkedDatabase_STATUS_ARMGenerator returns a generator of LinkedDatabase_STATUS_ARM instances for property testing.
func LinkedDatabase_STATUS_ARMGenerator() gopter.Gen {
	if linkedDatabase_STATUS_ARMGenerator != nil {
		return linkedDatabase_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLinkedDatabase_STATUS_ARM(generators)
	linkedDatabase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(LinkedDatabase_STATUS_ARM{}), generators)

	return linkedDatabase_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForLinkedDatabase_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLinkedDatabase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["State"] = gen.PtrOf(gen.OneConstOf(
		LinkedDatabase_State_STATUS_ARM_LinkFailed,
		LinkedDatabase_State_STATUS_ARM_Linked,
		LinkedDatabase_State_STATUS_ARM_Linking,
		LinkedDatabase_State_STATUS_ARM_UnlinkFailed,
		LinkedDatabase_State_STATUS_ARM_Unlinking))
}

func Test_Module_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Module_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForModule_STATUS_ARM, Module_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForModule_STATUS_ARM runs a test to see if a specific instance of Module_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForModule_STATUS_ARM(subject Module_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Module_STATUS_ARM
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

// Generator of Module_STATUS_ARM instances for property testing - lazily instantiated by Module_STATUS_ARMGenerator()
var module_STATUS_ARMGenerator gopter.Gen

// Module_STATUS_ARMGenerator returns a generator of Module_STATUS_ARM instances for property testing.
func Module_STATUS_ARMGenerator() gopter.Gen {
	if module_STATUS_ARMGenerator != nil {
		return module_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForModule_STATUS_ARM(generators)
	module_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Module_STATUS_ARM{}), generators)

	return module_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForModule_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForModule_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Args"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Version"] = gen.PtrOf(gen.AlphaString())
}

func Test_Persistence_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of Persistence_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForPersistence_STATUS_ARM, Persistence_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForPersistence_STATUS_ARM runs a test to see if a specific instance of Persistence_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForPersistence_STATUS_ARM(subject Persistence_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual Persistence_STATUS_ARM
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

// Generator of Persistence_STATUS_ARM instances for property testing - lazily instantiated by
// Persistence_STATUS_ARMGenerator()
var persistence_STATUS_ARMGenerator gopter.Gen

// Persistence_STATUS_ARMGenerator returns a generator of Persistence_STATUS_ARM instances for property testing.
func Persistence_STATUS_ARMGenerator() gopter.Gen {
	if persistence_STATUS_ARMGenerator != nil {
		return persistence_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForPersistence_STATUS_ARM(generators)
	persistence_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(Persistence_STATUS_ARM{}), generators)

	return persistence_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForPersistence_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForPersistence_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["AofEnabled"] = gen.PtrOf(gen.Bool())
	gens["AofFrequency"] = gen.PtrOf(gen.OneConstOf(Persistence_AofFrequency_STATUS_ARM_1S, Persistence_AofFrequency_STATUS_ARM_Always))
	gens["RdbEnabled"] = gen.PtrOf(gen.Bool())
	gens["RdbFrequency"] = gen.PtrOf(gen.OneConstOf(Persistence_RdbFrequency_STATUS_ARM_12H, Persistence_RdbFrequency_STATUS_ARM_1H, Persistence_RdbFrequency_STATUS_ARM_6H))
}

func Test_RedisEnterpriseDatabase_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RedisEnterpriseDatabase_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRedisEnterpriseDatabase_STATUS_ARM, RedisEnterpriseDatabase_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRedisEnterpriseDatabase_STATUS_ARM runs a test to see if a specific instance of RedisEnterpriseDatabase_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForRedisEnterpriseDatabase_STATUS_ARM(subject RedisEnterpriseDatabase_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RedisEnterpriseDatabase_STATUS_ARM
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

// Generator of RedisEnterpriseDatabase_STATUS_ARM instances for property testing - lazily instantiated by
// RedisEnterpriseDatabase_STATUS_ARMGenerator()
var redisEnterpriseDatabase_STATUS_ARMGenerator gopter.Gen

// RedisEnterpriseDatabase_STATUS_ARMGenerator returns a generator of RedisEnterpriseDatabase_STATUS_ARM instances for property testing.
// We first initialize redisEnterpriseDatabase_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RedisEnterpriseDatabase_STATUS_ARMGenerator() gopter.Gen {
	if redisEnterpriseDatabase_STATUS_ARMGenerator != nil {
		return redisEnterpriseDatabase_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_STATUS_ARM(generators)
	redisEnterpriseDatabase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_STATUS_ARM(generators)
	redisEnterpriseDatabase_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(RedisEnterpriseDatabase_STATUS_ARM{}), generators)

	return redisEnterpriseDatabase_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRedisEnterpriseDatabase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRedisEnterpriseDatabase_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(DatabaseProperties_STATUS_ARMGenerator())
}
