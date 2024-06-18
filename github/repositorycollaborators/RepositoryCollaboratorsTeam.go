// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositorycollaborators


type RepositoryCollaboratorsTeam struct {
	// Team ID or slug to add to the repository as a collaborator.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/repository_collaborators#team_id RepositoryCollaborators#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/repository_collaborators#permission RepositoryCollaborators#permission}.
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

