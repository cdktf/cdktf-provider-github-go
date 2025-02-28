// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetConditionsRefName struct {
	// Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/organization_ruleset#exclude OrganizationRuleset#exclude}
	Exclude *[]*string `field:"required" json:"exclude" yaml:"exclude"`
	// Array of ref names or patterns to include.
	//
	// One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/organization_ruleset#include OrganizationRuleset#include}
	Include *[]*string `field:"required" json:"include" yaml:"include"`
}

