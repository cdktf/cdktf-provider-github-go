package issue

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IssueConfig struct {
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
	// The GitHub repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/issue#repository Issue#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Title of the issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/issue#title Issue#title}
	Title *string `field:"required" json:"title" yaml:"title"`
	// List of Logins to assign to the issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/issue#assignees Issue#assignees}
	Assignees *[]*string `field:"optional" json:"assignees" yaml:"assignees"`
	// Body of the issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/issue#body Issue#body}
	Body *string `field:"optional" json:"body" yaml:"body"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/issue#id Issue#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// List of labels to attach to the issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/issue#labels Issue#labels}
	Labels *[]*string `field:"optional" json:"labels" yaml:"labels"`
	// Milestone number to assign to the issue.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.32.0/docs/resources/issue#milestone_number Issue#milestone_number}
	MilestoneNumber *float64 `field:"optional" json:"milestoneNumber" yaml:"milestoneNumber"`
}

