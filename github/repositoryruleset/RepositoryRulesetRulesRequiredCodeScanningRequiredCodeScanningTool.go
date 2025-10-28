// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesRequiredCodeScanningRequiredCodeScanningTool struct {
	// The severity level at which code scanning results that raise alerts block a reference update.
	//
	// Can be one of: `none`, `errors`, `errors_and_warnings`, `all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.1/docs/resources/repository_ruleset#alerts_threshold RepositoryRuleset#alerts_threshold}
	AlertsThreshold *string `field:"required" json:"alertsThreshold" yaml:"alertsThreshold"`
	// The severity level at which code scanning results that raise security alerts block a reference update.
	//
	// Can be one of: `none`, `critical`, `high_or_higher`, `medium_or_higher`, `all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.1/docs/resources/repository_ruleset#security_alerts_threshold RepositoryRuleset#security_alerts_threshold}
	SecurityAlertsThreshold *string `field:"required" json:"securityAlertsThreshold" yaml:"securityAlertsThreshold"`
	// The name of a code scanning tool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.1/docs/resources/repository_ruleset#tool RepositoryRuleset#tool}
	Tool *string `field:"required" json:"tool" yaml:"tool"`
}

