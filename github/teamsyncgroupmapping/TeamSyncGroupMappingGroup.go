// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsyncgroupmapping


type TeamSyncGroupMappingGroup struct {
	// The description of the IdP group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.3/docs/resources/team_sync_group_mapping#group_description TeamSyncGroupMapping#group_description}
	GroupDescription *string `field:"required" json:"groupDescription" yaml:"groupDescription"`
	// The ID of the IdP group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.3/docs/resources/team_sync_group_mapping#group_id TeamSyncGroupMapping#group_id}
	GroupId *string `field:"required" json:"groupId" yaml:"groupId"`
	// The name of the IdP group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.3/docs/resources/team_sync_group_mapping#group_name TeamSyncGroupMapping#group_name}
	GroupName *string `field:"required" json:"groupName" yaml:"groupName"`
}

