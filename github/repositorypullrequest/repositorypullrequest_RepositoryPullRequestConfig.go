package repositorypullrequest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryPullRequestConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_pull_request#base_ref RepositoryPullRequest#base_ref}.
	BaseRef *string `field:"required" json:"baseRef" yaml:"baseRef"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_pull_request#base_repository RepositoryPullRequest#base_repository}.
	BaseRepository *string `field:"required" json:"baseRepository" yaml:"baseRepository"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_pull_request#head_ref RepositoryPullRequest#head_ref}.
	HeadRef *string `field:"required" json:"headRef" yaml:"headRef"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_pull_request#title RepositoryPullRequest#title}.
	Title *string `field:"required" json:"title" yaml:"title"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_pull_request#body RepositoryPullRequest#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_pull_request#id RepositoryPullRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_pull_request#maintainer_can_modify RepositoryPullRequest#maintainer_can_modify}.
	MaintainerCanModify interface{} `field:"optional" json:"maintainerCanModify" yaml:"maintainerCanModify"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_pull_request#owner RepositoryPullRequest#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
}

