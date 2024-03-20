// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesPullRequest struct {
	// New, reviewable commits pushed will dismiss previous pull request review approvals. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.0/docs/resources/organization_ruleset#dismiss_stale_reviews_on_push OrganizationRuleset#dismiss_stale_reviews_on_push}
	DismissStaleReviewsOnPush interface{} `field:"optional" json:"dismissStaleReviewsOnPush" yaml:"dismissStaleReviewsOnPush"`
	// Require an approving review in pull requests that modify files that have a designated code owner. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.0/docs/resources/organization_ruleset#require_code_owner_review OrganizationRuleset#require_code_owner_review}
	RequireCodeOwnerReview interface{} `field:"optional" json:"requireCodeOwnerReview" yaml:"requireCodeOwnerReview"`
	// The number of approving reviews that are required before a pull request can be merged. Defaults to `0`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.0/docs/resources/organization_ruleset#required_approving_review_count OrganizationRuleset#required_approving_review_count}
	RequiredApprovingReviewCount *float64 `field:"optional" json:"requiredApprovingReviewCount" yaml:"requiredApprovingReviewCount"`
	// All conversations on code must be resolved before a pull request can be merged. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.0/docs/resources/organization_ruleset#required_review_thread_resolution OrganizationRuleset#required_review_thread_resolution}
	RequiredReviewThreadResolution interface{} `field:"optional" json:"requiredReviewThreadResolution" yaml:"requiredReviewThreadResolution"`
	// Whether the most recent reviewable push must be approved by someone other than the person who pushed it.
	//
	// Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.0/docs/resources/organization_ruleset#require_last_push_approval OrganizationRuleset#require_last_push_approval}
	RequireLastPushApproval interface{} `field:"optional" json:"requireLastPushApproval" yaml:"requireLastPushApproval"`
}

