// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionsorganizationpermissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsOrganizationPermissionsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The policy that controls the repositories in the organization that are allowed to run GitHub Actions.
	//
	// Can be one of: 'all', 'none', or 'selected'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/actions_organization_permissions#enabled_repositories ActionsOrganizationPermissions#enabled_repositories}
	EnabledRepositories *string `field:"required" json:"enabledRepositories" yaml:"enabledRepositories"`
	// The permissions policy that controls the actions that are allowed to run.
	//
	// Can be one of: 'all', 'local_only', or 'selected'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/actions_organization_permissions#allowed_actions ActionsOrganizationPermissions#allowed_actions}
	AllowedActions *string `field:"optional" json:"allowedActions" yaml:"allowedActions"`
	// allowed_actions_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/actions_organization_permissions#allowed_actions_config ActionsOrganizationPermissions#allowed_actions_config}
	AllowedActionsConfig *ActionsOrganizationPermissionsAllowedActionsConfig `field:"optional" json:"allowedActionsConfig" yaml:"allowedActionsConfig"`
	// enabled_repositories_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/actions_organization_permissions#enabled_repositories_config ActionsOrganizationPermissions#enabled_repositories_config}
	EnabledRepositoriesConfig *ActionsOrganizationPermissionsEnabledRepositoriesConfig `field:"optional" json:"enabledRepositoriesConfig" yaml:"enabledRepositoriesConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/actions_organization_permissions#id ActionsOrganizationPermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

