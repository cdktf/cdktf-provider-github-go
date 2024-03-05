// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRules struct {
	// branch_name_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#branch_name_pattern OrganizationRuleset#branch_name_pattern}
	BranchNamePattern *OrganizationRulesetRulesBranchNamePattern `field:"optional" json:"branchNamePattern" yaml:"branchNamePattern"`
	// commit_author_email_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#commit_author_email_pattern OrganizationRuleset#commit_author_email_pattern}
	CommitAuthorEmailPattern *OrganizationRulesetRulesCommitAuthorEmailPattern `field:"optional" json:"commitAuthorEmailPattern" yaml:"commitAuthorEmailPattern"`
	// commit_message_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#commit_message_pattern OrganizationRuleset#commit_message_pattern}
	CommitMessagePattern *OrganizationRulesetRulesCommitMessagePattern `field:"optional" json:"commitMessagePattern" yaml:"commitMessagePattern"`
	// committer_email_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#committer_email_pattern OrganizationRuleset#committer_email_pattern}
	CommitterEmailPattern *OrganizationRulesetRulesCommitterEmailPattern `field:"optional" json:"committerEmailPattern" yaml:"committerEmailPattern"`
	// Only allow users with bypass permission to create matching refs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#creation OrganizationRuleset#creation}
	Creation interface{} `field:"optional" json:"creation" yaml:"creation"`
	// Only allow users with bypass permissions to delete matching refs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#deletion OrganizationRuleset#deletion}
	Deletion interface{} `field:"optional" json:"deletion" yaml:"deletion"`
	// Prevent users with push access from force pushing to branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#non_fast_forward OrganizationRuleset#non_fast_forward}
	NonFastForward interface{} `field:"optional" json:"nonFastForward" yaml:"nonFastForward"`
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#pull_request OrganizationRuleset#pull_request}
	PullRequest *OrganizationRulesetRulesPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// Prevent merge commits from being pushed to matching branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#required_linear_history OrganizationRuleset#required_linear_history}
	RequiredLinearHistory interface{} `field:"optional" json:"requiredLinearHistory" yaml:"requiredLinearHistory"`
	// Commits pushed to matching branches must have verified signatures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#required_signatures OrganizationRuleset#required_signatures}
	RequiredSignatures interface{} `field:"optional" json:"requiredSignatures" yaml:"requiredSignatures"`
	// required_status_checks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#required_status_checks OrganizationRuleset#required_status_checks}
	RequiredStatusChecks *OrganizationRulesetRulesRequiredStatusChecks `field:"optional" json:"requiredStatusChecks" yaml:"requiredStatusChecks"`
	// required_workflows block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#required_workflows OrganizationRuleset#required_workflows}
	RequiredWorkflows *OrganizationRulesetRulesRequiredWorkflows `field:"optional" json:"requiredWorkflows" yaml:"requiredWorkflows"`
	// tag_name_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#tag_name_pattern OrganizationRuleset#tag_name_pattern}
	TagNamePattern *OrganizationRulesetRulesTagNamePattern `field:"optional" json:"tagNamePattern" yaml:"tagNamePattern"`
	// Only allow users with bypass permission to update matching refs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.1/docs/resources/organization_ruleset#update OrganizationRuleset#update}
	Update interface{} `field:"optional" json:"update" yaml:"update"`
}

