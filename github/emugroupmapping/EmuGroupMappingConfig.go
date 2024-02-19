// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package emugroupmapping

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EmuGroupMappingConfig struct {
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
	// Integer corresponding to the external group ID to be linked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.0/docs/resources/emu_group_mapping#group_id EmuGroupMapping#group_id}
	GroupId *float64 `field:"required" json:"groupId" yaml:"groupId"`
	// Slug of the GitHub team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.0/docs/resources/emu_group_mapping#team_slug EmuGroupMapping#team_slug}
	TeamSlug *string `field:"required" json:"teamSlug" yaml:"teamSlug"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.0.0/docs/resources/emu_group_mapping#id EmuGroupMapping#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

