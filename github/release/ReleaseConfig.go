// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package release

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReleaseConfig struct {
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
	// The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#repository Release#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#tag_name Release#tag_name}
	TagName *string `field:"required" json:"tagName" yaml:"tagName"`
	// Text describing the contents of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#body Release#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// If specified, a discussion of the specified category is created and linked to the release.
	//
	// The value must be a category that already exists in the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#discussion_category_name Release#discussion_category_name}
	DiscussionCategoryName *string `field:"optional" json:"discussionCategoryName" yaml:"discussionCategoryName"`
	// Set to 'false' to create a published release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#draft Release#draft}
	Draft interface{} `field:"optional" json:"draft" yaml:"draft"`
	// Set to 'true' to automatically generate the name and body for this release.
	//
	// If 'name' is specified, the specified name will be used; otherwise, a name will be automatically generated. If 'body' is specified, the body will be pre-pended to the automatically generated notes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#generate_release_notes Release#generate_release_notes}
	GenerateReleaseNotes interface{} `field:"optional" json:"generateReleaseNotes" yaml:"generateReleaseNotes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#id Release#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The name of the release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#name Release#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Set to 'false' to identify the release as a full release.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#prerelease Release#prerelease}
	Prerelease interface{} `field:"optional" json:"prerelease" yaml:"prerelease"`
	// The branch name or commit SHA the tag is created from.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/release#target_commitish Release#target_commitish}
	TargetCommitish *string `field:"optional" json:"targetCommitish" yaml:"targetCommitish"`
}

