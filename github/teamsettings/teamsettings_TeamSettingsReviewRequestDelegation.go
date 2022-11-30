package teamsettings


type TeamSettingsReviewRequestDelegation struct {
	// Algorithm to use when determining team members to be assigned to a pull request.
	//
	// Allowed values are ROUND_ROBIN and LOAD_BALANCE
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team_settings#algorithm TeamSettings#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The number of reviewers to be assigned to a pull request from this team.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team_settings#member_count TeamSettings#member_count}
	MemberCount *float64 `field:"optional" json:"memberCount" yaml:"memberCount"`
	// Notify the entire team when a pull request is assigned to a member of the team.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team_settings#notify TeamSettings#notify}
	Notify interface{} `field:"optional" json:"notify" yaml:"notify"`
}

