// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterpriseactionspermissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseActionsPermissionsConfig struct {
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
	// The policy that controls the organizations in the enterprise that are allowed to run GitHub Actions.
	//
	// Can be one of: 'all', 'none', or 'selected'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/enterprise_actions_permissions#enabled_organizations EnterpriseActionsPermissions#enabled_organizations}
	EnabledOrganizations *string `field:"required" json:"enabledOrganizations" yaml:"enabledOrganizations"`
	// The slug of the enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/enterprise_actions_permissions#enterprise_slug EnterpriseActionsPermissions#enterprise_slug}
	EnterpriseSlug *string `field:"required" json:"enterpriseSlug" yaml:"enterpriseSlug"`
	// The permissions policy that controls the actions that are allowed to run.
	//
	// Can be one of: 'all', 'local_only', or 'selected'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/enterprise_actions_permissions#allowed_actions EnterpriseActionsPermissions#allowed_actions}
	AllowedActions *string `field:"optional" json:"allowedActions" yaml:"allowedActions"`
	// allowed_actions_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/enterprise_actions_permissions#allowed_actions_config EnterpriseActionsPermissions#allowed_actions_config}
	AllowedActionsConfig *EnterpriseActionsPermissionsAllowedActionsConfig `field:"optional" json:"allowedActionsConfig" yaml:"allowedActionsConfig"`
	// enabled_organizations_config block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/enterprise_actions_permissions#enabled_organizations_config EnterpriseActionsPermissions#enabled_organizations_config}
	EnabledOrganizationsConfig *EnterpriseActionsPermissionsEnabledOrganizationsConfig `field:"optional" json:"enabledOrganizationsConfig" yaml:"enabledOrganizationsConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/enterprise_actions_permissions#id EnterpriseActionsPermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

