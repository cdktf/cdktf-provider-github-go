// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionsorganizationpermissions


type ActionsOrganizationPermissionsEnabledRepositoriesConfig struct {
	// List of repository IDs to enable for GitHub Actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/actions_organization_permissions#repository_ids ActionsOrganizationPermissions#repository_ids}
	RepositoryIds *[]*float64 `field:"required" json:"repositoryIds" yaml:"repositoryIds"`
}

