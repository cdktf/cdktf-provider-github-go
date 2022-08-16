// Prebuilt github Provider for Terraform CDK (cdktf)
package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-github-go/github/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-github-go/github/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/github/r/repository github_repository}.
type Repository interface {
	cdktf.TerraformResource
	AllowAutoMerge() interface{}
	SetAllowAutoMerge(val interface{})
	AllowAutoMergeInput() interface{}
	AllowMergeCommit() interface{}
	SetAllowMergeCommit(val interface{})
	AllowMergeCommitInput() interface{}
	AllowRebaseMerge() interface{}
	SetAllowRebaseMerge(val interface{})
	AllowRebaseMergeInput() interface{}
	AllowSquashMerge() interface{}
	SetAllowSquashMerge(val interface{})
	AllowSquashMergeInput() interface{}
	Archived() interface{}
	SetArchived(val interface{})
	ArchivedInput() interface{}
	ArchiveOnDestroy() interface{}
	SetArchiveOnDestroy(val interface{})
	ArchiveOnDestroyInput() interface{}
	AutoInit() interface{}
	SetAutoInit(val interface{})
	AutoInitInput() interface{}
	Branches() RepositoryBranchesList
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultBranch() *string
	SetDefaultBranch(val *string)
	DefaultBranchInput() *string
	DeleteBranchOnMerge() interface{}
	SetDeleteBranchOnMerge(val interface{})
	DeleteBranchOnMergeInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Etag() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	FullName() *string
	GitCloneUrl() *string
	GitignoreTemplate() *string
	SetGitignoreTemplate(val *string)
	GitignoreTemplateInput() *string
	HasDownloads() interface{}
	SetHasDownloads(val interface{})
	HasDownloadsInput() interface{}
	HasIssues() interface{}
	SetHasIssues(val interface{})
	HasIssuesInput() interface{}
	HasProjects() interface{}
	SetHasProjects(val interface{})
	HasProjectsInput() interface{}
	HasWiki() interface{}
	SetHasWiki(val interface{})
	HasWikiInput() interface{}
	HomepageUrl() *string
	SetHomepageUrl(val *string)
	HomepageUrlInput() *string
	HtmlUrl() *string
	HttpCloneUrl() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IgnoreVulnerabilityAlertsDuringRead() interface{}
	SetIgnoreVulnerabilityAlertsDuringRead(val interface{})
	IgnoreVulnerabilityAlertsDuringReadInput() interface{}
	IsTemplate() interface{}
	SetIsTemplate(val interface{})
	IsTemplateInput() interface{}
	LicenseTemplate() *string
	SetLicenseTemplate(val *string)
	LicenseTemplateInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
	NodeId() *string
	Pages() RepositoryPagesOutputReference
	PagesInput() *RepositoryPages
	Private() interface{}
	SetPrivate(val interface{})
	PrivateInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	RepoId() *float64
	SshCloneUrl() *string
	SvnUrl() *string
	Template() RepositoryTemplateOutputReference
	TemplateInput() *RepositoryTemplate
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Topics() *[]*string
	SetTopics(val *[]*string)
	TopicsInput() *[]*string
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
	VulnerabilityAlerts() interface{}
	SetVulnerabilityAlerts(val interface{})
	VulnerabilityAlertsInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutPages(value *RepositoryPages)
	PutTemplate(value *RepositoryTemplate)
	ResetAllowAutoMerge()
	ResetAllowMergeCommit()
	ResetAllowRebaseMerge()
	ResetAllowSquashMerge()
	ResetArchived()
	ResetArchiveOnDestroy()
	ResetAutoInit()
	ResetDefaultBranch()
	ResetDeleteBranchOnMerge()
	ResetDescription()
	ResetGitignoreTemplate()
	ResetHasDownloads()
	ResetHasIssues()
	ResetHasProjects()
	ResetHasWiki()
	ResetHomepageUrl()
	ResetId()
	ResetIgnoreVulnerabilityAlertsDuringRead()
	ResetIsTemplate()
	ResetLicenseTemplate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPages()
	ResetPrivate()
	ResetTemplate()
	ResetTopics()
	ResetVisibility()
	ResetVulnerabilityAlerts()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Repository
type jsiiProxy_Repository struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_Repository) AllowAutoMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAutoMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) AllowAutoMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAutoMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) AllowMergeCommit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMergeCommit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) AllowMergeCommitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowMergeCommitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) AllowRebaseMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRebaseMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) AllowRebaseMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowRebaseMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) AllowSquashMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSquashMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) AllowSquashMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowSquashMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Archived() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archived",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) ArchivedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archivedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) ArchiveOnDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveOnDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) ArchiveOnDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"archiveOnDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) AutoInit() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoInit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) AutoInitInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoInitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Branches() RepositoryBranchesList {
	var returns RepositoryBranchesList
	_jsii_.Get(
		j,
		"branches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) DefaultBranch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) DefaultBranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) DeleteBranchOnMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteBranchOnMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) DeleteBranchOnMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deleteBranchOnMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) FullName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) GitCloneUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitCloneUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) GitignoreTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitignoreTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) GitignoreTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gitignoreTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HasDownloads() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasDownloads",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HasDownloadsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasDownloadsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HasIssues() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasIssues",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HasIssuesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasIssuesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HasProjects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HasProjectsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HasWiki() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasWiki",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HasWikiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasWikiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HomepageUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homepageUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HomepageUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"homepageUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HtmlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) HttpCloneUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"httpCloneUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) IgnoreVulnerabilityAlertsDuringRead() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreVulnerabilityAlertsDuringRead",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) IgnoreVulnerabilityAlertsDuringReadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ignoreVulnerabilityAlertsDuringReadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) IsTemplate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) IsTemplateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) LicenseTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) LicenseTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) NodeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nodeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Pages() RepositoryPagesOutputReference {
	var returns RepositoryPagesOutputReference
	_jsii_.Get(
		j,
		"pages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) PagesInput() *RepositoryPages {
	var returns *RepositoryPages
	_jsii_.Get(
		j,
		"pagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Private() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"private",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) PrivateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"privateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) RepoId() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"repoId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) SshCloneUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sshCloneUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) SvnUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"svnUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Template() RepositoryTemplateOutputReference {
	var returns RepositoryTemplateOutputReference
	_jsii_.Get(
		j,
		"template",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) TemplateInput() *RepositoryTemplate {
	var returns *RepositoryTemplate
	_jsii_.Get(
		j,
		"templateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Topics() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) TopicsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"topicsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) VulnerabilityAlerts() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilityAlerts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Repository) VulnerabilityAlertsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vulnerabilityAlertsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/github/r/repository github_repository} Resource.
func NewRepository(scope constructs.Construct, id *string, config *RepositoryConfig) Repository {
	_init_.Initialize()

	j := jsiiProxy_Repository{}

	_jsii_.Create(
		"@cdktf/provider-github.Repository",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/github/r/repository github_repository} Resource.
func NewRepository_Override(r Repository, scope constructs.Construct, id *string, config *RepositoryConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.Repository",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_Repository) SetAllowAutoMerge(val interface{}) {
	_jsii_.Set(
		j,
		"allowAutoMerge",
		val,
	)
}

func (j *jsiiProxy_Repository) SetAllowMergeCommit(val interface{}) {
	_jsii_.Set(
		j,
		"allowMergeCommit",
		val,
	)
}

func (j *jsiiProxy_Repository) SetAllowRebaseMerge(val interface{}) {
	_jsii_.Set(
		j,
		"allowRebaseMerge",
		val,
	)
}

func (j *jsiiProxy_Repository) SetAllowSquashMerge(val interface{}) {
	_jsii_.Set(
		j,
		"allowSquashMerge",
		val,
	)
}

func (j *jsiiProxy_Repository) SetArchived(val interface{}) {
	_jsii_.Set(
		j,
		"archived",
		val,
	)
}

func (j *jsiiProxy_Repository) SetArchiveOnDestroy(val interface{}) {
	_jsii_.Set(
		j,
		"archiveOnDestroy",
		val,
	)
}

func (j *jsiiProxy_Repository) SetAutoInit(val interface{}) {
	_jsii_.Set(
		j,
		"autoInit",
		val,
	)
}

func (j *jsiiProxy_Repository) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_Repository) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_Repository) SetDefaultBranch(val *string) {
	_jsii_.Set(
		j,
		"defaultBranch",
		val,
	)
}

func (j *jsiiProxy_Repository) SetDeleteBranchOnMerge(val interface{}) {
	_jsii_.Set(
		j,
		"deleteBranchOnMerge",
		val,
	)
}

func (j *jsiiProxy_Repository) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Repository) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_Repository) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Repository) SetGitignoreTemplate(val *string) {
	_jsii_.Set(
		j,
		"gitignoreTemplate",
		val,
	)
}

