package projectcard

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ProjectCardConfig struct {
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
	// The ID of the project column.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.29.0/docs/resources/project_card#column_id ProjectCard#column_id}
	ColumnId *string `field:"required" json:"columnId" yaml:"columnId"`
	// 'github_issue.issue_id'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.29.0/docs/resources/project_card#content_id ProjectCard#content_id}
	ContentId *float64 `field:"optional" json:"contentId" yaml:"contentId"`
	// Must be either 'Issue' or 'PullRequest'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.29.0/docs/resources/project_card#content_type ProjectCard#content_type}
	ContentType *string `field:"optional" json:"contentType" yaml:"contentType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.29.0/docs/resources/project_card#id ProjectCard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The note contents of the card. Markdown supported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.29.0/docs/resources/project_card#note ProjectCard#note}
	Note *string `field:"optional" json:"note" yaml:"note"`
}

