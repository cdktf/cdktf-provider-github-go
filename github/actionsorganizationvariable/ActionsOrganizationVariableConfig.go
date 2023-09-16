// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionsorganizationvariable

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsOrganizationVariableConfig struct {
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
	// Value of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/actions_organization_variable#value ActionsOrganizationVariable#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Name of the variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/actions_organization_variable#variable_name ActionsOrganizationVariable#variable_name}
	VariableName *string `field:"required" json:"variableName" yaml:"variableName"`
	// Configures the access that repositories have to the organization variable.
	//
	// Must be one of 'all', 'private', or 'selected'. 'selected_repository_ids' is required if set to 'selected'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/actions_organization_variable#visibility ActionsOrganizationVariable#visibility}
	Visibility *string `field:"required" json:"visibility" yaml:"visibility"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/actions_organization_variable#id ActionsOrganizationVariable#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// An array of repository ids that can access the organization variable.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.37.0/docs/resources/actions_organization_variable#selected_repository_ids ActionsOrganizationVariable#selected_repository_ids}
	SelectedRepositoryIds *[]*float64 `field:"optional" json:"selectedRepositoryIds" yaml:"selectedRepositoryIds"`
}

