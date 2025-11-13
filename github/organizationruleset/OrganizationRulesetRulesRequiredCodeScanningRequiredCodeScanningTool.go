// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningTool struct {
	// The severity level at which code scanning results that raise alerts block a reference update.
	//
	// Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/organization_ruleset#alerts_threshold OrganizationRuleset#alerts_threshold}
	AlertsThreshold *string `field:"required" json:"alertsThreshold" yaml:"alertsThreshold"`
	// The severity level at which code scanning results that raise security alerts block a reference update.
	//
	// Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/organization_ruleset#security_alerts_threshold OrganizationRuleset#security_alerts_threshold}
	SecurityAlertsThreshold *string `field:"required" json:"securityAlertsThreshold" yaml:"securityAlertsThreshold"`
	// The name of a code scanning tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/organization_ruleset#tool OrganizationRuleset#tool}
	Tool *string `field:"required" json:"tool" yaml:"tool"`
}

