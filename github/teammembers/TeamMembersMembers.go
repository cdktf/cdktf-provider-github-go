// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teammembers


type TeamMembersMembers struct {
	// The user to add to the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.1/docs/resources/team_members#username TeamMembers#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The role of the user within the team. Must be one of 'member' or 'maintainer'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.1/docs/resources/team_members#role TeamMembers#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

