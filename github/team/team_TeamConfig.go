package team

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team#name Team#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team#create_default_maintainer Team#create_default_maintainer}.
	CreateDefaultMaintainer interface{} `field:"optional" json:"createDefaultMaintainer" yaml:"createDefaultMaintainer"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team#description Team#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team#id Team#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team#ldap_dn Team#ldap_dn}.
	LdapDn *string `field:"optional" json:"ldapDn" yaml:"ldapDn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team#parent_team_id Team#parent_team_id}.
	ParentTeamId *float64 `field:"optional" json:"parentTeamId" yaml:"parentTeamId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/team#privacy Team#privacy}.
	Privacy *string `field:"optional" json:"privacy" yaml:"privacy"`
}

