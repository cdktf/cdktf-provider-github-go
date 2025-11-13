// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagithubrepository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubRepositoryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/repository#description DataGithubRepository#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/repository#full_name DataGithubRepository#full_name}.
	FullName *string `field:"optional" json:"fullName" yaml:"fullName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/repository#homepage_url DataGithubRepository#homepage_url}.
	HomepageUrl *string `field:"optional" json:"homepageUrl" yaml:"homepageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/repository#id DataGithubRepository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/repository#name DataGithubRepository#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

