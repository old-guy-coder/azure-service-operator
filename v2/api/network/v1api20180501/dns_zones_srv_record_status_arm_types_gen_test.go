// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501

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

func Test_DnsZonesSRVRecord_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesSRVRecord_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesSRVRecord_STATUS_ARM, DnsZonesSRVRecord_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesSRVRecord_STATUS_ARM runs a test to see if a specific instance of DnsZonesSRVRecord_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesSRVRecord_STATUS_ARM(subject DnsZonesSRVRecord_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesSRVRecord_STATUS_ARM
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

// Generator of DnsZonesSRVRecord_STATUS_ARM instances for property testing - lazily instantiated by
// DnsZonesSRVRecord_STATUS_ARMGenerator()
var dnsZonesSRVRecord_STATUS_ARMGenerator gopter.Gen

// DnsZonesSRVRecord_STATUS_ARMGenerator returns a generator of DnsZonesSRVRecord_STATUS_ARM instances for property testing.
// We first initialize dnsZonesSRVRecord_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZonesSRVRecord_STATUS_ARMGenerator() gopter.Gen {
	if dnsZonesSRVRecord_STATUS_ARMGenerator != nil {
		return dnsZonesSRVRecord_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_STATUS_ARM(generators)
	dnsZonesSRVRecord_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsZonesSRVRecord_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsZonesSRVRecord_STATUS_ARM(generators)
	dnsZonesSRVRecord_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsZonesSRVRecord_STATUS_ARM{}), generators)

	return dnsZonesSRVRecord_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZonesSRVRecord_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsZonesSRVRecord_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesSRVRecord_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RecordSetProperties_STATUS_ARMGenerator())
}
