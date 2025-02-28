// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionsrunnergroup

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsRunnerGroupConfig struct {
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
	// Name of the runner group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/actions_runner_group#name ActionsRunnerGroup#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The visibility of the runner group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/actions_runner_group#visibility ActionsRunnerGroup#visibility}
	Visibility *string `field:"required" json:"visibility" yaml:"visibility"`
	// Whether public repositories can be added to the runner group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/actions_runner_group#allows_public_repositories ActionsRunnerGroup#allows_public_repositories}
	AllowsPublicRepositories interface{} `field:"optional" json:"allowsPublicRepositories" yaml:"allowsPublicRepositories"`
	// If 'true', the runner group will be restricted to running only the workflows specified in the 'selected_workflows' array.
	//
	// Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/actions_runner_group#restricted_to_workflows ActionsRunnerGroup#restricted_to_workflows}
	RestrictedToWorkflows interface{} `field:"optional" json:"restrictedToWorkflows" yaml:"restrictedToWorkflows"`
	// List of repository IDs that can access the runner group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/actions_runner_group#selected_repository_ids ActionsRunnerGroup#selected_repository_ids}
	SelectedRepositoryIds *[]*float64 `field:"optional" json:"selectedRepositoryIds" yaml:"selectedRepositoryIds"`
	// List of workflows the runner group should be allowed to run.
	//
	// This setting will be ignored unless restricted_to_workflows is set to 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/actions_runner_group#selected_workflows ActionsRunnerGroup#selected_workflows}
	SelectedWorkflows *[]*string `field:"optional" json:"selectedWorkflows" yaml:"selectedWorkflows"`
}

