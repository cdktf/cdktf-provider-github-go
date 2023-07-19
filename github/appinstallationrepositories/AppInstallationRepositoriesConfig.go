package appinstallationrepositories

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppInstallationRepositoriesConfig struct {
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
	// The GitHub app installation id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/resources/app_installation_repositories#installation_id AppInstallationRepositories#installation_id}
	InstallationId *string `field:"required" json:"installationId" yaml:"installationId"`
	// A list of repository names to install the app on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/resources/app_installation_repositories#selected_repositories AppInstallationRepositories#selected_repositories}
	SelectedRepositories *[]*string `field:"required" json:"selectedRepositories" yaml:"selectedRepositories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/resources/app_installation_repositories#id AppInstallationRepositories#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

