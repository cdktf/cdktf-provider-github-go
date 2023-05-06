package team

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamConfig struct {
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
	// The name of the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/team#name Team#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Adds a default maintainer to the team. Adds the creating user to the team when 'true'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/team#create_default_maintainer Team#create_default_maintainer}
	CreateDefaultMaintainer interface{} `field:"optional" json:"createDefaultMaintainer" yaml:"createDefaultMaintainer"`
	// A description of the team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/team#description Team#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/team#id Team#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The LDAP Distinguished Name of the group where membership will be synchronized. Only available in GitHub Enterprise Server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/team#ldap_dn Team#ldap_dn}
	LdapDn *string `field:"optional" json:"ldapDn" yaml:"ldapDn"`
	// The ID or slug of the parent team, if this is a nested team.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/team#parent_team_id Team#parent_team_id}
	ParentTeamId *string `field:"optional" json:"parentTeamId" yaml:"parentTeamId"`
	// The level of privacy for the team. Must be one of 'secret' or 'closed'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/team#privacy Team#privacy}
	Privacy *string `field:"optional" json:"privacy" yaml:"privacy"`
}

