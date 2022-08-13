// Prebuilt github Provider for Terraform CDK (cdktf)
package github


type BranchProtectionV3RequiredPullRequestReviews struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#dismissal_teams BranchProtectionV3#dismissal_teams}.
	DismissalTeams *[]*string `field:"optional" json:"dismissalTeams" yaml:"dismissalTeams"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#dismissal_users BranchProtectionV3#dismissal_users}.
	DismissalUsers *[]*string `field:"optional" json:"dismissalUsers" yaml:"dismissalUsers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#dismiss_stale_reviews BranchProtectionV3#dismiss_stale_reviews}.
	DismissStaleReviews interface{} `field:"optional" json:"dismissStaleReviews" yaml:"dismissStaleReviews"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#include_admins BranchProtectionV3#include_admins}.
	IncludeAdmins interface{} `field:"optional" json:"includeAdmins" yaml:"includeAdmins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#require_code_owner_reviews BranchProtectionV3#require_code_owner_reviews}.
	RequireCodeOwnerReviews interface{} `field:"optional" json:"requireCodeOwnerReviews" yaml:"requireCodeOwnerReviews"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#required_approving_review_count BranchProtectionV3#required_approving_review_count}.
	RequiredApprovingReviewCount *float64 `field:"optional" json:"requiredApprovingReviewCount" yaml:"requiredApprovingReviewCount"`
}

