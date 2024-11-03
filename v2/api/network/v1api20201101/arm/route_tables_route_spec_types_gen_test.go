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

func Test_RoutePropertiesFormat_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 100
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RoutePropertiesFormat via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRoutePropertiesFormat, RoutePropertiesFormatGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRoutePropertiesFormat runs a test to see if a specific instance of RoutePropertiesFormat round trips to JSON and back losslessly
func RunJSONSerializationTestForRoutePropertiesFormat(subject RoutePropertiesFormat) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RoutePropertiesFormat
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

// Generator of RoutePropertiesFormat instances for property testing - lazily instantiated by
// RoutePropertiesFormatGenerator()
var routePropertiesFormatGenerator gopter.Gen

// RoutePropertiesFormatGenerator returns a generator of RoutePropertiesFormat instances for property testing.
func RoutePropertiesFormatGenerator() gopter.Gen {
	if routePropertiesFormatGenerator != nil {
		return routePropertiesFormatGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRoutePropertiesFormat(generators)
	routePropertiesFormatGenerator = gen.Struct(reflect.TypeOf(RoutePropertiesFormat{}), generators)

	return routePropertiesFormatGenerator
}

// AddIndependentPropertyGeneratorsForRoutePropertiesFormat is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRoutePropertiesFormat(gens map[string]gopter.Gen) {
	gens["AddressPrefix"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopIpAddress"] = gen.PtrOf(gen.AlphaString())
	gens["NextHopType"] = gen.PtrOf(gen.OneConstOf(
		RouteNextHopType_Internet,
		RouteNextHopType_None,
		RouteNextHopType_VirtualAppliance,
		RouteNextHopType_VirtualNetworkGateway,
		RouteNextHopType_VnetLocal))
}

func Test_RouteTablesRoute_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of RouteTablesRoute_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForRouteTablesRoute_Spec, RouteTablesRoute_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForRouteTablesRoute_Spec runs a test to see if a specific instance of RouteTablesRoute_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForRouteTablesRoute_Spec(subject RouteTablesRoute_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual RouteTablesRoute_Spec
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

// Generator of RouteTablesRoute_Spec instances for property testing - lazily instantiated by
// RouteTablesRoute_SpecGenerator()
var routeTablesRoute_SpecGenerator gopter.Gen

// RouteTablesRoute_SpecGenerator returns a generator of RouteTablesRoute_Spec instances for property testing.
// We first initialize routeTablesRoute_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func RouteTablesRoute_SpecGenerator() gopter.Gen {
	if routeTablesRoute_SpecGenerator != nil {
		return routeTablesRoute_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablesRoute_Spec(generators)
	routeTablesRoute_SpecGenerator = gen.Struct(reflect.TypeOf(RouteTablesRoute_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForRouteTablesRoute_Spec(generators)
	AddRelatedPropertyGeneratorsForRouteTablesRoute_Spec(generators)
	routeTablesRoute_SpecGenerator = gen.Struct(reflect.TypeOf(RouteTablesRoute_Spec{}), generators)

	return routeTablesRoute_SpecGenerator
}

// AddIndependentPropertyGeneratorsForRouteTablesRoute_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForRouteTablesRoute_Spec(gens map[string]gopter.Gen) {
	gens["Name"] = gen.AlphaString()
}

// AddRelatedPropertyGeneratorsForRouteTablesRoute_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForRouteTablesRoute_Spec(gens map[string]gopter.Gen) {
	gens["Properties"] = gen.PtrOf(RoutePropertiesFormatGenerator())
}