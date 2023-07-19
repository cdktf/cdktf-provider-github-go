package organizationwebhook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationWebhookConfig struct {
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
	// A list of events which should trigger the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/resources/organization_webhook#events OrganizationWebhook#events}
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Indicate if the webhook should receive events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/resources/organization_webhook#active OrganizationWebhook#active}
	Active interface{} `field:"optional" json:"active" yaml:"active"`
	// configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/resources/organization_webhook#configuration OrganizationWebhook#configuration}
	Configuration *OrganizationWebhookConfiguration `field:"optional" json:"configuration" yaml:"configuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/resources/organization_webhook#id OrganizationWebhook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.31.0/docs/resources/organization_webhook#name OrganizationWebhook#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

