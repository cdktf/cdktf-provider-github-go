// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotection


type BranchProtectionRequiredPullRequestReviews struct {
	// The list of actor Names/IDs with dismissal access.
	//
	// If not empty, 'restrict_dismissals' is ignored. Actor names must either begin with a '/' for users or the organization name followed by a '/' for teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/branch_protection#dismissal_restrictions BranchProtection#dismissal_restrictions}
	DismissalRestrictions *[]*string `field:"optional" json:"dismissalRestrictions" yaml:"dismissalRestrictions"`
	// Dismiss approved reviews automatically when a new commit is pushed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/branch_protection#dismiss_stale_reviews BranchProtection#dismiss_stale_reviews}
	DismissStaleReviews interface{} `field:"optional" json:"dismissStaleReviews" yaml:"dismissStaleReviews"`
	// The list of actor Names/IDs that are allowed to bypass pull request requirements.
	//
	// Actor names must either begin with a '/' for users or the organization name followed by a '/' for teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/branch_protection#pull_request_bypassers BranchProtection#pull_request_bypassers}
	PullRequestBypassers *[]*string `field:"optional" json:"pullRequestBypassers" yaml:"pullRequestBypassers"`
	// Require an approved review in pull requests including files with a designated code owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/branch_protection#require_code_owner_reviews BranchProtection#require_code_owner_reviews}
	RequireCodeOwnerReviews interface{} `field:"optional" json:"requireCodeOwnerReviews" yaml:"requireCodeOwnerReviews"`
	// Require 'x' number of approvals to satisfy branch protection requirements.
	//
	// If this is specified it must be a number between 0-6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/branch_protection#required_approving_review_count BranchProtection#required_approving_review_count}
	RequiredApprovingReviewCount *float64 `field:"optional" json:"requiredApprovingReviewCount" yaml:"requiredApprovingReviewCount"`
	// Require that The most recent push must be approved by someone other than the last pusher.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/branch_protection#require_last_push_approval BranchProtection#require_last_push_approval}
	RequireLastPushApproval interface{} `field:"optional" json:"requireLastPushApproval" yaml:"requireLastPushApproval"`
	// Restrict pull request review dismissals.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/branch_protection#restrict_dismissals BranchProtection#restrict_dismissals}
	RestrictDismissals interface{} `field:"optional" json:"restrictDismissals" yaml:"restrictDismissals"`
}

