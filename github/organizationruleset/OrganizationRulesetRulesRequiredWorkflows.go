// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesRequiredWorkflows struct {
	// required_workflow block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/organization_ruleset#required_workflow OrganizationRuleset#required_workflow}
	RequiredWorkflow interface{} `field:"required" json:"requiredWorkflow" yaml:"requiredWorkflow"`
	// Allow repositories and branches to be created if a check would otherwise prohibit it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/organization_ruleset#do_not_enforce_on_create OrganizationRuleset#do_not_enforce_on_create}
	DoNotEnforceOnCreate interface{} `field:"optional" json:"doNotEnforceOnCreate" yaml:"doNotEnforceOnCreate"`
}

