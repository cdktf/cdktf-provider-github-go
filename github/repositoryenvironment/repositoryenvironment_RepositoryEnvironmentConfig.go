package repositoryenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#environment RepositoryEnvironment#environment}.
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#repository RepositoryEnvironment#repository}.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// deployment_branch_policy block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#deployment_branch_policy RepositoryEnvironment#deployment_branch_policy}
	DeploymentBranchPolicy *RepositoryEnvironmentDeploymentBranchPolicy `field:"optional" json:"deploymentBranchPolicy" yaml:"deploymentBranchPolicy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#id RepositoryEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// reviewers block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#reviewers RepositoryEnvironment#reviewers}
	Reviewers interface{} `field:"optional" json:"reviewers" yaml:"reviewers"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#wait_timer RepositoryEnvironment#wait_timer}.
	WaitTimer *float64 `field:"optional" json:"waitTimer" yaml:"waitTimer"`
}

