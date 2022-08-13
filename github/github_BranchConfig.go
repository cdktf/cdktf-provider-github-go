// Prebuilt github Provider for Terraform CDK (cdktf)
package github

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch#branch Branch#branch}.
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch#repository Branch#repository}.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch#id Branch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch#source_branch Branch#source_branch}.
	SourceBranch *string `field:"optional" json:"sourceBranch" yaml:"sourceBranch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch#source_sha Branch#source_sha}.
	SourceSha *string `field:"optional" json:"sourceSha" yaml:"sourceSha"`
}

