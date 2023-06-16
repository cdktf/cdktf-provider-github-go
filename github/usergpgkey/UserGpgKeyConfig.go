package usergpgkey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type UserGpgKeyConfig struct {
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
	// Your public GPG key, generated in ASCII-armored format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/resources/user_gpg_key#armored_public_key UserGpgKey#armored_public_key}
	ArmoredPublicKey *string `field:"required" json:"armoredPublicKey" yaml:"armoredPublicKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/resources/user_gpg_key#id UserGpgKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

