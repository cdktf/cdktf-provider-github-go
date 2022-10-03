package repositorywebhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryWebhookConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#events RepositoryWebhook#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#repository RepositoryWebhook#repository}.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#active RepositoryWebhook#active}.
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#configuration RepositoryWebhook#configuration}
	Configuration *RepositoryWebhookConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#id RepositoryWebhook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#name RepositoryWebhook#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

