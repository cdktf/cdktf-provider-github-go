package datagithubrepositoryfile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubRepositoryFileConfig struct {
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
	// The file path to manage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/data-sources/repository_file#file DataGithubRepositoryFile#file}
	File *string `field:"required" json:"file" yaml:"file"`
	// The repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/data-sources/repository_file#repository DataGithubRepositoryFile#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The branch name, defaults to the repository's default branch.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/data-sources/repository_file#branch DataGithubRepositoryFile#branch}
	Branch *string `field:"optional" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/data-sources/repository_file#id DataGithubRepositoryFile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

