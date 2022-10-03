package actionsorganizationpermissions


type ActionsOrganizationPermissionsAllowedActionsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_permissions#github_owned_allowed ActionsOrganizationPermissions#github_owned_allowed}.
	GithubOwnedAllowed interface{} `field:"required" json:"githubOwnedAllowed" yaml:"githubOwnedAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_permissions#patterns_allowed ActionsOrganizationPermissions#patterns_allowed}.
	PatternsAllowed *[]*string `field:"optional" json:"patternsAllowed" yaml:"patternsAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_permissions#verified_allowed ActionsOrganizationPermissions#verified_allowed}.
	VerifiedAllowed interface{} `field:"optional" json:"verifiedAllowed" yaml:"verifiedAllowed"`
}

