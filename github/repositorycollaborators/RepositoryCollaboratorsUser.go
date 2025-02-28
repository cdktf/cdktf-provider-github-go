// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositorycollaborators


type RepositoryCollaboratorsUser struct {
	// (Required) The user to add to the repository as a collaborator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/repository_collaborators#username RepositoryCollaborators#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.6.0/docs/resources/repository_collaborators#permission RepositoryCollaborators#permission}.
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

