// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package v1api20180501storage

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

func Test_DnsZonesTXTRecord_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 20
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZonesTXTRecord via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZonesTXTRecord, DnsZonesTXTRecordGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZonesTXTRecord runs a test to see if a specific instance of DnsZonesTXTRecord round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZonesTXTRecord(subject DnsZonesTXTRecord) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZonesTXTRecord
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

// Generator of DnsZonesTXTRecord instances for property testing - lazily instantiated by DnsZonesTXTRecordGenerator()
var dnsZonesTXTRecordGenerator gopter.Gen

// DnsZonesTXTRecordGenerator returns a generator of DnsZonesTXTRecord instances for property testing.
func DnsZonesTXTRecordGenerator() gopter.Gen {
	if dnsZonesTXTRecordGenerator != nil {
		return dnsZonesTXTRecordGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddRelatedPropertyGeneratorsForDnsZonesTXTRecord(generators)
	dnsZonesTXTRecordGenerator = gen.Struct(reflect.TypeOf(DnsZonesTXTRecord{}), generators)

	return dnsZonesTXTRecordGenerator
}

// AddRelatedPropertyGeneratorsForDnsZonesTXTRecord is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZonesTXTRecord(gens map[string]gopter.Gen) {
	gens["Spec"] = DnsZones_TXT_SpecGenerator()
	gens["Status"] = DnsZones_TXT_STATUSGenerator()
}

func Test_DnsZones_TXT_Spec_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_TXT_Spec via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_TXT_Spec, DnsZones_TXT_SpecGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_TXT_Spec runs a test to see if a specific instance of DnsZones_TXT_Spec round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_TXT_Spec(subject DnsZones_TXT_Spec) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_TXT_Spec
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

// Generator of DnsZones_TXT_Spec instances for property testing - lazily instantiated by DnsZones_TXT_SpecGenerator()
var dnsZones_TXT_SpecGenerator gopter.Gen

// DnsZones_TXT_SpecGenerator returns a generator of DnsZones_TXT_Spec instances for property testing.
// We first initialize dnsZones_TXT_SpecGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_TXT_SpecGenerator() gopter.Gen {
	if dnsZones_TXT_SpecGenerator != nil {
		return dnsZones_TXT_SpecGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_TXT_Spec(generators)
	dnsZones_TXT_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZones_TXT_Spec{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_TXT_Spec(generators)
	AddRelatedPropertyGeneratorsForDnsZones_TXT_Spec(generators)
	dnsZones_TXT_SpecGenerator = gen.Struct(reflect.TypeOf(DnsZones_TXT_Spec{}), generators)

	return dnsZones_TXT_SpecGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_TXT_Spec is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_TXT_Spec(gens map[string]gopter.Gen) {
	gens["AzureName"] = gen.AlphaString()
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["OriginalVersion"] = gen.AlphaString()
	gens["TTL"] = gen.PtrOf(gen.Int())
}

// AddRelatedPropertyGeneratorsForDnsZones_TXT_Spec is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_TXT_Spec(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecordGenerator())
	gens["ARecords"] = gen.SliceOf(ARecordGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecordGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecordGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecordGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecordGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecordGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecordGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecordGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecordGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResourceGenerator())
}

func Test_DnsZones_TXT_STATUS_WhenSerializedToJson_DeserializesAsEqual(t *testing.T) {
	t.Parallel()
	parameters := gopter.DefaultTestParameters()
	parameters.MinSuccessfulTests = 80
	parameters.MaxSize = 3
	properties := gopter.NewProperties(parameters)
	properties.Property(
		"Round trip of DnsZones_TXT_STATUS via JSON returns original",
		prop.ForAll(RunJSONSerializationTestForDnsZones_TXT_STATUS, DnsZones_TXT_STATUSGenerator()))
	properties.TestingRun(t, gopter.NewFormatedReporter(true, 240, os.Stdout))
}

// RunJSONSerializationTestForDnsZones_TXT_STATUS runs a test to see if a specific instance of DnsZones_TXT_STATUS round trips to JSON and back losslessly
func RunJSONSerializationTestForDnsZones_TXT_STATUS(subject DnsZones_TXT_STATUS) string {
	// Serialize to JSON
	bin, err := json.Marshal(subject)
	if err != nil {
		return err.Error()
	}

	// Deserialize back into memory
	var actual DnsZones_TXT_STATUS
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

// Generator of DnsZones_TXT_STATUS instances for property testing - lazily instantiated by
// DnsZones_TXT_STATUSGenerator()
var dnsZones_TXT_STATUSGenerator gopter.Gen

// DnsZones_TXT_STATUSGenerator returns a generator of DnsZones_TXT_STATUS instances for property testing.
// We first initialize dnsZones_TXT_STATUSGenerator with a simplified generator based on the
// fields with primitive types then replacing it with a more complex one that also handles complex fields
// to ensure any cycles in the object graph properly terminate.
func DnsZones_TXT_STATUSGenerator() gopter.Gen {
	if dnsZones_TXT_STATUSGenerator != nil {
		return dnsZones_TXT_STATUSGenerator
	}

	generators := make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_TXT_STATUS(generators)
	dnsZones_TXT_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZones_TXT_STATUS{}), generators)

	// The above call to gen.Struct() captures the map, so create a new one
	generators = make(map[string]gopter.Gen)
	AddIndependentPropertyGeneratorsForDnsZones_TXT_STATUS(generators)
	AddRelatedPropertyGeneratorsForDnsZones_TXT_STATUS(generators)
	dnsZones_TXT_STATUSGenerator = gen.Struct(reflect.TypeOf(DnsZones_TXT_STATUS{}), generators)

	return dnsZones_TXT_STATUSGenerator
}

