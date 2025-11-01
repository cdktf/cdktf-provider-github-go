// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionsorganizationsecretrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsOrganizationSecretRepositoryConfig struct {
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
	// The repository ID that can access the organization secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/actions_organization_secret_repository#repository_id ActionsOrganizationSecretRepository#repository_id}
	RepositoryId *float64 `field:"required" json:"repositoryId" yaml:"repositoryId"`
	// Name of the existing secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/actions_organization_secret_repository#secret_name ActionsOrganizationSecretRepository#secret_name}
	SecretName *string `field:"required" json:"secretName" yaml:"secretName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/actions_organization_secret_repository#id ActionsOrganizationSecretRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

