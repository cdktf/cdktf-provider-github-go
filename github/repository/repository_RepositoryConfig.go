package repository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryConfig struct {
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#name Repository#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#allow_auto_merge Repository#allow_auto_merge}.
	AllowAutoMerge interface{} `field:"optional" json:"allowAutoMerge" yaml:"allowAutoMerge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#allow_merge_commit Repository#allow_merge_commit}.
	AllowMergeCommit interface{} `field:"optional" json:"allowMergeCommit" yaml:"allowMergeCommit"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#allow_rebase_merge Repository#allow_rebase_merge}.
	AllowRebaseMerge interface{} `field:"optional" json:"allowRebaseMerge" yaml:"allowRebaseMerge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#allow_squash_merge Repository#allow_squash_merge}.
	AllowSquashMerge interface{} `field:"optional" json:"allowSquashMerge" yaml:"allowSquashMerge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#allow_update_branch Repository#allow_update_branch}.
	AllowUpdateBranch interface{} `field:"optional" json:"allowUpdateBranch" yaml:"allowUpdateBranch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#archived Repository#archived}.
	Archived interface{} `field:"optional" json:"archived" yaml:"archived"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#archive_on_destroy Repository#archive_on_destroy}.
	ArchiveOnDestroy interface{} `field:"optional" json:"archiveOnDestroy" yaml:"archiveOnDestroy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#auto_init Repository#auto_init}.
	AutoInit interface{} `field:"optional" json:"autoInit" yaml:"autoInit"`
	// Can only be set after initial repository creation, and only if the target branch exists.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#default_branch Repository#default_branch}
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#delete_branch_on_merge Repository#delete_branch_on_merge}.
	DeleteBranchOnMerge interface{} `field:"optional" json:"deleteBranchOnMerge" yaml:"deleteBranchOnMerge"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#description Repository#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#gitignore_template Repository#gitignore_template}.
	GitignoreTemplate *string `field:"optional" json:"gitignoreTemplate" yaml:"gitignoreTemplate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#has_downloads Repository#has_downloads}.
	HasDownloads interface{} `field:"optional" json:"hasDownloads" yaml:"hasDownloads"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#has_issues Repository#has_issues}.
	HasIssues interface{} `field:"optional" json:"hasIssues" yaml:"hasIssues"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#has_projects Repository#has_projects}.
	HasProjects interface{} `field:"optional" json:"hasProjects" yaml:"hasProjects"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#has_wiki Repository#has_wiki}.
	HasWiki interface{} `field:"optional" json:"hasWiki" yaml:"hasWiki"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#homepage_url Repository#homepage_url}.
	HomepageUrl *string `field:"optional" json:"homepageUrl" yaml:"homepageUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#id Repository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#ignore_vulnerability_alerts_during_read Repository#ignore_vulnerability_alerts_during_read}.
	IgnoreVulnerabilityAlertsDuringRead interface{} `field:"optional" json:"ignoreVulnerabilityAlertsDuringRead" yaml:"ignoreVulnerabilityAlertsDuringRead"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#is_template Repository#is_template}.
	IsTemplate interface{} `field:"optional" json:"isTemplate" yaml:"isTemplate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#license_template Repository#license_template}.
	LicenseTemplate *string `field:"optional" json:"licenseTemplate" yaml:"licenseTemplate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#merge_commit_message Repository#merge_commit_message}.
	MergeCommitMessage *string `field:"optional" json:"mergeCommitMessage" yaml:"mergeCommitMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#merge_commit_title Repository#merge_commit_title}.
	MergeCommitTitle *string `field:"optional" json:"mergeCommitTitle" yaml:"mergeCommitTitle"`
	// pages block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#pages Repository#pages}
	Pages *RepositoryPages `field:"optional" json:"pages" yaml:"pages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#private Repository#private}.
	Private interface{} `field:"optional" json:"private" yaml:"private"`
	// security_and_analysis block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#security_and_analysis Repository#security_and_analysis}
	SecurityAndAnalysis *RepositorySecurityAndAnalysis `field:"optional" json:"securityAndAnalysis" yaml:"securityAndAnalysis"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#squash_merge_commit_message Repository#squash_merge_commit_message}.
	SquashMergeCommitMessage *string `field:"optional" json:"squashMergeCommitMessage" yaml:"squashMergeCommitMessage"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#squash_merge_commit_title Repository#squash_merge_commit_title}.
	SquashMergeCommitTitle *string `field:"optional" json:"squashMergeCommitTitle" yaml:"squashMergeCommitTitle"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#template Repository#template}
	Template *RepositoryTemplate `field:"optional" json:"template" yaml:"template"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#topics Repository#topics}.
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#visibility Repository#visibility}.
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#vulnerability_alerts Repository#vulnerability_alerts}.
	VulnerabilityAlerts interface{} `field:"optional" json:"vulnerabilityAlerts" yaml:"vulnerabilityAlerts"`
}

