package repositorymilestone

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryMilestoneConfig struct {
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
	// The owner of the GitHub Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.0/docs/resources/repository_milestone#owner RepositoryMilestone#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the GitHub Repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.0/docs/resources/repository_milestone#repository RepositoryMilestone#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The title of the milestone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.0/docs/resources/repository_milestone#title RepositoryMilestone#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// A description of the milestone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.0/docs/resources/repository_milestone#description RepositoryMilestone#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The milestone due date. In 'yyyy-mm-dd' format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.0/docs/resources/repository_milestone#due_date RepositoryMilestone#due_date}
	DueDate *string `field:"optional" json:"dueDate" yaml:"dueDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.0/docs/resources/repository_milestone#id RepositoryMilestone#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The state of the milestone. Either 'open' or 'closed'. Default: 'open'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.0/docs/resources/repository_milestone#state RepositoryMilestone#state}
	State *string `field:"optional" json:"state" yaml:"state"`
}

