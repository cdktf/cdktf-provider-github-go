package datagithubrelease

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubReleaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/data-sources/release#owner DataGithubRelease#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/data-sources/release#repository DataGithubRelease#repository}.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/data-sources/release#retrieve_by DataGithubRelease#retrieve_by}.
	RetrieveBy *string `field:"required" json:"retrieveBy" yaml:"retrieveBy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/data-sources/release#id DataGithubRelease#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/data-sources/release#release_id DataGithubRelease#release_id}.
	ReleaseId *float64 `field:"optional" json:"releaseId" yaml:"releaseId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/data-sources/release#release_tag DataGithubRelease#release_tag}.
	ReleaseTag *string `field:"optional" json:"releaseTag" yaml:"releaseTag"`
}

