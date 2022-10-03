package repositorywebhook


type RepositoryWebhookConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#url RepositoryWebhook#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#content_type RepositoryWebhook#content_type}.
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#insecure_ssl RepositoryWebhook#insecure_ssl}.
	InsecureSsl interface{} `field:"optional" json:"insecureSsl" yaml:"insecureSsl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_webhook#secret RepositoryWebhook#secret}.
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

