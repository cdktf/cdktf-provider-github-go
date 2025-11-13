// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterpriseorganization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseOrganizationConfig struct {
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
	// List of organization owner usernames.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/enterprise_organization#admin_logins EnterpriseOrganization#admin_logins}
	AdminLogins *[]*string `field:"required" json:"adminLogins" yaml:"adminLogins"`
	// The billing email address.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/enterprise_organization#billing_email EnterpriseOrganization#billing_email}
	BillingEmail *string `field:"required" json:"billingEmail" yaml:"billingEmail"`
	// The ID of the enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/enterprise_organization#enterprise_id EnterpriseOrganization#enterprise_id}
	EnterpriseId *string `field:"required" json:"enterpriseId" yaml:"enterpriseId"`
	// The name of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/enterprise_organization#name EnterpriseOrganization#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The description of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/enterprise_organization#description EnterpriseOrganization#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The display name of the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/enterprise_organization#display_name EnterpriseOrganization#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

