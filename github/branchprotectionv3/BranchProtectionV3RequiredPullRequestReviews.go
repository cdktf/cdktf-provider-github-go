package branchprotectionv3


type BranchProtectionV3RequiredPullRequestReviews struct {
	// bypass_pull_request_allowances block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/branch_protection_v3#bypass_pull_request_allowances BranchProtectionV3#bypass_pull_request_allowances}
	BypassPullRequestAllowances *BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowances `field:"optional" json:"bypassPullRequestAllowances" yaml:"bypassPullRequestAllowances"`
	// The list of team slugs with dismissal access.
	//
	// Always use slug of the team, not its name. Each team already has to have access to the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/branch_protection_v3#dismissal_teams BranchProtectionV3#dismissal_teams}
	DismissalTeams *[]*string `field:"optional" json:"dismissalTeams" yaml:"dismissalTeams"`
	// The list of user logins with dismissal access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/branch_protection_v3#dismissal_users BranchProtectionV3#dismissal_users}
	DismissalUsers *[]*string `field:"optional" json:"dismissalUsers" yaml:"dismissalUsers"`
	// Dismiss approved reviews automatically when a new commit is pushed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/branch_protection_v3#dismiss_stale_reviews BranchProtectionV3#dismiss_stale_reviews}
	DismissStaleReviews interface{} `field:"optional" json:"dismissStaleReviews" yaml:"dismissStaleReviews"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/branch_protection_v3#include_admins BranchProtectionV3#include_admins}.
	IncludeAdmins interface{} `field:"optional" json:"includeAdmins" yaml:"includeAdmins"`
	// Require an approved review in pull requests including files with a designated code owner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/branch_protection_v3#require_code_owner_reviews BranchProtectionV3#require_code_owner_reviews}
	RequireCodeOwnerReviews interface{} `field:"optional" json:"requireCodeOwnerReviews" yaml:"requireCodeOwnerReviews"`
	// Require 'x' number of approvals to satisfy branch protection requirements.
	//
	// If this is specified it must be a number between 0-6.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/branch_protection_v3#required_approving_review_count BranchProtectionV3#required_approving_review_count}
	RequiredApprovingReviewCount *float64 `field:"optional" json:"requiredApprovingReviewCount" yaml:"requiredApprovingReviewCount"`
}

