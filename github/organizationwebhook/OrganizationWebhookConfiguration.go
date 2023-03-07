package organizationwebhook


type OrganizationWebhookConfiguration struct {
	// The URL of the webhook.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_webhook#url OrganizationWebhook#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The content type for the payload. Valid values are either 'form' or 'json'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_webhook#content_type OrganizationWebhook#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Insecure SSL boolean toggle. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_webhook#insecure_ssl OrganizationWebhook#insecure_ssl}
	InsecureSsl interface{} `field:"optional" json:"insecureSsl" yaml:"insecureSsl"`
	// The shared secret for the webhook.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_webhook#secret OrganizationWebhook#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}
