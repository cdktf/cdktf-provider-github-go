package teammembers


type TeamMembersMembers struct {
	// The user to add to the team.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team_members#username TeamMembers#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// The role of the user within the team. Must be one of 'member' or 'maintainer'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team_members#role TeamMembers#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}
