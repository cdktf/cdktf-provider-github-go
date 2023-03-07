package userinvitationaccepter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserInvitationAccepterConfig struct {
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
	// Allow the ID to be unset.
	//
	// This will result in the resource being skipped when the ID is not set instead of returning an error.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/user_invitation_accepter#allow_empty_id UserInvitationAccepter#allow_empty_id}
	AllowEmptyId interface{} `field:"optional" json:"allowEmptyId" yaml:"allowEmptyId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/user_invitation_accepter#id UserInvitationAccepter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// ID of the invitation to accept. Must be set when 'allow_empty_id' is 'false'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/user_invitation_accepter#invitation_id UserInvitationAccepter#invitation_id}
	InvitationId *string `field:"optional" json:"invitationId" yaml:"invitationId"`
}
