// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionsorganizationpermissions


type ActionsOrganizationPermissionsAllowedActionsConfig struct {
	// Whether GitHub-owned actions are allowed in the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/actions_organization_permissions#github_owned_allowed ActionsOrganizationPermissions#github_owned_allowed}
	GithubOwnedAllowed interface{} `field:"required" json:"githubOwnedAllowed" yaml:"githubOwnedAllowed"`
	// Specifies a list of string-matching patterns to allow specific action(s).
	//
	// Wildcards, tags, and SHAs are allowed. For example, 'monalisa/octocat@', 'monalisa/octocat@v2', 'monalisa/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/actions_organization_permissions#patterns_allowed ActionsOrganizationPermissions#patterns_allowed}
	PatternsAllowed *[]*string `field:"optional" json:"patternsAllowed" yaml:"patternsAllowed"`
	// Whether actions in GitHub Marketplace from verified creators are allowed.
	//
	// Set to 'true' to allow all GitHub Marketplace actions by verified creators.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/actions_organization_permissions#verified_allowed ActionsOrganizationPermissions#verified_allowed}
	VerifiedAllowed interface{} `field:"optional" json:"verifiedAllowed" yaml:"verifiedAllowed"`
}

