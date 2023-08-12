package datagithubrepositorypullrequests

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v10/datagithubrepositorypullrequests/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/data-sources/repository_pull_requests github_repository_pull_requests}.
type DataGithubRepositoryPullRequests interface {
	cdktf.TerraformDataSource
	BaseRef() *string
	SetBaseRef(val *string)
	BaseRefInput() *string
	BaseRepository() *string
	SetBaseRepository(val *string)
	BaseRepositoryInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HeadRef() *string
	SetHeadRef(val *string)
	HeadRefInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Results() DataGithubRepositoryPullRequestsResultsList
	SortBy() *string
	SetSortBy(val *string)
	SortByInput() *string
	SortDirection() *string
	SetSortDirection(val *string)
	SortDirectionInput() *string
	State() *string
	SetState(val *string)
	StateInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	ResetBaseRef()
	ResetHeadRef()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetSortBy()
	ResetSortDirection()
	ResetState()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGithubRepositoryPullRequests
type jsiiProxy_DataGithubRepositoryPullRequests struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) BaseRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) BaseRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) BaseRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) BaseRepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseRepositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) HeadRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) HeadRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) Results() DataGithubRepositoryPullRequestsResultsList {
	var returns DataGithubRepositoryPullRequestsResultsList
	_jsii_.Get(
		j,
		"results",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) SortBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) SortByInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortByInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) SortDirection() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortDirection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) SortDirectionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sortDirectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) StateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"stateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/data-sources/repository_pull_requests github_repository_pull_requests} Data Source.
func NewDataGithubRepositoryPullRequests(scope constructs.Construct, id *string, config *DataGithubRepositoryPullRequestsConfig) DataGithubRepositoryPullRequests {
	_init_.Initialize()

	if err := validateNewDataGithubRepositoryPullRequestsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGithubRepositoryPullRequests{}

	_jsii_.Create(
		"@cdktf/provider-github.dataGithubRepositoryPullRequests.DataGithubRepositoryPullRequests",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/data-sources/repository_pull_requests github_repository_pull_requests} Data Source.
func NewDataGithubRepositoryPullRequests_Override(d DataGithubRepositoryPullRequests, scope constructs.Construct, id *string, config *DataGithubRepositoryPullRequestsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.dataGithubRepositoryPullRequests.DataGithubRepositoryPullRequests",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetBaseRef(val *string) {
	if err := j.validateSetBaseRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseRef",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetBaseRepository(val *string) {
	if err := j.validateSetBaseRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"baseRepository",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetHeadRef(val *string) {
	if err := j.validateSetHeadRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headRef",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetOwner(val *string) {
	if err := j.validateSetOwnerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetSortBy(val *string) {
	if err := j.validateSetSortByParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortBy",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetSortDirection(val *string) {
	if err := j.validateSetSortDirectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sortDirection",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequests)SetState(val *string) {
	if err := j.validateSetStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"state",
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
func DataGithubRepositoryPullRequests_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGithubRepositoryPullRequests_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.dataGithubRepositoryPullRequests.DataGithubRepositoryPullRequests",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGithubRepositoryPullRequests_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGithubRepositoryPullRequests_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.dataGithubRepositoryPullRequests.DataGithubRepositoryPullRequests",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGithubRepositoryPullRequests_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGithubRepositoryPullRequests_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.dataGithubRepositoryPullRequests.DataGithubRepositoryPullRequests",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGithubRepositoryPullRequests_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.dataGithubRepositoryPullRequests.DataGithubRepositoryPullRequests",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ResetBaseRef() {
	_jsii_.InvokeVoid(
		d,
		"resetBaseRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ResetHeadRef() {
	_jsii_.InvokeVoid(
		d,
		"resetHeadRef",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ResetOwner() {
	_jsii_.InvokeVoid(
		d,
		"resetOwner",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ResetSortBy() {
	_jsii_.InvokeVoid(
		d,
		"resetSortBy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ResetSortDirection() {
	_jsii_.InvokeVoid(
		d,
		"resetSortDirection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ResetState() {
	_jsii_.InvokeVoid(
		d,
		"resetState",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequests) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

