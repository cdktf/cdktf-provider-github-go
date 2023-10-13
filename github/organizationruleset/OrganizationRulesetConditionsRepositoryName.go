// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetConditionsRepositoryName struct {
	// Array of repository names or patterns to exclude. The condition will not pass if any of these patterns match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/organization_ruleset#exclude OrganizationRuleset#exclude}
	Exclude *[]*string `field:"required" json:"exclude" yaml:"exclude"`
	// Array of repository names or patterns to include.
	//
	// One of these patterns must match for the condition to pass. Also accepts `~ALL` to include all repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/organization_ruleset#include OrganizationRuleset#include}
	Include *[]*string `field:"required" json:"include" yaml:"include"`
	// Whether renaming of target repositories is prevented.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/organization_ruleset#protected OrganizationRuleset#protected}
	Protected interface{} `field:"optional" json:"protected" yaml:"protected"`
}