func (j *jsiiProxy_Repository) SetHasDownloads(val interface{}) {
	_jsii_.Set(
		j,
		"hasDownloads",
		val,
	)
}

func (j *jsiiProxy_Repository) SetHasIssues(val interface{}) {
	_jsii_.Set(
		j,
		"hasIssues",
		val,
	)
}

func (j *jsiiProxy_Repository) SetHasProjects(val interface{}) {
	_jsii_.Set(
		j,
		"hasProjects",
		val,
	)
}

func (j *jsiiProxy_Repository) SetHasWiki(val interface{}) {
	_jsii_.Set(
		j,
		"hasWiki",
		val,
	)
}

func (j *jsiiProxy_Repository) SetHomepageUrl(val *string) {
	_jsii_.Set(
		j,
		"homepageUrl",
		val,
	)
}

func (j *jsiiProxy_Repository) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_Repository) SetIgnoreVulnerabilityAlertsDuringRead(val interface{}) {
	_jsii_.Set(
		j,
		"ignoreVulnerabilityAlertsDuringRead",
		val,
	)
}

func (j *jsiiProxy_Repository) SetIsTemplate(val interface{}) {
	_jsii_.Set(
		j,
		"isTemplate",
		val,
	)
}

func (j *jsiiProxy_Repository) SetLicenseTemplate(val *string) {
	_jsii_.Set(
		j,
		"licenseTemplate",
		val,
	)
}

