package actionsrepositorypermissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsRepositoryPermissionsConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_repository_permissions#repository ActionsRepositoryPermissions#repository}.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_repository_permissions#allowed_actions ActionsRepositoryPermissions#allowed_actions}.
	AllowedActions *string `field:"optional" json:"allowedActions" yaml:"allowedActions"`
	// allowed_actions_config block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_repository_permissions#allowed_actions_config ActionsRepositoryPermissions#allowed_actions_config}
	AllowedActionsConfig *ActionsRepositoryPermissionsAllowedActionsConfig `field:"optional" json:"allowedActionsConfig" yaml:"allowedActionsConfig"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_repository_permissions#enabled ActionsRepositoryPermissions#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_repository_permissions#id ActionsRepositoryPermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

