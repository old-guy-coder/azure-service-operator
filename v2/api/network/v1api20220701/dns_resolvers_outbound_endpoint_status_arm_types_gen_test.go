// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20220701

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

func Test_DnsResolversOutboundEndpoint_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsResolversOutboundEndpoint_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsResolversOutboundEndpoint_STATUS_ARM, DnsResolversOutboundEndpoint_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsResolversOutboundEndpoint_STATUS_ARM runs a test to see if a specific instance of DnsResolversOutboundEndpoint_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsResolversOutboundEndpoint_STATUS_ARM(subject DnsResolversOutboundEndpoint_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsResolversOutboundEndpoint_STATUS_ARM
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

// Generator of DnsResolversOutboundEndpoint_STATUS_ARM instances for property testing - lazily instantiated by
// DnsResolversOutboundEndpoint_STATUS_ARMGenerator()
var dnsResolversOutboundEndpoint_STATUS_ARMGenerator gopter.Gen

// DnsResolversOutboundEndpoint_STATUS_ARMGenerator returns a generator of DnsResolversOutboundEndpoint_STATUS_ARM instances for property testing.
// We first initialize dnsResolversOutboundEndpoint_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsResolversOutboundEndpoint_STATUS_ARMGenerator() gopter.Gen {
	if dnsResolversOutboundEndpoint_STATUS_ARMGenerator != nil {
		return dnsResolversOutboundEndpoint_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS_ARM(generators)
	dnsResolversOutboundEndpoint_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsResolversOutboundEndpoint_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS_ARM(generators)
	dnsResolversOutboundEndpoint_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(DnsResolversOutboundEndpoint_STATUS_ARM{}), generators)

	return dnsResolversOutboundEndpoint_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Location"] = gen.PtrOf(gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["Tags"] = gen.MapOf(
		gen.AlphaString(),
		gen.AlphaString())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsResolversOutboundEndpoint_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(OutboundEndpointProperties_STATUS_ARMGenerator())
	gens["SystemData"] = gen.PtrOf(SystemData_STATUS_ARMGenerator())
}

func Test_OutboundEndpointProperties_STATUS_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of OutboundEndpointProperties_STATUS_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForOutboundEndpointProperties_STATUS_ARM, OutboundEndpointProperties_STATUS_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForOutboundEndpointProperties_STATUS_ARM runs a test to see if a specific instance of OutboundEndpointProperties_STATUS_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForOutboundEndpointProperties_STATUS_ARM(subject OutboundEndpointProperties_STATUS_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual OutboundEndpointProperties_STATUS_ARM
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

// Generator of OutboundEndpointProperties_STATUS_ARM instances for property testing - lazily instantiated by
// OutboundEndpointProperties_STATUS_ARMGenerator()
var outboundEndpointProperties_STATUS_ARMGenerator gopter.Gen

// OutboundEndpointProperties_STATUS_ARMGenerator returns a generator of OutboundEndpointProperties_STATUS_ARM instances for property testing.
// We first initialize outboundEndpointProperties_STATUS_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func OutboundEndpointProperties_STATUS_ARMGenerator() gopter.Gen {
	if outboundEndpointProperties_STATUS_ARMGenerator != nil {
		return outboundEndpointProperties_STATUS_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOutboundEndpointProperties_STATUS_ARM(generators)
	outboundEndpointProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(OutboundEndpointProperties_STATUS_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForOutboundEndpointProperties_STATUS_ARM(generators)
	AddRelatedPropertyGeneratorsForOutboundEndpointProperties_STATUS_ARM(generators)
	outboundEndpointProperties_STATUS_ARMGenerator = gen.Struct(reflect.TypeOf(OutboundEndpointProperties_STATUS_ARM{}), generators)

	return outboundEndpointProperties_STATUS_ARMGenerator
}

// AddIndependentPropertyGeneratorsForOutboundEndpointProperties_STATUS_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForOutboundEndpointProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["ProvisioningState"] = gen.PtrOf(gen.OneConstOf(
		DnsresolverProvisioningState_STATUS_ARM_Canceled,
		DnsresolverProvisioningState_STATUS_ARM_Creating,
		DnsresolverProvisioningState_STATUS_ARM_Deleting,
		DnsresolverProvisioningState_STATUS_ARM_Failed,
		DnsresolverProvisioningState_STATUS_ARM_Succeeded,
		DnsresolverProvisioningState_STATUS_ARM_Updating))
	gens["ResourceGuid"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForOutboundEndpointProperties_STATUS_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForOutboundEndpointProperties_STATUS_ARM(gens map[string]gopter.Gen) {
	gens["Subnet"] = gen.PtrOf(DnsresolverSubResource_STATUS_ARMGenerator())
}
