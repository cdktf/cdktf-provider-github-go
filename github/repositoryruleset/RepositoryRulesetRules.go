// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRules struct {
	// branch_name_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#branch_name_pattern RepositoryRuleset#branch_name_pattern}
	BranchNamePattern *RepositoryRulesetRulesBranchNamePattern `field:"optional" json:"branchNamePattern" yaml:"branchNamePattern"`
	// commit_author_email_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#commit_author_email_pattern RepositoryRuleset#commit_author_email_pattern}
	CommitAuthorEmailPattern *RepositoryRulesetRulesCommitAuthorEmailPattern `field:"optional" json:"commitAuthorEmailPattern" yaml:"commitAuthorEmailPattern"`
	// commit_message_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#commit_message_pattern RepositoryRuleset#commit_message_pattern}
	CommitMessagePattern *RepositoryRulesetRulesCommitMessagePattern `field:"optional" json:"commitMessagePattern" yaml:"commitMessagePattern"`
	// committer_email_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#committer_email_pattern RepositoryRuleset#committer_email_pattern}
	CommitterEmailPattern *RepositoryRulesetRulesCommitterEmailPattern `field:"optional" json:"committerEmailPattern" yaml:"committerEmailPattern"`
	// Only allow users with bypass permission to create matching refs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#creation RepositoryRuleset#creation}
	Creation interface{} `field:"optional" json:"creation" yaml:"creation"`
	// Only allow users with bypass permissions to delete matching refs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#deletion RepositoryRuleset#deletion}
	Deletion interface{} `field:"optional" json:"deletion" yaml:"deletion"`
	// Prevent users with push access from force pushing to branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#non_fast_forward RepositoryRuleset#non_fast_forward}
	NonFastForward interface{} `field:"optional" json:"nonFastForward" yaml:"nonFastForward"`
	// pull_request block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#pull_request RepositoryRuleset#pull_request}
	PullRequest *RepositoryRulesetRulesPullRequest `field:"optional" json:"pullRequest" yaml:"pullRequest"`
	// required_deployments block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#required_deployments RepositoryRuleset#required_deployments}
	RequiredDeployments *RepositoryRulesetRulesRequiredDeployments `field:"optional" json:"requiredDeployments" yaml:"requiredDeployments"`
	// Prevent merge commits from being pushed to matching branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#required_linear_history RepositoryRuleset#required_linear_history}
	RequiredLinearHistory interface{} `field:"optional" json:"requiredLinearHistory" yaml:"requiredLinearHistory"`
	// Commits pushed to matching branches must have verified signatures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#required_signatures RepositoryRuleset#required_signatures}
	RequiredSignatures interface{} `field:"optional" json:"requiredSignatures" yaml:"requiredSignatures"`
	// required_status_checks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#required_status_checks RepositoryRuleset#required_status_checks}
	RequiredStatusChecks *RepositoryRulesetRulesRequiredStatusChecks `field:"optional" json:"requiredStatusChecks" yaml:"requiredStatusChecks"`
	// tag_name_pattern block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#tag_name_pattern RepositoryRuleset#tag_name_pattern}
	TagNamePattern *RepositoryRulesetRulesTagNamePattern `field:"optional" json:"tagNamePattern" yaml:"tagNamePattern"`
	// Only allow users with bypass permission to update matching refs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#update RepositoryRuleset#update}
	Update interface{} `field:"optional" json:"update" yaml:"update"`
	// Branch can pull changes from its upstream repository.
	//
	// This is only applicable to forked repositories. Requires `update` to be set to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/repository_ruleset#update_allows_fetch_and_merge RepositoryRuleset#update_allows_fetch_and_merge}
	UpdateAllowsFetchAndMerge interface{} `field:"optional" json:"updateAllowsFetchAndMerge" yaml:"updateAllowsFetchAndMerge"`
}

