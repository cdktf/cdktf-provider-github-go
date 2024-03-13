// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package codespacesorganizationsecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CodespacesOrganizationSecretConfig struct {
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
	// Name of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/codespaces_organization_secret#secret_name CodespacesOrganizationSecret#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Configures the access that repositories have to the organization secret.
	//
	// Must be one of 'all', 'private' or 'selected'. 'selected_repository_ids' is required if set to 'selected'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/codespaces_organization_secret#visibility CodespacesOrganizationSecret#visibility}
	Visibility *string `field:"required" json:"visibility" yaml:"visibility"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/codespaces_organization_secret#encrypted_value CodespacesOrganizationSecret#encrypted_value}
	EncryptedValue *string `field:"optional" json:"encryptedValue" yaml:"encryptedValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/codespaces_organization_secret#id CodespacesOrganizationSecret#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Plaintext value of the secret to be encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/codespaces_organization_secret#plaintext_value CodespacesOrganizationSecret#plaintext_value}
	PlaintextValue *string `field:"optional" json:"plaintextValue" yaml:"plaintextValue"`
	// An array of repository ids that can access the organization secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/codespaces_organization_secret#selected_repository_ids CodespacesOrganizationSecret#selected_repository_ids}
	SelectedRepositoryIds *[]*float64 `field:"optional" json:"selectedRepositoryIds" yaml:"selectedRepositoryIds"`
}

