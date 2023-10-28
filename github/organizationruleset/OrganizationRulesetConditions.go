// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetConditions struct {
	// ref_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.41.0/docs/resources/organization_ruleset#ref_name OrganizationRuleset#ref_name}
	RefName *OrganizationRulesetConditionsRefName `field:"required" json:"refName" yaml:"refName"`
	// The repository IDs that the ruleset applies to. One of these IDs must match for the condition to pass.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.41.0/docs/resources/organization_ruleset#repository_id OrganizationRuleset#repository_id}
	RepositoryId *[]*float64 `field:"optional" json:"repositoryId" yaml:"repositoryId"`
	// repository_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.41.0/docs/resources/organization_ruleset#repository_name OrganizationRuleset#repository_name}
	RepositoryName *OrganizationRulesetConditionsRepositoryName `field:"optional" json:"repositoryName" yaml:"repositoryName"`
}