// AddIndependentPropertyGeneratorsForDnsZones_TXT_STATUS is a factory method for creating gopter generators
func AddIndependentPropertyGeneratorsForDnsZones_TXT_STATUS(gens map[string]gopter.Gen) {
	gens["Etag"] = gen.PtrOf(gen.AlphaString())
	gens["Fqdn"] = gen.PtrOf(gen.AlphaString())
	gens["Id"] = gen.PtrOf(gen.AlphaString())
	gens["Metadata"] = gen.MapOf(gen.AlphaString(), gen.AlphaString())
	gens["Name"] = gen.PtrOf(gen.AlphaString())
	gens["ProvisioningState"] = gen.PtrOf(gen.AlphaString())
	gens["TTL"] = gen.PtrOf(gen.Int())
	gens["Type"] = gen.PtrOf(gen.AlphaString())
}

// AddRelatedPropertyGeneratorsForDnsZones_TXT_STATUS is a factory method for creating gopter generators
func AddRelatedPropertyGeneratorsForDnsZones_TXT_STATUS(gens map[string]gopter.Gen) {
	gens["AAAARecords"] = gen.SliceOf(AaaaRecord_STATUSGenerator())
	gens["ARecords"] = gen.SliceOf(ARecord_STATUSGenerator())
	gens["CNAMERecord"] = gen.PtrOf(CnameRecord_STATUSGenerator())
	gens["CaaRecords"] = gen.SliceOf(CaaRecord_STATUSGenerator())
	gens["MXRecords"] = gen.SliceOf(MxRecord_STATUSGenerator())
	gens["NSRecords"] = gen.SliceOf(NsRecord_STATUSGenerator())
	gens["PTRRecords"] = gen.SliceOf(PtrRecord_STATUSGenerator())
	gens["SOARecord"] = gen.PtrOf(SoaRecord_STATUSGenerator())
	gens["SRVRecords"] = gen.SliceOf(SrvRecord_STATUSGenerator())
	gens["TXTRecords"] = gen.SliceOf(TxtRecord_STATUSGenerator())
	gens["TargetResource"] = gen.PtrOf(SubResource_STATUSGenerator())
}