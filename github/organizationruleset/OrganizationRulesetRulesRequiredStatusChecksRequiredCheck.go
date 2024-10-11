// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesRequiredStatusChecksRequiredCheck struct {
	// The status check context name that must be present on the commit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/organization_ruleset#context OrganizationRuleset#context}
	Context *string `field:"required" json:"context" yaml:"context"`
	// The optional integration ID that this status check must originate from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/organization_ruleset#integration_id OrganizationRuleset#integration_id}
	IntegrationId *float64 `field:"optional" json:"integrationId" yaml:"integrationId"`
}

