package organizationsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v10/organizationsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/resources/organization_settings github_organization_settings}.
type OrganizationSettings interface {
	cdktf.TerraformResource
	AdvancedSecurityEnabledForNewRepositories() interface{}
	SetAdvancedSecurityEnabledForNewRepositories(val interface{})
	AdvancedSecurityEnabledForNewRepositoriesInput() interface{}
	BillingEmail() *string
	SetBillingEmail(val *string)
	BillingEmailInput() *string
	Blog() *string
	SetBlog(val *string)
	BlogInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Company() *string
	SetCompany(val *string)
	CompanyInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	DefaultRepositoryPermission() *string
	SetDefaultRepositoryPermission(val *string)
	DefaultRepositoryPermissionInput() *string
	DependabotAlertsEnabledForNewRepositories() interface{}
	SetDependabotAlertsEnabledForNewRepositories(val interface{})
	DependabotAlertsEnabledForNewRepositoriesInput() interface{}
	DependabotSecurityUpdatesEnabledForNewRepositories() interface{}
	SetDependabotSecurityUpdatesEnabledForNewRepositories(val interface{})
	DependabotSecurityUpdatesEnabledForNewRepositoriesInput() interface{}
	DependencyGraphEnabledForNewRepositories() interface{}
	SetDependencyGraphEnabledForNewRepositories(val interface{})
	DependencyGraphEnabledForNewRepositoriesInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Email() *string
	SetEmail(val *string)
	EmailInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HasOrganizationProjects() interface{}
	SetHasOrganizationProjects(val interface{})
	HasOrganizationProjectsInput() interface{}
	HasRepositoryProjects() interface{}
	SetHasRepositoryProjects(val interface{})
	HasRepositoryProjectsInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MembersCanCreateInternalRepositories() interface{}
	SetMembersCanCreateInternalRepositories(val interface{})
	MembersCanCreateInternalRepositoriesInput() interface{}
	MembersCanCreatePages() interface{}
	SetMembersCanCreatePages(val interface{})
	MembersCanCreatePagesInput() interface{}
	MembersCanCreatePrivatePages() interface{}
	SetMembersCanCreatePrivatePages(val interface{})
	MembersCanCreatePrivatePagesInput() interface{}
	MembersCanCreatePrivateRepositories() interface{}
	SetMembersCanCreatePrivateRepositories(val interface{})
	MembersCanCreatePrivateRepositoriesInput() interface{}
	MembersCanCreatePublicPages() interface{}
	SetMembersCanCreatePublicPages(val interface{})
	MembersCanCreatePublicPagesInput() interface{}
	MembersCanCreatePublicRepositories() interface{}
	SetMembersCanCreatePublicRepositories(val interface{})
	MembersCanCreatePublicRepositoriesInput() interface{}
	MembersCanCreateRepositories() interface{}
	SetMembersCanCreateRepositories(val interface{})
	MembersCanCreateRepositoriesInput() interface{}
	MembersCanForkPrivateRepositories() interface{}
	SetMembersCanForkPrivateRepositories(val interface{})
	MembersCanForkPrivateRepositoriesInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	// The tree node.
	Node() constructs.Node
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
	SecretScanningEnabledForNewRepositories() interface{}
	SetSecretScanningEnabledForNewRepositories(val interface{})
	SecretScanningEnabledForNewRepositoriesInput() interface{}
	SecretScanningPushProtectionEnabledForNewRepositories() interface{}
	SetSecretScanningPushProtectionEnabledForNewRepositories(val interface{})
	SecretScanningPushProtectionEnabledForNewRepositoriesInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TwitterUsername() *string
	SetTwitterUsername(val *string)
	TwitterUsernameInput() *string
	WebCommitSignoffRequired() interface{}
	SetWebCommitSignoffRequired(val interface{})
	WebCommitSignoffRequiredInput() interface{}
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
	ResetAdvancedSecurityEnabledForNewRepositories()
	ResetBlog()
	ResetCompany()
	ResetDefaultRepositoryPermission()
	ResetDependabotAlertsEnabledForNewRepositories()
	ResetDependabotSecurityUpdatesEnabledForNewRepositories()
	ResetDependencyGraphEnabledForNewRepositories()
	ResetDescription()
	ResetEmail()
	ResetHasOrganizationProjects()
	ResetHasRepositoryProjects()
	ResetId()
	ResetLocation()
	ResetMembersCanCreateInternalRepositories()
	ResetMembersCanCreatePages()
	ResetMembersCanCreatePrivatePages()
	ResetMembersCanCreatePrivateRepositories()
	ResetMembersCanCreatePublicPages()
	ResetMembersCanCreatePublicRepositories()
	ResetMembersCanCreateRepositories()
	ResetMembersCanForkPrivateRepositories()
	ResetName()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecretScanningEnabledForNewRepositories()
	ResetSecretScanningPushProtectionEnabledForNewRepositories()
	ResetTwitterUsername()
	ResetWebCommitSignoffRequired()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for OrganizationSettings
type jsiiProxy_OrganizationSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_OrganizationSettings) AdvancedSecurityEnabledForNewRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedSecurityEnabledForNewRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) AdvancedSecurityEnabledForNewRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedSecurityEnabledForNewRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) BillingEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) BillingEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Blog() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) BlogInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"blogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Company() *string {
	var returns *string
	_jsii_.Get(
		j,
		"company",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) CompanyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"companyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DefaultRepositoryPermission() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRepositoryPermission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DefaultRepositoryPermissionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRepositoryPermissionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DependabotAlertsEnabledForNewRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependabotAlertsEnabledForNewRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DependabotAlertsEnabledForNewRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependabotAlertsEnabledForNewRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DependabotSecurityUpdatesEnabledForNewRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependabotSecurityUpdatesEnabledForNewRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DependabotSecurityUpdatesEnabledForNewRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependabotSecurityUpdatesEnabledForNewRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DependencyGraphEnabledForNewRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependencyGraphEnabledForNewRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DependencyGraphEnabledForNewRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dependencyGraphEnabledForNewRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Email() *string {
	var returns *string
	_jsii_.Get(
		j,
		"email",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) EmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) HasOrganizationProjects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasOrganizationProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) HasOrganizationProjectsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasOrganizationProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) HasRepositoryProjects() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasRepositoryProjects",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) HasRepositoryProjectsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hasRepositoryProjectsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreateInternalRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreateInternalRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreateInternalRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreateInternalRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePrivatePages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePrivatePages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePrivatePagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePrivatePagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePrivateRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePrivateRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePrivateRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePrivateRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePublicPages() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePublicPages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePublicPagesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePublicPagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePublicRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePublicRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreatePublicRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreatePublicRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreateRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreateRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanCreateRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanCreateRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanForkPrivateRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanForkPrivateRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) MembersCanForkPrivateRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"membersCanForkPrivateRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) SecretScanningEnabledForNewRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningEnabledForNewRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) SecretScanningEnabledForNewRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningEnabledForNewRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) SecretScanningPushProtectionEnabledForNewRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningPushProtectionEnabledForNewRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) SecretScanningPushProtectionEnabledForNewRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningPushProtectionEnabledForNewRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) TwitterUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twitterUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) TwitterUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"twitterUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) WebCommitSignoffRequired() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webCommitSignoffRequired",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationSettings) WebCommitSignoffRequiredInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"webCommitSignoffRequiredInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/resources/organization_settings github_organization_settings} Resource.
