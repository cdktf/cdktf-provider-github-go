// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotectionv3

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchProtectionV3Config struct {
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
	// The Git branch to protect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/branch_protection_v3#branch BranchProtectionV3#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The GitHub repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/branch_protection_v3#repository BranchProtectionV3#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Setting this to 'true' enforces status checks for repository administrators.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/branch_protection_v3#enforce_admins BranchProtectionV3#enforce_admins}
	EnforceAdmins interface{} `field:"optional" json:"enforceAdmins" yaml:"enforceAdmins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/branch_protection_v3#id BranchProtectionV3#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Setting this to 'true' requires all conversations on code must be resolved before a pull request can be merged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/branch_protection_v3#require_conversation_resolution BranchProtectionV3#require_conversation_resolution}
	RequireConversationResolution interface{} `field:"optional" json:"requireConversationResolution" yaml:"requireConversationResolution"`
	// required_pull_request_reviews block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/branch_protection_v3#required_pull_request_reviews BranchProtectionV3#required_pull_request_reviews}
	RequiredPullRequestReviews *BranchProtectionV3RequiredPullRequestReviews `field:"optional" json:"requiredPullRequestReviews" yaml:"requiredPullRequestReviews"`
	// required_status_checks block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/branch_protection_v3#required_status_checks BranchProtectionV3#required_status_checks}
	RequiredStatusChecks *BranchProtectionV3RequiredStatusChecks `field:"optional" json:"requiredStatusChecks" yaml:"requiredStatusChecks"`
	// Setting this to 'true' requires all commits to be signed with GPG.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/branch_protection_v3#require_signed_commits BranchProtectionV3#require_signed_commits}
	RequireSignedCommits interface{} `field:"optional" json:"requireSignedCommits" yaml:"requireSignedCommits"`
	// restrictions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/branch_protection_v3#restrictions BranchProtectionV3#restrictions}
	Restrictions *BranchProtectionV3Restrictions `field:"optional" json:"restrictions" yaml:"restrictions"`
}

