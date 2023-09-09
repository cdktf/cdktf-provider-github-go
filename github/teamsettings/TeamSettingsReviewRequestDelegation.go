// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsettings


type TeamSettingsReviewRequestDelegation struct {
	// The algorithm to use when assigning pull requests to team members. Supported values are 'ROUND_ROBIN' and 'LOAD_BALANCE'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.35.0/docs/resources/team_settings#algorithm TeamSettings#algorithm}
	Algorithm *string `field:"optional" json:"algorithm" yaml:"algorithm"`
	// The number of team members to assign to a pull request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.35.0/docs/resources/team_settings#member_count TeamSettings#member_count}
	MemberCount *float64 `field:"optional" json:"memberCount" yaml:"memberCount"`
	// whether to notify the entire team when at least one member is also assigned to the pull request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.35.0/docs/resources/team_settings#notify TeamSettings#notify}
	Notify interface{} `field:"optional" json:"notify" yaml:"notify"`
}

