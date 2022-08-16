// Prebuilt github Provider for Terraform CDK (cdktf)
package github

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsOrganizationPermissionsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_permissions#enabled_repositories ActionsOrganizationPermissions#enabled_repositories}.
	EnabledRepositories *string `field:"required" json:"enabledRepositories" yaml:"enabledRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_permissions#allowed_actions ActionsOrganizationPermissions#allowed_actions}.
	AllowedActions *string `field:"optional" json:"allowedActions" yaml:"allowedActions"`
	// allowed_actions_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_permissions#allowed_actions_config ActionsOrganizationPermissions#allowed_actions_config}
	AllowedActionsConfig *ActionsOrganizationPermissionsAllowedActionsConfig `field:"optional" json:"allowedActionsConfig" yaml:"allowedActionsConfig"`
	// enabled_repositories_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_permissions#enabled_repositories_config ActionsOrganizationPermissions#enabled_repositories_config}
	EnabledRepositoriesConfig *ActionsOrganizationPermissionsEnabledRepositoriesConfig `field:"optional" json:"enabledRepositoriesConfig" yaml:"enabledRepositoriesConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_permissions#id ActionsOrganizationPermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}
