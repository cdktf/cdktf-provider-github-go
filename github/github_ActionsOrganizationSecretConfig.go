// Prebuilt github Provider for Terraform CDK (cdktf)
package github

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsOrganizationSecretConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_secret#secret_name ActionsOrganizationSecret#secret_name}.
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_secret#visibility ActionsOrganizationSecret#visibility}.
	Visibility *string `field:"required" json:"visibility" yaml:"visibility"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_secret#encrypted_value ActionsOrganizationSecret#encrypted_value}.
	EncryptedValue *string `field:"optional" json:"encryptedValue" yaml:"encryptedValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_secret#id ActionsOrganizationSecret#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_secret#plaintext_value ActionsOrganizationSecret#plaintext_value}.
	PlaintextValue *string `field:"optional" json:"plaintextValue" yaml:"plaintextValue"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/actions_organization_secret#selected_repository_ids ActionsOrganizationSecret#selected_repository_ids}.
	SelectedRepositoryIds *[]*float64 `field:"optional" json:"selectedRepositoryIds" yaml:"selectedRepositoryIds"`
}

