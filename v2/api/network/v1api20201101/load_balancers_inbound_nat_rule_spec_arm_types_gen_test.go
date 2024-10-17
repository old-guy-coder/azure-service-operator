// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20201101

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

func Test_InboundNatRulePropertiesFormat_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of InboundNatRulePropertiesFormat_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForInboundNatRulePropertiesFormat_ARM, InboundNatRulePropertiesFormat_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForInboundNatRulePropertiesFormat_ARM runs a test to see if a specific instance of InboundNatRulePropertiesFormat_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForInboundNatRulePropertiesFormat_ARM(subject InboundNatRulePropertiesFormat_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual InboundNatRulePropertiesFormat_ARM
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

// Generator of InboundNatRulePropertiesFormat_ARM instances for property testing - lazily instantiated by
// InboundNatRulePropertiesFormat_ARMGenerator()
var inboundNatRulePropertiesFormat_ARMGenerator gopter.Gen

// InboundNatRulePropertiesFormat_ARMGenerator returns a generator of InboundNatRulePropertiesFormat_ARM instances for property testing.
// We first initialize inboundNatRulePropertiesFormat_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func InboundNatRulePropertiesFormat_ARMGenerator() gopter.Gen {
	if inboundNatRulePropertiesFormat_ARMGenerator != nil {
		return inboundNatRulePropertiesFormat_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_ARM(generators)
	inboundNatRulePropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(InboundNatRulePropertiesFormat_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_ARM(generators)
	AddRelatedPropertyGeneratorsForInboundNatRulePropertiesFormat_ARM(generators)
	inboundNatRulePropertiesFormat_ARMGenerator = gen.Struct(reflect.TypeOf(InboundNatRulePropertiesFormat_ARM{}), generators)

	return inboundNatRulePropertiesFormat_ARMGenerator
}

// AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForInboundNatRulePropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["BackendPort"] = gen.PtrOf(gen.Int())
	gens["EnableFloatingIP"] = gen.PtrOf(gen.Bool())
	gens["EnableTcpReset"] = gen.PtrOf(gen.Bool())
	gens["FrontendPort"] = gen.PtrOf(gen.Int())
	gens["IdleTimeoutInMinutes"] = gen.PtrOf(gen.Int())
	gens["Protocol"] = gen.PtrOf(gen.OneConstOf(TransportProtocol_ARM_All, TransportProtocol_ARM_Tcp, TransportProtocol_ARM_Udp))
}

// AddRelatedPropertyGeneratorsForInboundNatRulePropertiesFormat_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForInboundNatRulePropertiesFormat_ARM(gens map[string]gopter.Gen) {
	gens["FrontendIPConfiguration"] = gen.PtrOf(SubResource_ARMGenerator())
}

func Test_LoadBalancersInboundNatRule_Spec_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of LoadBalancersInboundNatRule_Spec_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForLoadBalancersInboundNatRule_Spec_ARM, LoadBalancersInboundNatRule_Spec_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForLoadBalancersInboundNatRule_Spec_ARM runs a test to see if a specific instance of LoadBalancersInboundNatRule_Spec_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForLoadBalancersInboundNatRule_Spec_ARM(subject LoadBalancersInboundNatRule_Spec_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual LoadBalancersInboundNatRule_Spec_ARM
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

// Generator of LoadBalancersInboundNatRule_Spec_ARM instances for property testing - lazily instantiated by
// LoadBalancersInboundNatRule_Spec_ARMGenerator()
var loadBalancersInboundNatRule_Spec_ARMGenerator gopter.Gen

// LoadBalancersInboundNatRule_Spec_ARMGenerator returns a generator of LoadBalancersInboundNatRule_Spec_ARM instances for property testing.
// We first initialize loadBalancersInboundNatRule_Spec_ARMGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func LoadBalancersInboundNatRule_Spec_ARMGenerator() gopter.Gen {
	if loadBalancersInboundNatRule_Spec_ARMGenerator != nil {
		return loadBalancersInboundNatRule_Spec_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancersInboundNatRule_Spec_ARM(generators)
	loadBalancersInboundNatRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(LoadBalancersInboundNatRule_Spec_ARM{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForLoadBalancersInboundNatRule_Spec_ARM(generators)
	AddRelatedPropertyGeneratorsForLoadBalancersInboundNatRule_Spec_ARM(generators)
	loadBalancersInboundNatRule_Spec_ARMGenerator = gen.Struct(reflect.TypeOf(LoadBalancersInboundNatRule_Spec_ARM{}), generators)

	return loadBalancersInboundNatRule_Spec_ARMGenerator
}

// AddIndependentPropertyGeneratorsForLoadBalancersInboundNatRule_Spec_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForLoadBalancersInboundNatRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForLoadBalancersInboundNatRule_Spec_ARM is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForLoadBalancersInboundNatRule_Spec_ARM(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(InboundNatRulePropertiesFormat_ARMGenerator())
}

func Test_SubResource_ARM_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of SubResource_ARM via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForSubResource_ARM, SubResource_ARMGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForSubResource_ARM runs a test to see if a specific instance of SubResource_ARM round trips to JSON and back losslessly
func RunJSONSerializationTestForSubResource_ARM(subject SubResource_ARM) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual SubResource_ARM
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

// Generator of SubResource_ARM instances for property testing - lazily instantiated by SubResource_ARMGenerator()
var subResource_ARMGenerator gopter.Gen

// SubResource_ARMGenerator returns a generator of SubResource_ARM instances for property testing.
func SubResource_ARMGenerator() gopter.Gen {
	if subResource_ARMGenerator != nil {
		return subResource_ARMGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForSubResource_ARM(generators)
	subResource_ARMGenerator = gen.Struct(reflect.TypeOf(SubResource_ARM{}), generators)

	return subResource_ARMGenerator
}

// AddIndependentPropertyGeneratorsForSubResource_ARM is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForSubResource_ARM(gens map[string]gopter.Gen) {
	gens["Id"] = gen.PtrOf(gen.AlphaString())
}
