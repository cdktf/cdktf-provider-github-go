// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositorypullrequest

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryPullRequestConfig struct {
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
	// Name of the branch serving as the base of the Pull Request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_pull_request#base_ref RepositoryPullRequest#base_ref}
	BaseRef *string `field:"required" json:"baseRef" yaml:"baseRef"`
	// Name of the base repository to retrieve the Pull Requests from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_pull_request#base_repository RepositoryPullRequest#base_repository}
	BaseRepository *string `field:"required" json:"baseRepository" yaml:"baseRepository"`
	// Name of the branch serving as the head of the Pull Request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_pull_request#head_ref RepositoryPullRequest#head_ref}
	HeadRef *string `field:"required" json:"headRef" yaml:"headRef"`
	// The title of the Pull Request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_pull_request#title RepositoryPullRequest#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// Body of the Pull Request.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_pull_request#body RepositoryPullRequest#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_pull_request#id RepositoryPullRequest#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Controls whether the base repository maintainers can modify the Pull Request. Default: 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_pull_request#maintainer_can_modify RepositoryPullRequest#maintainer_can_modify}
	MaintainerCanModify interface{} `field:"optional" json:"maintainerCanModify" yaml:"maintainerCanModify"`
	// Owner of the repository. If not provided, the provider's default owner is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_pull_request#owner RepositoryPullRequest#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
}

