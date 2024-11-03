// Code generated by azure-service-operator-codegen. DO NOT EDIT.
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT license.
package arm

import "github.com/Azure/azure-service-operator/v2/pkg/genruntime"

type RedisFirewallRule_Spec struct {
	Name string `json:"name,omitempty"`

	// Properties: redis cache firewall rule properties
	Properties *RedisFirewallRuleProperties `json:"properties,omitempty"`
}

var _ genruntime.ARMResourceSpec = &RedisFirewallRule_Spec{}

// GetAPIVersion returns the ARM API version of the resource. This is always "2023-04-01"
func (rule RedisFirewallRule_Spec) GetAPIVersion() string {
	return "2023-04-01"
}

// GetName returns the Name of the resource
func (rule *RedisFirewallRule_Spec) GetName() string {
	return rule.Name
}

// GetType returns the ARM Type of the resource. This is always "Microsoft.Cache/redis/firewallRules"
func (rule *RedisFirewallRule_Spec) GetType() string {
	return "Microsoft.Cache/redis/firewallRules"
}

// Specifies a range of IP addresses permitted to connect to the cache
type RedisFirewallRuleProperties struct {
	// EndIP: highest IP address included in the range
	EndIP *string `json:"endIP,omitempty"`

	// StartIP: lowest IP address included in the range
	StartIP *string `json:"startIP,omitempty"`
}