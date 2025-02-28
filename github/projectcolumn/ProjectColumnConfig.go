// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package projectcolumn

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectColumnConfig struct {
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
	// The name of the column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/project_column#name ProjectColumn#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The ID of an existing project that the column will be created in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/project_column#project_id ProjectColumn#project_id}
	ProjectId *string `field:"required" json:"projectId" yaml:"projectId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/project_column#id ProjectColumn#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

