// Prebuilt github Provider for Terraform CDK (cdktf)
package github


type BranchProtectionRequiredPullRequestReviews struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection#dismissal_restrictions BranchProtection#dismissal_restrictions}.
	DismissalRestrictions *[]*string `field:"optional" json:"dismissalRestrictions" yaml:"dismissalRestrictions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection#dismiss_stale_reviews BranchProtection#dismiss_stale_reviews}.
	DismissStaleReviews interface{} `field:"optional" json:"dismissStaleReviews" yaml:"dismissStaleReviews"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection#pull_request_bypassers BranchProtection#pull_request_bypassers}.
	PullRequestBypassers *[]*string `field:"optional" json:"pullRequestBypassers" yaml:"pullRequestBypassers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection#require_code_owner_reviews BranchProtection#require_code_owner_reviews}.
	RequireCodeOwnerReviews interface{} `field:"optional" json:"requireCodeOwnerReviews" yaml:"requireCodeOwnerReviews"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection#required_approving_review_count BranchProtection#required_approving_review_count}.
	RequiredApprovingReviewCount *float64 `field:"optional" json:"requiredApprovingReviewCount" yaml:"requiredApprovingReviewCount"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection#restrict_dismissals BranchProtection#restrict_dismissals}.
	RestrictDismissals interface{} `field:"optional" json:"restrictDismissals" yaml:"restrictDismissals"`
}

