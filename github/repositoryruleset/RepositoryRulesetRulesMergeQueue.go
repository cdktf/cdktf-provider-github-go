// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesMergeQueue struct {
	// Maximum time for a required status check to report a conclusion.
	//
	// After this much time has elapsed, checks that have not reported a conclusion will be assumed to have failed. Defaults to `60`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#check_response_timeout_minutes RepositoryRuleset#check_response_timeout_minutes}
	CheckResponseTimeoutMinutes *float64 `field:"optional" json:"checkResponseTimeoutMinutes" yaml:"checkResponseTimeoutMinutes"`
	// When set to ALLGREEN, the merge commit created by merge queue for each PR in the group must pass all required checks to merge.
	//
	// When set to HEADGREEN, only the commit at the head of the merge group, i.e. the commit containing changes from all of the PRs in the group, must pass its required checks to merge. Can be one of: ALLGREEN, HEADGREEN. Defaults to `ALLGREEN`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#grouping_strategy RepositoryRuleset#grouping_strategy}
	GroupingStrategy *string `field:"optional" json:"groupingStrategy" yaml:"groupingStrategy"`
	// Limit the number of queued pull requests requesting checks and workflow runs at the same time. Defaults to `5`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#max_entries_to_build RepositoryRuleset#max_entries_to_build}
	MaxEntriesToBuild *float64 `field:"optional" json:"maxEntriesToBuild" yaml:"maxEntriesToBuild"`
	// The maximum number of PRs that will be merged together in a group. Defaults to `5`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#max_entries_to_merge RepositoryRuleset#max_entries_to_merge}
	MaxEntriesToMerge *float64 `field:"optional" json:"maxEntriesToMerge" yaml:"maxEntriesToMerge"`
	// Method to use when merging changes from queued pull requests.
	//
	// Can be one of: MERGE, SQUASH, REBASE. Defaults to `MERGE`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#merge_method RepositoryRuleset#merge_method}
	MergeMethod *string `field:"optional" json:"mergeMethod" yaml:"mergeMethod"`
	// The minimum number of PRs that will be merged together in a group. Defaults to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#min_entries_to_merge RepositoryRuleset#min_entries_to_merge}
	MinEntriesToMerge *float64 `field:"optional" json:"minEntriesToMerge" yaml:"minEntriesToMerge"`
	// The time merge queue should wait after the first PR is added to the queue for the minimum group size to be met.
	//
	// After this time has elapsed, the minimum group size will be ignored and a smaller group will be merged. Defaults to `5`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#min_entries_to_merge_wait_minutes RepositoryRuleset#min_entries_to_merge_wait_minutes}
	MinEntriesToMergeWaitMinutes *float64 `field:"optional" json:"minEntriesToMergeWaitMinutes" yaml:"minEntriesToMergeWaitMinutes"`
}

