package repositoryautolinkreference

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryAutolinkReferenceConfig struct {
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
	// This prefix appended by a number will generate a link any time it is found in an issue, pull request, or commit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/repository_autolink_reference#key_prefix RepositoryAutolinkReference#key_prefix}
	KeyPrefix *string `field:"required" json:"keyPrefix" yaml:"keyPrefix"`
	// The repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/repository_autolink_reference#repository RepositoryAutolinkReference#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// The template of the target URL used for the links;
	//
	// must be a valid URL and contain `<num>` for the reference number
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/repository_autolink_reference#target_url_template RepositoryAutolinkReference#target_url_template}
	TargetUrlTemplate *string `field:"required" json:"targetUrlTemplate" yaml:"targetUrlTemplate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/repository_autolink_reference#id RepositoryAutolinkReference#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether this autolink reference matches alphanumeric characters. If false, this autolink reference only matches numeric characters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/repository_autolink_reference#is_alphanumeric RepositoryAutolinkReference#is_alphanumeric}
	IsAlphanumeric interface{} `field:"optional" json:"isAlphanumeric" yaml:"isAlphanumeric"`
}

