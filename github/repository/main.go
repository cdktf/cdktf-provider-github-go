package repository

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-github.repository.Repository",
		reflect.TypeOf((*Repository)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowAutoMerge", GoGetter: "AllowAutoMerge"},
			_jsii_.MemberProperty{JsiiProperty: "allowAutoMergeInput", GoGetter: "AllowAutoMergeInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowMergeCommit", GoGetter: "AllowMergeCommit"},
			_jsii_.MemberProperty{JsiiProperty: "allowMergeCommitInput", GoGetter: "AllowMergeCommitInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowRebaseMerge", GoGetter: "AllowRebaseMerge"},
			_jsii_.MemberProperty{JsiiProperty: "allowRebaseMergeInput", GoGetter: "AllowRebaseMergeInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowSquashMerge", GoGetter: "AllowSquashMerge"},
			_jsii_.MemberProperty{JsiiProperty: "allowSquashMergeInput", GoGetter: "AllowSquashMergeInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowUpdateBranch", GoGetter: "AllowUpdateBranch"},
			_jsii_.MemberProperty{JsiiProperty: "allowUpdateBranchInput", GoGetter: "AllowUpdateBranchInput"},
			_jsii_.MemberProperty{JsiiProperty: "archived", GoGetter: "Archived"},
			_jsii_.MemberProperty{JsiiProperty: "archivedInput", GoGetter: "ArchivedInput"},
			_jsii_.MemberProperty{JsiiProperty: "archiveOnDestroy", GoGetter: "ArchiveOnDestroy"},
			_jsii_.MemberProperty{JsiiProperty: "archiveOnDestroyInput", GoGetter: "ArchiveOnDestroyInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoInit", GoGetter: "AutoInit"},
			_jsii_.MemberProperty{JsiiProperty: "autoInitInput", GoGetter: "AutoInitInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranch", GoGetter: "DefaultBranch"},
			_jsii_.MemberProperty{JsiiProperty: "defaultBranchInput", GoGetter: "DefaultBranchInput"},
			_jsii_.MemberProperty{JsiiProperty: "deleteBranchOnMerge", GoGetter: "DeleteBranchOnMerge"},
			_jsii_.MemberProperty{JsiiProperty: "deleteBranchOnMergeInput", GoGetter: "DeleteBranchOnMergeInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "fullName", GoGetter: "FullName"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gitCloneUrl", GoGetter: "GitCloneUrl"},
			_jsii_.MemberProperty{JsiiProperty: "gitignoreTemplate", GoGetter: "GitignoreTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "gitignoreTemplateInput", GoGetter: "GitignoreTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "hasDiscussions", GoGetter: "HasDiscussions"},
			_jsii_.MemberProperty{JsiiProperty: "hasDiscussionsInput", GoGetter: "HasDiscussionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "hasDownloads", GoGetter: "HasDownloads"},
			_jsii_.MemberProperty{JsiiProperty: "hasDownloadsInput", GoGetter: "HasDownloadsInput"},
			_jsii_.MemberProperty{JsiiProperty: "hasIssues", GoGetter: "HasIssues"},
			_jsii_.MemberProperty{JsiiProperty: "hasIssuesInput", GoGetter: "HasIssuesInput"},
			_jsii_.MemberProperty{JsiiProperty: "hasProjects", GoGetter: "HasProjects"},
			_jsii_.MemberProperty{JsiiProperty: "hasProjectsInput", GoGetter: "HasProjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "hasWiki", GoGetter: "HasWiki"},
			_jsii_.MemberProperty{JsiiProperty: "hasWikiInput", GoGetter: "HasWikiInput"},
			_jsii_.MemberProperty{JsiiProperty: "homepageUrl", GoGetter: "HomepageUrl"},
			_jsii_.MemberProperty{JsiiProperty: "homepageUrlInput", GoGetter: "HomepageUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "htmlUrl", GoGetter: "HtmlUrl"},
			_jsii_.MemberProperty{JsiiProperty: "httpCloneUrl", GoGetter: "HttpCloneUrl"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreVulnerabilityAlertsDuringRead", GoGetter: "IgnoreVulnerabilityAlertsDuringRead"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreVulnerabilityAlertsDuringReadInput", GoGetter: "IgnoreVulnerabilityAlertsDuringReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isTemplate", GoGetter: "IsTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "isTemplateInput", GoGetter: "IsTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "licenseTemplate", GoGetter: "LicenseTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "licenseTemplateInput", GoGetter: "LicenseTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "mergeCommitMessage", GoGetter: "MergeCommitMessage"},
			_jsii_.MemberProperty{JsiiProperty: "mergeCommitMessageInput", GoGetter: "MergeCommitMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "mergeCommitTitle", GoGetter: "MergeCommitTitle"},
			_jsii_.MemberProperty{JsiiProperty: "mergeCommitTitleInput", GoGetter: "MergeCommitTitleInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeId", GoGetter: "NodeId"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pages", GoGetter: "Pages"},
			_jsii_.MemberProperty{JsiiProperty: "pagesInput", GoGetter: "PagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "private", GoGetter: "Private"},
			_jsii_.MemberProperty{JsiiProperty: "privateInput", GoGetter: "PrivateInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putPages", GoMethod: "PutPages"},
			_jsii_.MemberMethod{JsiiMethod: "putSecurityAndAnalysis", GoMethod: "PutSecurityAndAnalysis"},
			_jsii_.MemberMethod{JsiiMethod: "putTemplate", GoMethod: "PutTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "repoId", GoGetter: "RepoId"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowAutoMerge", GoMethod: "ResetAllowAutoMerge"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowMergeCommit", GoMethod: "ResetAllowMergeCommit"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowRebaseMerge", GoMethod: "ResetAllowRebaseMerge"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowSquashMerge", GoMethod: "ResetAllowSquashMerge"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowUpdateBranch", GoMethod: "ResetAllowUpdateBranch"},
			_jsii_.MemberMethod{JsiiMethod: "resetArchived", GoMethod: "ResetArchived"},
			_jsii_.MemberMethod{JsiiMethod: "resetArchiveOnDestroy", GoMethod: "ResetArchiveOnDestroy"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoInit", GoMethod: "ResetAutoInit"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultBranch", GoMethod: "ResetDefaultBranch"},
			_jsii_.MemberMethod{JsiiMethod: "resetDeleteBranchOnMerge", GoMethod: "ResetDeleteBranchOnMerge"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetGitignoreTemplate", GoMethod: "ResetGitignoreTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetHasDiscussions", GoMethod: "ResetHasDiscussions"},
			_jsii_.MemberMethod{JsiiMethod: "resetHasDownloads", GoMethod: "ResetHasDownloads"},
			_jsii_.MemberMethod{JsiiMethod: "resetHasIssues", GoMethod: "ResetHasIssues"},
			_jsii_.MemberMethod{JsiiMethod: "resetHasProjects", GoMethod: "ResetHasProjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetHasWiki", GoMethod: "ResetHasWiki"},
			_jsii_.MemberMethod{JsiiMethod: "resetHomepageUrl", GoMethod: "ResetHomepageUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreVulnerabilityAlertsDuringRead", GoMethod: "ResetIgnoreVulnerabilityAlertsDuringRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsTemplate", GoMethod: "ResetIsTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetLicenseTemplate", GoMethod: "ResetLicenseTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetMergeCommitMessage", GoMethod: "ResetMergeCommitMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetMergeCommitTitle", GoMethod: "ResetMergeCommitTitle"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPages", GoMethod: "ResetPages"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivate", GoMethod: "ResetPrivate"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityAndAnalysis", GoMethod: "ResetSecurityAndAnalysis"},
			_jsii_.MemberMethod{JsiiMethod: "resetSquashMergeCommitMessage", GoMethod: "ResetSquashMergeCommitMessage"},
			_jsii_.MemberMethod{JsiiMethod: "resetSquashMergeCommitTitle", GoMethod: "ResetSquashMergeCommitTitle"},
			_jsii_.MemberMethod{JsiiMethod: "resetTemplate", GoMethod: "ResetTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetTopics", GoMethod: "ResetTopics"},
			_jsii_.MemberMethod{JsiiMethod: "resetVisibility", GoMethod: "ResetVisibility"},
			_jsii_.MemberMethod{JsiiMethod: "resetVulnerabilityAlerts", GoMethod: "ResetVulnerabilityAlerts"},
			_jsii_.MemberProperty{JsiiProperty: "securityAndAnalysis", GoGetter: "SecurityAndAnalysis"},
			_jsii_.MemberProperty{JsiiProperty: "securityAndAnalysisInput", GoGetter: "SecurityAndAnalysisInput"},
			_jsii_.MemberProperty{JsiiProperty: "squashMergeCommitMessage", GoGetter: "SquashMergeCommitMessage"},
			_jsii_.MemberProperty{JsiiProperty: "squashMergeCommitMessageInput", GoGetter: "SquashMergeCommitMessageInput"},
			_jsii_.MemberProperty{JsiiProperty: "squashMergeCommitTitle", GoGetter: "SquashMergeCommitTitle"},
			_jsii_.MemberProperty{JsiiProperty: "squashMergeCommitTitleInput", GoGetter: "SquashMergeCommitTitleInput"},
			_jsii_.MemberProperty{JsiiProperty: "sshCloneUrl", GoGetter: "SshCloneUrl"},
			_jsii_.MemberProperty{JsiiProperty: "svnUrl", GoGetter: "SvnUrl"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "template", GoGetter: "Template"},
			_jsii_.MemberProperty{JsiiProperty: "templateInput", GoGetter: "TemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "topics", GoGetter: "Topics"},
			_jsii_.MemberProperty{JsiiProperty: "topicsInput", GoGetter: "TopicsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "visibility", GoGetter: "Visibility"},
			_jsii_.MemberProperty{JsiiProperty: "visibilityInput", GoGetter: "VisibilityInput"},
			_jsii_.MemberProperty{JsiiProperty: "vulnerabilityAlerts", GoGetter: "VulnerabilityAlerts"},
			_jsii_.MemberProperty{JsiiProperty: "vulnerabilityAlertsInput", GoGetter: "VulnerabilityAlertsInput"},
		},
		func() interface{} {
			j := jsiiProxy_Repository{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.repository.RepositoryConfig",
		reflect.TypeOf((*RepositoryConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.repository.RepositoryPages",
		reflect.TypeOf((*RepositoryPages)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.repository.RepositoryPagesOutputReference",
		reflect.TypeOf((*RepositoryPagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "buildType", GoGetter: "BuildType"},
			_jsii_.MemberProperty{JsiiProperty: "buildTypeInput", GoGetter: "BuildTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "cname", GoGetter: "Cname"},
			_jsii_.MemberProperty{JsiiProperty: "cnameInput", GoGetter: "CnameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "custom404", GoGetter: "Custom404"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "htmlUrl", GoGetter: "HtmlUrl"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putSource", GoMethod: "PutSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetBuildType", GoMethod: "ResetBuildType"},
			_jsii_.MemberMethod{JsiiMethod: "resetCname", GoMethod: "ResetCname"},
			_jsii_.MemberMethod{JsiiMethod: "resetSource", GoMethod: "ResetSource"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberProperty{JsiiProperty: "sourceInput", GoGetter: "SourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
		},
		func() interface{} {
			j := jsiiProxy_RepositoryPagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.repository.RepositoryPagesSource",
		reflect.TypeOf((*RepositoryPagesSource)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.repository.RepositoryPagesSourceOutputReference",
		reflect.TypeOf((*RepositoryPagesSourceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "branch", GoGetter: "Branch"},
			_jsii_.MemberProperty{JsiiProperty: "branchInput", GoGetter: "BranchInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RepositoryPagesSourceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysis",
		reflect.TypeOf((*RepositorySecurityAndAnalysis)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysisAdvancedSecurity",
		reflect.TypeOf((*RepositorySecurityAndAnalysisAdvancedSecurity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysisAdvancedSecurityOutputReference",
		reflect.TypeOf((*RepositorySecurityAndAnalysisAdvancedSecurityOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RepositorySecurityAndAnalysisAdvancedSecurityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysisOutputReference",
		reflect.TypeOf((*RepositorySecurityAndAnalysisOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurity", GoGetter: "AdvancedSecurity"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurityInput", GoGetter: "AdvancedSecurityInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putAdvancedSecurity", GoMethod: "PutAdvancedSecurity"},
			_jsii_.MemberMethod{JsiiMethod: "putSecretScanning", GoMethod: "PutSecretScanning"},
			_jsii_.MemberMethod{JsiiMethod: "putSecretScanningPushProtection", GoMethod: "PutSecretScanningPushProtection"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdvancedSecurity", GoMethod: "ResetAdvancedSecurity"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretScanning", GoMethod: "ResetSecretScanning"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretScanningPushProtection", GoMethod: "ResetSecretScanningPushProtection"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanning", GoGetter: "SecretScanning"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanningInput", GoGetter: "SecretScanningInput"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanningPushProtection", GoGetter: "SecretScanningPushProtection"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanningPushProtectionInput", GoGetter: "SecretScanningPushProtectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RepositorySecurityAndAnalysisOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysisSecretScanning",
		reflect.TypeOf((*RepositorySecurityAndAnalysisSecretScanning)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysisSecretScanningOutputReference",
		reflect.TypeOf((*RepositorySecurityAndAnalysisSecretScanningOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RepositorySecurityAndAnalysisSecretScanningOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysisSecretScanningPushProtection",
		reflect.TypeOf((*RepositorySecurityAndAnalysisSecretScanningPushProtection)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysisSecretScanningPushProtectionOutputReference",
		reflect.TypeOf((*RepositorySecurityAndAnalysisSecretScanningPushProtectionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RepositorySecurityAndAnalysisSecretScanningPushProtectionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.repository.RepositoryTemplate",
		reflect.TypeOf((*RepositoryTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.repository.RepositoryTemplateOutputReference",
		reflect.TypeOf((*RepositoryTemplateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeAllBranches", GoGetter: "IncludeAllBranches"},
			_jsii_.MemberProperty{JsiiProperty: "includeAllBranchesInput", GoGetter: "IncludeAllBranchesInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "ownerInput", GoGetter: "OwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "repository", GoGetter: "Repository"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryInput", GoGetter: "RepositoryInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeAllBranches", GoMethod: "ResetIncludeAllBranches"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_RepositoryTemplateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
