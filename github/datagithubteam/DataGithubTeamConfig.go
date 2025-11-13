// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagithubteam

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubTeamConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/team#slug DataGithubTeam#slug}.
	Slug *string `field:"required" json:"slug" yaml:"slug"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/team#id DataGithubTeam#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/team#membership_type DataGithubTeam#membership_type}.
	MembershipType *string `field:"optional" json:"membershipType" yaml:"membershipType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/team#results_per_page DataGithubTeam#results_per_page}.
	ResultsPerPage *float64 `field:"optional" json:"resultsPerPage" yaml:"resultsPerPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/data-sources/team#summary_only DataGithubTeam#summary_only}.
	SummaryOnly interface{} `field:"optional" json:"summaryOnly" yaml:"summaryOnly"`
}

