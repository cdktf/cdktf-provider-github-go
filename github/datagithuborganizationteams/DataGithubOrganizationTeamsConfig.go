package datagithuborganizationteams

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubOrganizationTeamsConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.1/docs/data-sources/organization_teams#id DataGithubOrganizationTeams#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.1/docs/data-sources/organization_teams#results_per_page DataGithubOrganizationTeams#results_per_page}.
	ResultsPerPage *float64 `field:"optional" json:"resultsPerPage" yaml:"resultsPerPage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.1/docs/data-sources/organization_teams#root_teams_only DataGithubOrganizationTeams#root_teams_only}.
	RootTeamsOnly interface{} `field:"optional" json:"rootTeamsOnly" yaml:"rootTeamsOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.1/docs/data-sources/organization_teams#summary_only DataGithubOrganizationTeams#summary_only}.
	SummaryOnly interface{} `field:"optional" json:"summaryOnly" yaml:"summaryOnly"`
}

