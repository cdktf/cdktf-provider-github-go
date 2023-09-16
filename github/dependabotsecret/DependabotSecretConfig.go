// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dependabotsecret

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DependabotSecretConfig struct {
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
	// Name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/dependabot_secret#repository DependabotSecret#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Name of the secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/dependabot_secret#secret_name DependabotSecret#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Encrypted value of the secret using the GitHub public key in Base64 format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/dependabot_secret#encrypted_value DependabotSecret#encrypted_value}
	EncryptedValue *string `field:"optional" json:"encryptedValue" yaml:"encryptedValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/dependabot_secret#id DependabotSecret#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Plaintext value of the secret to be encrypted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/dependabot_secret#plaintext_value DependabotSecret#plaintext_value}
	PlaintextValue *string `field:"optional" json:"plaintextValue" yaml:"plaintextValue"`
}

