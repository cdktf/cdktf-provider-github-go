// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagithubrepositorypullrequests

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubRepositoryPullRequestsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/repository_pull_requests#base_repository DataGithubRepositoryPullRequests#base_repository}.
	BaseRepository *string `field:"required" json:"baseRepository" yaml:"baseRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/repository_pull_requests#base_ref DataGithubRepositoryPullRequests#base_ref}.
	BaseRef *string `field:"optional" json:"baseRef" yaml:"baseRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/repository_pull_requests#head_ref DataGithubRepositoryPullRequests#head_ref}.
	HeadRef *string `field:"optional" json:"headRef" yaml:"headRef"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/repository_pull_requests#id DataGithubRepositoryPullRequests#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/repository_pull_requests#owner DataGithubRepositoryPullRequests#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/repository_pull_requests#sort_by DataGithubRepositoryPullRequests#sort_by}.
	SortBy *string `field:"optional" json:"sortBy" yaml:"sortBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/repository_pull_requests#sort_direction DataGithubRepositoryPullRequests#sort_direction}.
	SortDirection *string `field:"optional" json:"sortDirection" yaml:"sortDirection"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/repository_pull_requests#state DataGithubRepositoryPullRequests#state}.
	State *string `field:"optional" json:"state" yaml:"state"`
}

