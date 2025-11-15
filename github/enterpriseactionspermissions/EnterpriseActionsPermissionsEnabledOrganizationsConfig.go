// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterpriseactionspermissions


type EnterpriseActionsPermissionsEnabledOrganizationsConfig struct {
	// List of organization IDs to enable for GitHub Actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/enterprise_actions_permissions#organization_ids EnterpriseActionsPermissions#organization_ids}
	OrganizationIds *[]*float64 `field:"required" json:"organizationIds" yaml:"organizationIds"`
}

