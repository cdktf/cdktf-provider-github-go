// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositorywebhook


type RepositoryWebhookConfiguration struct {
	// The URL of the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/repository_webhook#url RepositoryWebhook#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The content type for the payload. Valid values are either 'form' or 'json'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/repository_webhook#content_type RepositoryWebhook#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Insecure SSL boolean toggle. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/repository_webhook#insecure_ssl RepositoryWebhook#insecure_ssl}
	InsecureSsl interface{} `field:"optional" json:"insecureSsl" yaml:"insecureSsl"`
	// The shared secret for the webhook.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/repository_webhook#secret RepositoryWebhook#secret}
	Secret *string `field:"optional" json:"secret" yaml:"secret"`
}

