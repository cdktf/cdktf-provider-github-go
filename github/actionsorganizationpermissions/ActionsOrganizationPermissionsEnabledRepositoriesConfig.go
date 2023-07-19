package actionsorganizationpermissions


type ActionsOrganizationPermissionsEnabledRepositoriesConfig struct {
	// List of repository IDs to enable for GitHub Actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/resources/actions_organization_permissions#repository_ids ActionsOrganizationPermissions#repository_ids}
	RepositoryIds *[]*float64 `field:"required" json:"repositoryIds" yaml:"repositoryIds"`
}

