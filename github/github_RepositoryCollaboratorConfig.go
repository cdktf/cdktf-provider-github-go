// Prebuilt github Provider for Terraform CDK (cdktf)
package github

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryCollaboratorConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_collaborator#repository RepositoryCollaborator#repository}.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_collaborator#username RepositoryCollaborator#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_collaborator#id RepositoryCollaborator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_collaborator#permission RepositoryCollaborator#permission}.
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_collaborator#permission_diff_suppression RepositoryCollaborator#permission_diff_suppression}.
	PermissionDiffSuppression interface{} `field:"optional" json:"permissionDiffSuppression" yaml:"permissionDiffSuppression"`
}

