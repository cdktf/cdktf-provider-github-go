// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterpriseactionspermissions


type EnterpriseActionsPermissionsAllowedActionsConfig struct {
	// Whether GitHub-owned actions are allowed in the enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.4.0/docs/resources/enterprise_actions_permissions#github_owned_allowed EnterpriseActionsPermissions#github_owned_allowed}
	GithubOwnedAllowed interface{} `field:"required" json:"githubOwnedAllowed" yaml:"githubOwnedAllowed"`
	// Specifies a list of string-matching patterns to allow specific action(s).
	//
	// Wildcards, tags, and SHAs are allowed. For example, 'monalisa/octocat@', 'monalisa/octocat@v2', 'monalisa/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.4.0/docs/resources/enterprise_actions_permissions#patterns_allowed EnterpriseActionsPermissions#patterns_allowed}
	PatternsAllowed *[]*string `field:"optional" json:"patternsAllowed" yaml:"patternsAllowed"`
	// Whether actions in GitHub Marketplace from verified creators are allowed.
	//
	// Set to 'true' to allow all GitHub Marketplace actions by verified creators.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.4.0/docs/resources/enterprise_actions_permissions#verified_allowed EnterpriseActionsPermissions#verified_allowed}
	VerifiedAllowed interface{} `field:"optional" json:"verifiedAllowed" yaml:"verifiedAllowed"`
}

