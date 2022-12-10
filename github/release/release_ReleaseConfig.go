package release

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReleaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#repository Release#repository}.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#tag_name Release#tag_name}.
	TagName *string `field:"required" json:"tagName" yaml:"tagName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#body Release#body}.
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#discussion_category_name Release#discussion_category_name}.
	DiscussionCategoryName *string `field:"optional" json:"discussionCategoryName" yaml:"discussionCategoryName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#draft Release#draft}.
	Draft interface{} `field:"optional" json:"draft" yaml:"draft"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#generate_release_notes Release#generate_release_notes}.
	GenerateReleaseNotes interface{} `field:"optional" json:"generateReleaseNotes" yaml:"generateReleaseNotes"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#id Release#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#name Release#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#prerelease Release#prerelease}.
	Prerelease interface{} `field:"optional" json:"prerelease" yaml:"prerelease"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/release#target_commitish Release#target_commitish}.
	TargetCommitish *string `field:"optional" json:"targetCommitish" yaml:"targetCommitish"`
}

