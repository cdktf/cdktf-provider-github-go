// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repository

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryConfig struct {
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
	// The name of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#name Repository#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Set to 'true' to allow auto-merging pull requests on the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#allow_auto_merge Repository#allow_auto_merge}
	AllowAutoMerge interface{} `field:"optional" json:"allowAutoMerge" yaml:"allowAutoMerge"`
	// Set to 'false' to disable merge commits on the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#allow_merge_commit Repository#allow_merge_commit}
	AllowMergeCommit interface{} `field:"optional" json:"allowMergeCommit" yaml:"allowMergeCommit"`
	// Set to 'false' to disable rebase merges on the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#allow_rebase_merge Repository#allow_rebase_merge}
	AllowRebaseMerge interface{} `field:"optional" json:"allowRebaseMerge" yaml:"allowRebaseMerge"`
	// Set to 'false' to disable squash merges on the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#allow_squash_merge Repository#allow_squash_merge}
	AllowSquashMerge interface{} `field:"optional" json:"allowSquashMerge" yaml:"allowSquashMerge"`
	// Set to 'true' to always suggest updating pull request branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#allow_update_branch Repository#allow_update_branch}
	AllowUpdateBranch interface{} `field:"optional" json:"allowUpdateBranch" yaml:"allowUpdateBranch"`
	// Specifies if the repository should be archived. Defaults to 'false'. NOTE Currently, the API does not support unarchiving.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#archived Repository#archived}
	Archived interface{} `field:"optional" json:"archived" yaml:"archived"`
	// Set to 'true' to archive the repository instead of deleting on destroy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#archive_on_destroy Repository#archive_on_destroy}
	ArchiveOnDestroy interface{} `field:"optional" json:"archiveOnDestroy" yaml:"archiveOnDestroy"`
	// Set to 'true' to produce an initial commit in the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#auto_init Repository#auto_init}
	AutoInit interface{} `field:"optional" json:"autoInit" yaml:"autoInit"`
	// Can only be set after initial repository creation, and only if the target branch exists.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#default_branch Repository#default_branch}
	DefaultBranch *string `field:"optional" json:"defaultBranch" yaml:"defaultBranch"`
	// Automatically delete head branch after a pull request is merged. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#delete_branch_on_merge Repository#delete_branch_on_merge}
	DeleteBranchOnMerge interface{} `field:"optional" json:"deleteBranchOnMerge" yaml:"deleteBranchOnMerge"`
	// A description of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#description Repository#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Use the name of the template without the extension. For example, 'Haskell'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#gitignore_template Repository#gitignore_template}
	GitignoreTemplate *string `field:"optional" json:"gitignoreTemplate" yaml:"gitignoreTemplate"`
	// Set to 'true' to enable GitHub Discussions on the repository. Defaults to 'false'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#has_discussions Repository#has_discussions}
	HasDiscussions interface{} `field:"optional" json:"hasDiscussions" yaml:"hasDiscussions"`
	// Set to 'true' to enable the (deprecated) downloads features on the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#has_downloads Repository#has_downloads}
	HasDownloads interface{} `field:"optional" json:"hasDownloads" yaml:"hasDownloads"`
	// Set to 'true' to enable the GitHub Issues features on the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#has_issues Repository#has_issues}
	HasIssues interface{} `field:"optional" json:"hasIssues" yaml:"hasIssues"`
	// Set to 'true' to enable the GitHub Projects features on the repository.
	//
	// Per the GitHub documentation when in an organization that has disabled repository projects it will default to 'false' and will otherwise default to 'true'. If you specify 'true' when it has been disabled it will return an error.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#has_projects Repository#has_projects}
	HasProjects interface{} `field:"optional" json:"hasProjects" yaml:"hasProjects"`
	// Set to 'true' to enable the GitHub Wiki features on the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#has_wiki Repository#has_wiki}
	HasWiki interface{} `field:"optional" json:"hasWiki" yaml:"hasWiki"`
	// URL of a page describing the project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#homepage_url Repository#homepage_url}
	HomepageUrl *string `field:"optional" json:"homepageUrl" yaml:"homepageUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#id Repository#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Set to true to not call the vulnerability alerts endpoint so the resource can also be used without admin permissions during read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#ignore_vulnerability_alerts_during_read Repository#ignore_vulnerability_alerts_during_read}
	IgnoreVulnerabilityAlertsDuringRead interface{} `field:"optional" json:"ignoreVulnerabilityAlertsDuringRead" yaml:"ignoreVulnerabilityAlertsDuringRead"`
	// Set to 'true' to tell GitHub that this is a template repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#is_template Repository#is_template}
	IsTemplate interface{} `field:"optional" json:"isTemplate" yaml:"isTemplate"`
	// Use the name of the template without the extension. For example, 'mit' or 'mpl-2.0'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#license_template Repository#license_template}
	LicenseTemplate *string `field:"optional" json:"licenseTemplate" yaml:"licenseTemplate"`
	// Can be 'PR_BODY', 'PR_TITLE', or 'BLANK' for a default merge commit message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#merge_commit_message Repository#merge_commit_message}
	MergeCommitMessage *string `field:"optional" json:"mergeCommitMessage" yaml:"mergeCommitMessage"`
	// Can be 'PR_TITLE' or 'MERGE_MESSAGE' for a default merge commit title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#merge_commit_title Repository#merge_commit_title}
	MergeCommitTitle *string `field:"optional" json:"mergeCommitTitle" yaml:"mergeCommitTitle"`
	// pages block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#pages Repository#pages}
	Pages *RepositoryPages `field:"optional" json:"pages" yaml:"pages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#private Repository#private}.
	Private interface{} `field:"optional" json:"private" yaml:"private"`
	// security_and_analysis block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#security_and_analysis Repository#security_and_analysis}
	SecurityAndAnalysis *RepositorySecurityAndAnalysis `field:"optional" json:"securityAndAnalysis" yaml:"securityAndAnalysis"`
	// Can be 'PR_BODY', 'COMMIT_MESSAGES', or 'BLANK' for a default squash merge commit message.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#squash_merge_commit_message Repository#squash_merge_commit_message}
	SquashMergeCommitMessage *string `field:"optional" json:"squashMergeCommitMessage" yaml:"squashMergeCommitMessage"`
	// Can be 'PR_TITLE' or 'COMMIT_OR_PR_TITLE' for a default squash merge commit title.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#squash_merge_commit_title Repository#squash_merge_commit_title}
	SquashMergeCommitTitle *string `field:"optional" json:"squashMergeCommitTitle" yaml:"squashMergeCommitTitle"`
	// template block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#template Repository#template}
	Template *RepositoryTemplate `field:"optional" json:"template" yaml:"template"`
	// The list of topics of the repository.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#topics Repository#topics}
	Topics *[]*string `field:"optional" json:"topics" yaml:"topics"`
	// Can be 'public' or 'private'.
	//
	// If your organization is associated with an enterprise account using GitHub Enterprise Cloud or GitHub Enterprise Server 2.20+, visibility can also be 'internal'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#visibility Repository#visibility}
	Visibility *string `field:"optional" json:"visibility" yaml:"visibility"`
	// Set to 'true' to enable security alerts for vulnerable dependencies.
	//
	// Enabling requires alerts to be enabled on the owner level. (Note for importing: GitHub enables the alerts on public repos but disables them on private repos by default). Note that vulnerability alerts have not been successfully tested on any GitHub Enterprise instance and may be unavailable in those settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs/resources/repository#vulnerability_alerts Repository#vulnerability_alerts}
	VulnerabilityAlerts interface{} `field:"optional" json:"vulnerabilityAlerts" yaml:"vulnerabilityAlerts"`
}

