package actionsorganizationpermissions


type ActionsOrganizationPermissionsEnabledRepositoriesConfig struct {
	// List of repository IDs to enable for GitHub Actions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_permissions#repository_ids ActionsOrganizationPermissions#repository_ids}
	RepositoryIds *[]*float64 `field:"required" json:"repositoryIds" yaml:"repositoryIds"`
}
