// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package membership

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MembershipConfig struct {
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
	// The user to add to the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.41.0/docs/resources/membership#username Membership#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Instead of removing the member from the org, you can choose to downgrade their membership to 'member' when this resource is destroyed.
	//
	// This is useful when wanting to downgrade admins while keeping them in the organization
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.41.0/docs/resources/membership#downgrade_on_destroy Membership#downgrade_on_destroy}
	DowngradeOnDestroy interface{} `field:"optional" json:"downgradeOnDestroy" yaml:"downgradeOnDestroy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.41.0/docs/resources/membership#id Membership#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The role of the user within the organization. Must be one of 'member' or 'admin'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.41.0/docs/resources/membership#role Membership#role}
	Role *string `field:"optional" json:"role" yaml:"role"`
}

