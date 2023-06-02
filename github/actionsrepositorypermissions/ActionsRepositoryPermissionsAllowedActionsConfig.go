package actionsrepositorypermissions


type ActionsRepositoryPermissionsAllowedActionsConfig struct {
	// Whether GitHub-owned actions are allowed in the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/actions_repository_permissions#github_owned_allowed ActionsRepositoryPermissions#github_owned_allowed}
	GithubOwnedAllowed interface{} `field:"required" json:"githubOwnedAllowed" yaml:"githubOwnedAllowed"`
	// Specifies a list of string-matching patterns to allow specific action(s).
	//
	// Wildcards, tags, and SHAs are allowed. For example, 'monalisa/octocat@', 'monalisa/octocat@v2', 'monalisa/'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/actions_repository_permissions#patterns_allowed ActionsRepositoryPermissions#patterns_allowed}
	PatternsAllowed *[]*string `field:"optional" json:"patternsAllowed" yaml:"patternsAllowed"`
	// Whether actions in GitHub Marketplace from verified creators are allowed.
	//
	// Set to 'true' to allow all GitHub Marketplace actions by verified creators.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/actions_repository_permissions#verified_allowed ActionsRepositoryPermissions#verified_allowed}
	VerifiedAllowed interface{} `field:"optional" json:"verifiedAllowed" yaml:"verifiedAllowed"`
}

