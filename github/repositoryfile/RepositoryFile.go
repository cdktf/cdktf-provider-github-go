package repositoryfile

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v7/repositoryfile/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/github/r/repository_file github_repository_file}.
type RepositoryFile interface {
	cdktf.TerraformResource
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CommitAuthor() *string
	SetCommitAuthor(val *string)
	CommitAuthorInput() *string
	CommitEmail() *string
	SetCommitEmail(val *string)
	CommitEmailInput() *string
	CommitMessage() *string
	SetCommitMessage(val *string)
	CommitMessageInput() *string
	CommitSha() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Content() *string
	SetContent(val *string)
	ContentInput() *string
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	File() *string
	SetFile(val *string)
	FileInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	// The tree node.
	Node() constructs.Node
	OverwriteOnCreate() interface{}
	SetOverwriteOnCreate(val interface{})
	OverwriteOnCreateInput() interface{}
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
	Ref() *string
	Repository() *string
	SetRepository(val *string)
	RepositoryInput() *string
	Sha() *string
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
	ResetBranch()
	ResetCommitAuthor()
	ResetCommitEmail()
	ResetCommitMessage()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOverwriteOnCreate()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for RepositoryFile
type jsiiProxy_RepositoryFile struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_RepositoryFile) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitAuthor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitAuthor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitAuthorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitAuthorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) CommitSha() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commitSha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Content() *string {
	var returns *string
	_jsii_.Get(
		j,
		"content",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"contentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) File() *string {
	var returns *string
	_jsii_.Get(
		j,
		"file",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) FileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) OverwriteOnCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteOnCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) OverwriteOnCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overwriteOnCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Ref() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ref",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) Sha() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryFile) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/github/r/repository_file github_repository_file} Resource.
func NewRepositoryFile(scope constructs.Construct, id *string, config *RepositoryFileConfig) RepositoryFile {
	_init_.Initialize()

	if err := validateNewRepositoryFileParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositoryFile{}

	_jsii_.Create(
		"@cdktf/provider-github.repositoryFile.RepositoryFile",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/github/r/repository_file github_repository_file} Resource.
func NewRepositoryFile_Override(r RepositoryFile, scope constructs.Construct, id *string, config *RepositoryFileConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.repositoryFile.RepositoryFile",
		[]interface{}{scope, id, config},
		r,
	)
}

func (j *jsiiProxy_RepositoryFile)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetCommitAuthor(val *string) {
	if err := j.validateSetCommitAuthorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitAuthor",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetCommitEmail(val *string) {
	if err := j.validateSetCommitEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitEmail",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetCommitMessage(val *string) {
	if err := j.validateSetCommitMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"commitMessage",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetContent(val *string) {
	if err := j.validateSetContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"content",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetFile(val *string) {
	if err := j.validateSetFileParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"file",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetOverwriteOnCreate(val interface{}) {
	if err := j.validateSetOverwriteOnCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overwriteOnCreate",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_RepositoryFile)SetRepository(val *string) {
	if err := j.validateSetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repository",
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
func RepositoryFile_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRepositoryFile_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.repositoryFile.RepositoryFile",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RepositoryFile_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRepositoryFile_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.repositoryFile.RepositoryFile",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func RepositoryFile_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateRepositoryFile_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.repositoryFile.RepositoryFile",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func RepositoryFile_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.repositoryFile.RepositoryFile",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (r *jsiiProxy_RepositoryFile) AddOverride(path *string, value interface{}) {
	if err := r.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (r *jsiiProxy_RepositoryFile) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) OverrideLogicalId(newLogicalId *string) {
	if err := r.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (r *jsiiProxy_RepositoryFile) ResetBranch() {
	_jsii_.InvokeVoid(
		r,
		"resetBranch",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetCommitAuthor() {
	_jsii_.InvokeVoid(
		r,
		"resetCommitAuthor",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetCommitEmail() {
	_jsii_.InvokeVoid(
		r,
		"resetCommitEmail",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetCommitMessage() {
	_jsii_.InvokeVoid(
		r,
		"resetCommitMessage",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetId() {
	_jsii_.InvokeVoid(
		r,
		"resetId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		r,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) ResetOverwriteOnCreate() {
	_jsii_.InvokeVoid(
		r,
		"resetOverwriteOnCreate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryFile) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryFile) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

