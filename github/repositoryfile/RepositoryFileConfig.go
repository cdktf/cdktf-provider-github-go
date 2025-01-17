// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryfile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryFileConfig struct {
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
	// The file's content.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#content RepositoryFile#content}
	Content *string `field:"required" json:"content" yaml:"content"`
	// The file path to manage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#file RepositoryFile#file}
	File *string `field:"required" json:"file" yaml:"file"`
	// The repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#repository RepositoryFile#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Automatically create the branch if it could not be found.
	//
	// Subsequent reads if the branch is deleted will occur from 'autocreate_branch_source_branch'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#autocreate_branch RepositoryFile#autocreate_branch}
	AutocreateBranch interface{} `field:"optional" json:"autocreateBranch" yaml:"autocreateBranch"`
	// The branch name to start from, if 'autocreate_branch' is set. Defaults to 'main'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#autocreate_branch_source_branch RepositoryFile#autocreate_branch_source_branch}
	AutocreateBranchSourceBranch *string `field:"optional" json:"autocreateBranchSourceBranch" yaml:"autocreateBranchSourceBranch"`
	// The commit hash to start from, if 'autocreate_branch' is set.
	//
	// Defaults to the tip of 'autocreate_branch_source_branch'. If provided, 'autocreate_branch_source_branch' is ignored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#autocreate_branch_source_sha RepositoryFile#autocreate_branch_source_sha}
	AutocreateBranchSourceSha *string `field:"optional" json:"autocreateBranchSourceSha" yaml:"autocreateBranchSourceSha"`
	// The branch name, defaults to the repository's default branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#branch RepositoryFile#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// The commit author name, defaults to the authenticated user's name.
	//
	// GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#commit_author RepositoryFile#commit_author}
	CommitAuthor *string `field:"optional" json:"commitAuthor" yaml:"commitAuthor"`
	// The commit author email address, defaults to the authenticated user's email address.
	//
	// GitHub app users may omit author and email information so GitHub can verify commits as the GitHub App.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#commit_email RepositoryFile#commit_email}
	CommitEmail *string `field:"optional" json:"commitEmail" yaml:"commitEmail"`
	// The commit message when creating, updating or deleting the file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#commit_message RepositoryFile#commit_message}
	CommitMessage *string `field:"optional" json:"commitMessage" yaml:"commitMessage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#id RepositoryFile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Enable overwriting existing files, defaults to "false".
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_file#overwrite_on_create RepositoryFile#overwrite_on_create}
	OverwriteOnCreate interface{} `field:"optional" json:"overwriteOnCreate" yaml:"overwriteOnCreate"`
}

