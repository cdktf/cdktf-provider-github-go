// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationcustomproperties

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationCustomPropertiesConfig struct {
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
	// The name of the custom property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/organization_custom_properties#property_name OrganizationCustomProperties#property_name}
	PropertyName *string `field:"required" json:"propertyName" yaml:"propertyName"`
	// The allowed values of the custom property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/organization_custom_properties#allowed_values OrganizationCustomProperties#allowed_values}
	AllowedValues *[]*string `field:"optional" json:"allowedValues" yaml:"allowedValues"`
	// The default value of the custom property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/organization_custom_properties#default_value OrganizationCustomProperties#default_value}
	DefaultValue *string `field:"optional" json:"defaultValue" yaml:"defaultValue"`
	// The description of the custom property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/organization_custom_properties#description OrganizationCustomProperties#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/organization_custom_properties#id OrganizationCustomProperties#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether the custom property is required.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/organization_custom_properties#required OrganizationCustomProperties#required}
	Required interface{} `field:"optional" json:"required" yaml:"required"`
	// The type of the custom property.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/organization_custom_properties#value_type OrganizationCustomProperties#value_type}
	ValueType *string `field:"optional" json:"valueType" yaml:"valueType"`
}

