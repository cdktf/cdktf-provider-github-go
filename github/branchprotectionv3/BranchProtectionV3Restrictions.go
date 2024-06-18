// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotectionv3


type BranchProtectionV3Restrictions struct {
	// The list of app slugs with push access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection_v3#apps BranchProtectionV3#apps}
	Apps *[]*string `field:"optional" json:"apps" yaml:"apps"`
	// The list of team slugs with push access.
	//
	// Always use slug of the team, not its name. Each team already has to have access to the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection_v3#teams BranchProtectionV3#teams}
	Teams *[]*string `field:"optional" json:"teams" yaml:"teams"`
	// The list of user logins with push access.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/branch_protection_v3#users BranchProtectionV3#users}
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

