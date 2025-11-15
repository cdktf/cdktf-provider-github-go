// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package workflowrepositorypermissions

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkflowRepositoryPermissionsConfig struct {
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
	// The GitHub repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/workflow_repository_permissions#repository WorkflowRepositoryPermissions#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Whether GitHub Actions can approve pull requests. Enabling this can be a security risk.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/workflow_repository_permissions#can_approve_pull_request_reviews WorkflowRepositoryPermissions#can_approve_pull_request_reviews}
	CanApprovePullRequestReviews interface{} `field:"optional" json:"canApprovePullRequestReviews" yaml:"canApprovePullRequestReviews"`
	// The default workflow permissions granted to the GITHUB_TOKEN when running workflows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/workflow_repository_permissions#default_workflow_permissions WorkflowRepositoryPermissions#default_workflow_permissions}
	DefaultWorkflowPermissions *string `field:"optional" json:"defaultWorkflowPermissions" yaml:"defaultWorkflowPermissions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/workflow_repository_permissions#id WorkflowRepositoryPermissions#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

