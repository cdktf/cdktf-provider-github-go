// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotection

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchProtectionConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Identifies the protection rule pattern.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#pattern BranchProtection#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// The name or node ID of the repository associated with this branch protection rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#repository_id BranchProtection#repository_id}
	RepositoryId *string `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// Setting this to 'true' to allow the branch to be deleted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#allows_deletions BranchProtection#allows_deletions}
	AllowsDeletions interface{} `field:"optional" json:"allowsDeletions" yaml:"allowsDeletions"`
	// Setting this to 'true' to allow force pushes on the branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#allows_force_pushes BranchProtection#allows_force_pushes}
	AllowsForcePushes interface{} `field:"optional" json:"allowsForcePushes" yaml:"allowsForcePushes"`
	// Setting this to 'true' enforces status checks for repository administrators.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#enforce_admins BranchProtection#enforce_admins}
	EnforceAdmins interface{} `field:"optional" json:"enforceAdmins" yaml:"enforceAdmins"`
	// The list of actor Names/IDs that are allowed to bypass force push restrictions.
	//
	// Actor names must either begin with a '/' for users or the organization name followed by a '/' for teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#force_push_bypassers BranchProtection#force_push_bypassers}
	ForcePushBypassers *[]*string `field:"optional" json:"forcePushBypassers" yaml:"forcePushBypassers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#id BranchProtection#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Setting this to 'true' will make the branch read-only and preventing any pushes to it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#lock_branch BranchProtection#lock_branch}
	LockBranch interface{} `field:"optional" json:"lockBranch" yaml:"lockBranch"`
	// Setting this to 'true' requires all conversations on code must be resolved before a pull request can be merged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#require_conversation_resolution BranchProtection#require_conversation_resolution}
	RequireConversationResolution interface{} `field:"optional" json:"requireConversationResolution" yaml:"requireConversationResolution"`
	// Setting this to 'true' enforces a linear commit Git history, which prevents anyone from pushing merge commits to a branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#required_linear_history BranchProtection#required_linear_history}
	RequiredLinearHistory interface{} `field:"optional" json:"requiredLinearHistory" yaml:"requiredLinearHistory"`
	// required_pull_request_reviews block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#required_pull_request_reviews BranchProtection#required_pull_request_reviews}
	RequiredPullRequestReviews interface{} `field:"optional" json:"requiredPullRequestReviews" yaml:"requiredPullRequestReviews"`
	// required_status_checks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#required_status_checks BranchProtection#required_status_checks}
	RequiredStatusChecks interface{} `field:"optional" json:"requiredStatusChecks" yaml:"requiredStatusChecks"`
	// Setting this to 'true' requires all commits to be signed with GPG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#require_signed_commits BranchProtection#require_signed_commits}
	RequireSignedCommits interface{} `field:"optional" json:"requireSignedCommits" yaml:"requireSignedCommits"`
	// restrict_pushes block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection#restrict_pushes BranchProtection#restrict_pushes}
	RestrictPushes interface{} `field:"optional" json:"restrictPushes" yaml:"restrictPushes"`
}

