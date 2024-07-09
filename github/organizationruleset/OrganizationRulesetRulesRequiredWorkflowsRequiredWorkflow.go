// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesRequiredWorkflowsRequiredWorkflow struct {
	// The path to the workflow YAML definition file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.3/docs/resources/organization_ruleset#path OrganizationRuleset#path}
	Path *string `field:"required" json:"path" yaml:"path"`
	// The repository in which the workflow is defined.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.3/docs/resources/organization_ruleset#repository_id OrganizationRuleset#repository_id}
	RepositoryId *float64 `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// The ref (branch or tag) of the workflow file to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.3/docs/resources/organization_ruleset#ref OrganizationRuleset#ref}
	Ref *string `field:"optional" json:"ref" yaml:"ref"`
}

