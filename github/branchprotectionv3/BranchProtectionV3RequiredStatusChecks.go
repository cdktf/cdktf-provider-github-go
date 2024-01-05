// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotectionv3


type BranchProtectionV3RequiredStatusChecks struct {
	// The list of status checks to require in order to merge into this branch.
	//
	// No status checks are required by default. Checks should be strings containing the 'context' and 'app_id' like so 'context:app_id'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/branch_protection_v3#checks BranchProtectionV3#checks}
	Checks *[]*string `field:"optional" json:"checks" yaml:"checks"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/branch_protection_v3#contexts BranchProtectionV3#contexts}.
	Contexts *[]*string `field:"optional" json:"contexts" yaml:"contexts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/branch_protection_v3#include_admins BranchProtectionV3#include_admins}.
	IncludeAdmins interface{} `field:"optional" json:"includeAdmins" yaml:"includeAdmins"`
	// Require branches to be up to date before merging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/branch_protection_v3#strict BranchProtectionV3#strict}
	Strict interface{} `field:"optional" json:"strict" yaml:"strict"`
}

