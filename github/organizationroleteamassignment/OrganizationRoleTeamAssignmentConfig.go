// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationroleteamassignment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRoleTeamAssignmentConfig struct {
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
	// The GitHub organization role id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.3/docs/resources/organization_role_team_assignment#role_id OrganizationRoleTeamAssignment#role_id}
	RoleId *string `field:"required" json:"roleId" yaml:"roleId"`
	// The GitHub team slug.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.3/docs/resources/organization_role_team_assignment#team_slug OrganizationRoleTeamAssignment#team_slug}
	TeamSlug *string `field:"required" json:"teamSlug" yaml:"teamSlug"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.3/docs/resources/organization_role_team_assignment#id OrganizationRoleTeamAssignment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

