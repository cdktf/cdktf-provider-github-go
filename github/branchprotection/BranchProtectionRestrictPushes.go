// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotection


type BranchProtectionRestrictPushes struct {
	// Restrict pushes that create matching branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/branch_protection#blocks_creations BranchProtection#blocks_creations}
	BlocksCreations interface{} `field:"optional" json:"blocksCreations" yaml:"blocksCreations"`
	// The list of actor Names/IDs that may push to the branch.
	//
	// Actor names must either begin with a '/' for users or the organization name followed by a '/' for teams.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/branch_protection#push_allowances BranchProtection#push_allowances}
	PushAllowances *[]*string `field:"optional" json:"pushAllowances" yaml:"pushAllowances"`
}

