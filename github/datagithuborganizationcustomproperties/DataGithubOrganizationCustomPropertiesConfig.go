// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagithuborganizationcustomproperties

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubOrganizationCustomPropertiesConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/data-sources/organization_custom_properties#property_name DataGithubOrganizationCustomProperties#property_name}.
	PropertyName *string `field:"required" json:"propertyName" yaml:"propertyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/data-sources/organization_custom_properties#allowed_values DataGithubOrganizationCustomProperties#allowed_values}.
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/data-sources/organization_custom_properties#default_value DataGithubOrganizationCustomProperties#default_value}.
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/data-sources/organization_custom_properties#description DataGithubOrganizationCustomProperties#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/data-sources/organization_custom_properties#id DataGithubOrganizationCustomProperties#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/data-sources/organization_custom_properties#required DataGithubOrganizationCustomProperties#required}.
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/data-sources/organization_custom_properties#value_type DataGithubOrganizationCustomProperties#value_type}.
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

