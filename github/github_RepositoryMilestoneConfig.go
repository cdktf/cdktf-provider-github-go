// Prebuilt github Provider for Terraform CDK (cdktf)
package github

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryMilestoneConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_milestone#owner RepositoryMilestone#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_milestone#repository RepositoryMilestone#repository}.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_milestone#title RepositoryMilestone#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_milestone#description RepositoryMilestone#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// in yyyy-mm-dd format.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_milestone#due_date RepositoryMilestone#due_date}
	DueDate *string `field:"optional" json:"dueDate" yaml:"dueDate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_milestone#id RepositoryMilestone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_milestone#state RepositoryMilestone#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

