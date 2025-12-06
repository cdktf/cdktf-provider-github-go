// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionshostedrunner


type ActionsHostedRunnerImage struct {
	// The image ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#id ActionsHostedRunner#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The image source (github, partner, or custom).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#source ActionsHostedRunner#source}
	Source *string `field:"optional" json:"source" yaml:"source"`
}

