// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagithubapptoken

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubAppTokenConfig struct {
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
	// The GitHub App ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/app_token#app_id DataGithubAppToken#app_id}
	AppId *string `field:"required" json:"appId" yaml:"appId"`
	// The GitHub App installation instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/app_token#installation_id DataGithubAppToken#installation_id}
	InstallationId *string `field:"required" json:"installationId" yaml:"installationId"`
	// The GitHub App PEM file contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/app_token#pem_file DataGithubAppToken#pem_file}
	PemFile *string `field:"required" json:"pemFile" yaml:"pemFile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.36.0/docs/data-sources/app_token#id DataGithubAppToken#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

