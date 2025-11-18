// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositorycollaborators


type RepositoryCollaboratorsIgnoreTeam struct {
	// ID or slug of the team to ignore.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.3/docs/resources/repository_collaborators#team_id RepositoryCollaborators#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
}

