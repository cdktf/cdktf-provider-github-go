// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionshostedrunner

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsHostedRunnerConfig struct {
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
	// image block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#image ActionsHostedRunner#image}
	Image *ActionsHostedRunnerImage `field:"required" json:"image" yaml:"image"`
	// Name of the hosted runner.
	//
	// Must be between 1 and 64 characters and may only contain upper and lowercase letters a-z, numbers 0-9, '.', '-', and '_'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#name ActionsHostedRunner#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The runner group ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#runner_group_id ActionsHostedRunner#runner_group_id}
	RunnerGroupId *float64 `field:"required" json:"runnerGroupId" yaml:"runnerGroupId"`
	// Machine size (e.g., '4-core', '8-core'). Can be updated to scale the runner.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#size ActionsHostedRunner#size}
	Size *string `field:"required" json:"size" yaml:"size"`
	// Whether this runner should be used to generate custom images. Cannot be changed after creation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#image_gen ActionsHostedRunner#image_gen}
	ImageGen interface{} `field:"optional" json:"imageGen" yaml:"imageGen"`
	// The version of the runner image to deploy. This is relevant only for runners using custom images.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#image_version ActionsHostedRunner#image_version}
	ImageVersion *string `field:"optional" json:"imageVersion" yaml:"imageVersion"`
	// Maximum number of runners to scale up to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#maximum_runners ActionsHostedRunner#maximum_runners}
	MaximumRunners *float64 `field:"optional" json:"maximumRunners" yaml:"maximumRunners"`
	// Whether to enable static public IP.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#public_ip_enabled ActionsHostedRunner#public_ip_enabled}
	PublicIpEnabled interface{} `field:"optional" json:"publicIpEnabled" yaml:"publicIpEnabled"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/actions_hosted_runner#timeouts ActionsHostedRunner#timeouts}
	Timeouts *ActionsHostedRunnerTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}

