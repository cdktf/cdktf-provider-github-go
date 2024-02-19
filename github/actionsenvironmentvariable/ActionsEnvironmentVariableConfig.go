// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionsenvironmentvariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsEnvironmentVariableConfig struct {
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
	// Name of the environment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.0/docs/resources/actions_environment_variable#environment ActionsEnvironmentVariable#environment}
	Environment *string `field:"required" json:"environment" yaml:"environment"`
	// Name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.0/docs/resources/actions_environment_variable#repository ActionsEnvironmentVariable#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.0/docs/resources/actions_environment_variable#value ActionsEnvironmentVariable#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.0/docs/resources/actions_environment_variable#variable_name ActionsEnvironmentVariable#variable_name}
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.0/docs/resources/actions_environment_variable#id ActionsEnvironmentVariable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