func (j *jsiiProxy_Repository) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_Repository) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_Repository) SetPrivate(val interface{}) {
	_jsii_.Set(
		j,
		"private",
		val,
	)
}

func (j *jsiiProxy_Repository) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_Repository) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_Repository) SetTopics(val *[]*string) {
	_jsii_.Set(
		j,
		"topics",
		val,
	)
}

func (j *jsiiProxy_Repository) SetVisibility(val *string) {
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

func (j *jsiiProxy_Repository) SetVulnerabilityAlerts(val interface{}) {
	_jsii_.Set(
		j,
		"vulnerabilityAlerts",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Repository_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.Repository",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func Repository_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.Repository",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_Repository) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_Repository) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_Repository) PutPages(value *RepositoryPages) {
	_jsii_.InvokeVoid(
		r,
		"putPages",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Repository) PutTemplate(value *RepositoryTemplate) {
	_jsii_.InvokeVoid(
		r,
		"putTemplate",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_Repository) ResetAllowAutoMerge() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowAutoMerge",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetAllowMergeCommit() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowMergeCommit",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetAllowRebaseMerge() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowRebaseMerge",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetAllowSquashMerge() {
	_jsii_.InvokeVoid(
		r,
		"resetAllowSquashMerge",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetArchived() {
	_jsii_.InvokeVoid(
		r,
		"resetArchived",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetArchiveOnDestroy() {
	_jsii_.InvokeVoid(
		r,
		"resetArchiveOnDestroy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetAutoInit() {
	_jsii_.InvokeVoid(
		r,
		"resetAutoInit",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetDefaultBranch() {
	_jsii_.InvokeVoid(
		r,
		"resetDefaultBranch",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetDeleteBranchOnMerge() {
	_jsii_.InvokeVoid(
		r,
		"resetDeleteBranchOnMerge",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetDescription() {
	_jsii_.InvokeVoid(
		r,
		"resetDescription",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetGitignoreTemplate() {
	_jsii_.InvokeVoid(
		r,
		"resetGitignoreTemplate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetHasDownloads() {
	_jsii_.InvokeVoid(
		r,
		"resetHasDownloads",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetHasIssues() {
	_jsii_.InvokeVoid(
		r,
		"resetHasIssues",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetHasProjects() {
	_jsii_.InvokeVoid(
		r,
		"resetHasProjects",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetHasWiki() {
	_jsii_.InvokeVoid(
		r,
		"resetHasWiki",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetHomepageUrl() {
	_jsii_.InvokeVoid(
		r,
		"resetHomepageUrl",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetIgnoreVulnerabilityAlertsDuringRead() {
	_jsii_.InvokeVoid(
		r,
		"resetIgnoreVulnerabilityAlertsDuringRead",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetIsTemplate() {
	_jsii_.InvokeVoid(
		r,
		"resetIsTemplate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetLicenseTemplate() {
	_jsii_.InvokeVoid(
		r,
		"resetLicenseTemplate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetPages() {
	_jsii_.InvokeVoid(
		r,
		"resetPages",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetPrivate() {
	_jsii_.InvokeVoid(
		r,
		"resetPrivate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetTemplate() {
	_jsii_.InvokeVoid(
		r,
		"resetTemplate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetTopics() {
	_jsii_.InvokeVoid(
		r,
		"resetTopics",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetVisibility() {
	_jsii_.InvokeVoid(
		r,
		"resetVisibility",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) ResetVulnerabilityAlerts() {
	_jsii_.InvokeVoid(
		r,
		"resetVulnerabilityAlerts",
		nil, // no parameters
	)
}

func (r *jsiiProxy_Repository) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_Repository) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
