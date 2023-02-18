package actionsrunnergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsRunnerGroupConfig struct {
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
	// Name of the runner group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_runner_group#name ActionsRunnerGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The visibility of the runner group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_runner_group#visibility ActionsRunnerGroup#visibility}
	Visibility *string `field:"required" json:"visibility" yaml:"visibility"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_runner_group#id ActionsRunnerGroup#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of repository IDs that can access the runner group.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_runner_group#selected_repository_ids ActionsRunnerGroup#selected_repository_ids}
	SelectedRepositoryIds *[]*float64 `field:"optional" json:"selectedRepositoryIds" yaml:"selectedRepositoryIds"`
}

