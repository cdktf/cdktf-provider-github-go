// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesRequiredStatusChecks struct {
	// required_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#required_check OrganizationRuleset#required_check}
	RequiredCheck interface{} `field:"required" json:"requiredCheck" yaml:"requiredCheck"`
	// Whether pull requests targeting a matching branch must be tested with the latest code.
	//
	// This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#strict_required_status_checks_policy OrganizationRuleset#strict_required_status_checks_policy}
	StrictRequiredStatusChecksPolicy interface{} `field:"optional" json:"strictRequiredStatusChecksPolicy" yaml:"strictRequiredStatusChecksPolicy"`
}

