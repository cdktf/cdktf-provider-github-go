package datagithubipranges

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v6/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v6/datagithubipranges/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/github/d/ip_ranges github_ip_ranges}.
type DataGithubIpRanges interface {
	cdktf.TerraformDataSource
	Actions() *[]*string
	ActionsIpv4() *[]*string
	ActionsIpv6() *[]*string
	Api() *[]*string
	ApiIpv4() *[]*string
	ApiIpv6() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Dependabot() *[]*string
	DependabotIpv4() *[]*string
	DependabotIpv6() *[]*string
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
	Git() *[]*string
	GitIpv4() *[]*string
	GitIpv6() *[]*string
	Hooks() *[]*string
	HooksIpv4() *[]*string
	HooksIpv6() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	Importer() *[]*string
	ImporterIpv4() *[]*string
	ImporterIpv6() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	Pages() *[]*string
	PagesIpv4() *[]*string
	PagesIpv6() *[]*string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Web() *[]*string
	WebIpv4() *[]*string
	WebIpv6() *[]*string
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
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataGithubIpRanges
type jsiiProxy_DataGithubIpRanges struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataGithubIpRanges) Actions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) ActionsIpv4() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) ActionsIpv6() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"actionsIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Api() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"api",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) ApiIpv4() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) ApiIpv6() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apiIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Dependabot() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependabot",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) DependabotIpv4() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependabotIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) DependabotIpv6() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependabotIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Git() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"git",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) GitIpv4() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gitIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) GitIpv6() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"gitIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Hooks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hooks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) HooksIpv4() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hooksIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) HooksIpv6() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"hooksIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Importer() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) ImporterIpv4() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importerIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) ImporterIpv6() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"importerIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Pages() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) PagesIpv4() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pagesIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) PagesIpv6() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pagesIpv6",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) Web() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"web",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) WebIpv4() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webIpv4",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubIpRanges) WebIpv6() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"webIpv6",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/github/d/ip_ranges github_ip_ranges} Data Source.
func NewDataGithubIpRanges(scope constructs.Construct, id *string, config *DataGithubIpRangesConfig) DataGithubIpRanges {
	_init_.Initialize()

	if err := validateNewDataGithubIpRangesParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGithubIpRanges{}

	_jsii_.Create(
		"@cdktf/provider-github.dataGithubIpRanges.DataGithubIpRanges",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/github/d/ip_ranges github_ip_ranges} Data Source.
func NewDataGithubIpRanges_Override(d DataGithubIpRanges, scope constructs.Construct, id *string, config *DataGithubIpRangesConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.dataGithubIpRanges.DataGithubIpRanges",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataGithubIpRanges)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataGithubIpRanges)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataGithubIpRanges)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataGithubIpRanges)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataGithubIpRanges)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataGithubIpRanges)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
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
func DataGithubIpRanges_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGithubIpRanges_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.dataGithubIpRanges.DataGithubIpRanges",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGithubIpRanges_IsTerraformDataSource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGithubIpRanges_IsTerraformDataSourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.dataGithubIpRanges.DataGithubIpRanges",
		"isTerraformDataSource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DataGithubIpRanges_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDataGithubIpRanges_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.dataGithubIpRanges.DataGithubIpRanges",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataGithubIpRanges_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.dataGithubIpRanges.DataGithubIpRanges",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataGithubIpRanges) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataGithubIpRanges) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGithubIpRanges) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGithubIpRanges) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGithubIpRanges) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGithubIpRanges) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGithubIpRanges) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGithubIpRanges) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGithubIpRanges) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGithubIpRanges) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGithubIpRanges) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGithubIpRanges) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataGithubIpRanges) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubIpRanges) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataGithubIpRanges) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubIpRanges) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubIpRanges) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubIpRanges) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

