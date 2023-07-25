package appinstallationrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppInstallationRepositoryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/app_installation_repository#installation_id AppInstallationRepository#installation_id}
	InstallationId *string `field:"required" json:"installationId" yaml:"installationId"`
	// The repository to install the app on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/app_installation_repository#repository AppInstallationRepository#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/app_installation_repository#id AppInstallationRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

