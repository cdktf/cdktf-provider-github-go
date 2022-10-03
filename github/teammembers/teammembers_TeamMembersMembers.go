package teammembers


type TeamMembersMembers struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team_members#username TeamMembers#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team_members#role TeamMembers#role}.
	Role *string `field:"optional" json:"role" yaml:"role"`
}

