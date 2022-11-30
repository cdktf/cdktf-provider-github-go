package actionsrepositorypermissions


type ActionsRepositoryPermissionsAllowedActionsConfig struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_repository_permissions#github_owned_allowed ActionsRepositoryPermissions#github_owned_allowed}.
	GithubOwnedAllowed interface{} `field:"required" json:"githubOwnedAllowed" yaml:"githubOwnedAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_repository_permissions#patterns_allowed ActionsRepositoryPermissions#patterns_allowed}.
	PatternsAllowed *[]*string `field:"optional" json:"patternsAllowed" yaml:"patternsAllowed"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_repository_permissions#verified_allowed ActionsRepositoryPermissions#verified_allowed}.
	VerifiedAllowed interface{} `field:"optional" json:"verifiedAllowed" yaml:"verifiedAllowed"`
}