func NewOrganizationSettings(scope constructs.Construct, id *string, config *OrganizationSettingsConfig) OrganizationSettings {
	_init_.Initialize()

	if err := validateNewOrganizationSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationSettings{}

	_jsii_.Create(
		"@cdktf/provider-github.organizationSettings.OrganizationSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/resources/organization_settings github_organization_settings} Resource.
func NewOrganizationSettings_Override(o OrganizationSettings, scope constructs.Construct, id *string, config *OrganizationSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.organizationSettings.OrganizationSettings",
		[]interface{}{scope, id, config},
		o,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetAdvancedSecurityEnabledForNewRepositories(val interface{}) {
	if err := j.validateSetAdvancedSecurityEnabledForNewRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedSecurityEnabledForNewRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetBillingEmail(val *string) {
	if err := j.validateSetBillingEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingEmail",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetBlog(val *string) {
	if err := j.validateSetBlogParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blog",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetCompany(val *string) {
	if err := j.validateSetCompanyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"company",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetDefaultRepositoryPermission(val *string) {
	if err := j.validateSetDefaultRepositoryPermissionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRepositoryPermission",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetDependabotAlertsEnabledForNewRepositories(val interface{}) {
	if err := j.validateSetDependabotAlertsEnabledForNewRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependabotAlertsEnabledForNewRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetDependabotSecurityUpdatesEnabledForNewRepositories(val interface{}) {
	if err := j.validateSetDependabotSecurityUpdatesEnabledForNewRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependabotSecurityUpdatesEnabledForNewRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetDependencyGraphEnabledForNewRepositories(val interface{}) {
	if err := j.validateSetDependencyGraphEnabledForNewRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dependencyGraphEnabledForNewRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetEmail(val *string) {
	if err := j.validateSetEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"email",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetHasOrganizationProjects(val interface{}) {
	if err := j.validateSetHasOrganizationProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasOrganizationProjects",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetHasRepositoryProjects(val interface{}) {
	if err := j.validateSetHasRepositoryProjectsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasRepositoryProjects",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetMembersCanCreateInternalRepositories(val interface{}) {
	if err := j.validateSetMembersCanCreateInternalRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membersCanCreateInternalRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetMembersCanCreatePages(val interface{}) {
	if err := j.validateSetMembersCanCreatePagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membersCanCreatePages",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetMembersCanCreatePrivatePages(val interface{}) {
	if err := j.validateSetMembersCanCreatePrivatePagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membersCanCreatePrivatePages",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetMembersCanCreatePrivateRepositories(val interface{}) {
	if err := j.validateSetMembersCanCreatePrivateRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membersCanCreatePrivateRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetMembersCanCreatePublicPages(val interface{}) {
	if err := j.validateSetMembersCanCreatePublicPagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membersCanCreatePublicPages",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetMembersCanCreatePublicRepositories(val interface{}) {
	if err := j.validateSetMembersCanCreatePublicRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membersCanCreatePublicRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetMembersCanCreateRepositories(val interface{}) {
	if err := j.validateSetMembersCanCreateRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membersCanCreateRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetMembersCanForkPrivateRepositories(val interface{}) {
	if err := j.validateSetMembersCanForkPrivateRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"membersCanForkPrivateRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetSecretScanningEnabledForNewRepositories(val interface{}) {
	if err := j.validateSetSecretScanningEnabledForNewRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretScanningEnabledForNewRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetSecretScanningPushProtectionEnabledForNewRepositories(val interface{}) {
	if err := j.validateSetSecretScanningPushProtectionEnabledForNewRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretScanningPushProtectionEnabledForNewRepositories",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetTwitterUsername(val *string) {
	if err := j.validateSetTwitterUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"twitterUsername",
		val,
	)
}

func (j *jsiiProxy_OrganizationSettings)SetWebCommitSignoffRequired(val interface{}) {
	if err := j.validateSetWebCommitSignoffRequiredParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"webCommitSignoffRequired",
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
func OrganizationSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrganizationSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.organizationSettings.OrganizationSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OrganizationSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrganizationSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.organizationSettings.OrganizationSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func OrganizationSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateOrganizationSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.organizationSettings.OrganizationSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func OrganizationSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.organizationSettings.OrganizationSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (o *jsiiProxy_OrganizationSettings) AddOverride(path *string, value interface{}) {
	if err := o.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (o *jsiiProxy_OrganizationSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) OverrideLogicalId(newLogicalId *string) {
	if err := o.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetAdvancedSecurityEnabledForNewRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetAdvancedSecurityEnabledForNewRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetBlog() {
	_jsii_.InvokeVoid(
		o,
		"resetBlog",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetCompany() {
	_jsii_.InvokeVoid(
		o,
		"resetCompany",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetDefaultRepositoryPermission() {
	_jsii_.InvokeVoid(
		o,
		"resetDefaultRepositoryPermission",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetDependabotAlertsEnabledForNewRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetDependabotAlertsEnabledForNewRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetDependabotSecurityUpdatesEnabledForNewRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetDependabotSecurityUpdatesEnabledForNewRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetDependencyGraphEnabledForNewRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetDependencyGraphEnabledForNewRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetDescription() {
	_jsii_.InvokeVoid(
		o,
		"resetDescription",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetEmail() {
	_jsii_.InvokeVoid(
		o,
		"resetEmail",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetHasOrganizationProjects() {
	_jsii_.InvokeVoid(
		o,
		"resetHasOrganizationProjects",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetHasRepositoryProjects() {
	_jsii_.InvokeVoid(
		o,
		"resetHasRepositoryProjects",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetId() {
	_jsii_.InvokeVoid(
		o,
		"resetId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetLocation() {
	_jsii_.InvokeVoid(
		o,
		"resetLocation",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetMembersCanCreateInternalRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetMembersCanCreateInternalRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetMembersCanCreatePages() {
	_jsii_.InvokeVoid(
		o,
		"resetMembersCanCreatePages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetMembersCanCreatePrivatePages() {
	_jsii_.InvokeVoid(
		o,
		"resetMembersCanCreatePrivatePages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetMembersCanCreatePrivateRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetMembersCanCreatePrivateRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetMembersCanCreatePublicPages() {
	_jsii_.InvokeVoid(
		o,
		"resetMembersCanCreatePublicPages",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetMembersCanCreatePublicRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetMembersCanCreatePublicRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetMembersCanCreateRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetMembersCanCreateRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetMembersCanForkPrivateRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetMembersCanForkPrivateRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetName() {
	_jsii_.InvokeVoid(
		o,
		"resetName",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		o,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetSecretScanningEnabledForNewRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetSecretScanningEnabledForNewRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetSecretScanningPushProtectionEnabledForNewRepositories() {
	_jsii_.InvokeVoid(
		o,
		"resetSecretScanningPushProtectionEnabledForNewRepositories",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetTwitterUsername() {
	_jsii_.InvokeVoid(
		o,
		"resetTwitterUsername",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) ResetWebCommitSignoffRequired() {
	_jsii_.InvokeVoid(
		o,
		"resetWebCommitSignoffRequired",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		o,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

