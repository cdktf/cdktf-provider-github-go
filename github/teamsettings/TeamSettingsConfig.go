// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamSettingsConfig struct {
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
	// The GitHub team id or the GitHub team slug.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/team_settings#team_id TeamSettings#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/team_settings#id TeamSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// review_request_delegation block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/team_settings#review_request_delegation TeamSettings#review_request_delegation}
	ReviewRequestDelegation *TeamSettingsReviewRequestDelegation `field:"optional" json:"reviewRequestDelegation" yaml:"reviewRequestDelegation"`
}

