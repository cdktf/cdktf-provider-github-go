// Prebuilt github Provider for Terraform CDK (cdktf)
package github


type OrganizationWebhookConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_webhook#url OrganizationWebhook#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_webhook#content_type OrganizationWebhook#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_webhook#insecure_ssl OrganizationWebhook#insecure_ssl}.
	InsecureSsl interface{} `field:"optional" json:"insecureSsl" yaml:"insecureSsl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_webhook#secret OrganizationWebhook#secret}.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

