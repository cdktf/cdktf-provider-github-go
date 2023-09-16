// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagithubdependabotsecrets

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubDependabotSecretsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/data-sources/dependabot_secrets#full_name DataGithubDependabotSecrets#full_name}.
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/data-sources/dependabot_secrets#id DataGithubDependabotSecrets#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/data-sources/dependabot_secrets#name DataGithubDependabotSecrets#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

