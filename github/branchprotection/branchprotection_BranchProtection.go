package branchprotection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v5/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v5/branchprotection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/github/r/branch_protection github_branch_protection}.
type BranchProtection interface {
	cdktf.TerraformResource
	AllowsDeletions() interface{}
	SetAllowsDeletions(val interface{})
	AllowsDeletionsInput() interface{}
	AllowsForcePushes() interface{}
	SetAllowsForcePushes(val interface{})
	AllowsForcePushesInput() interface{}
	BlocksCreations() interface{}
	SetBlocksCreations(val interface{})
	BlocksCreationsInput() interface{}
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnforceAdmins() interface{}
	SetEnforceAdmins(val interface{})
	EnforceAdminsInput() interface{}
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
	LockBranch() interface{}
	SetLockBranch(val interface{})
	LockBranchInput() interface{}
	// The tree node.
	Node() constructs.Node
	Pattern() *string
	SetPattern(val *string)
	PatternInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushRestrictions() *[]*string
	SetPushRestrictions(val *[]*string)
	PushRestrictionsInput() *[]*string
	// Experimental.
	RawOverrides() interface{}
	RepositoryId() *string
	SetRepositoryId(val *string)
	RepositoryIdInput() *string
	RequireConversationResolution() interface{}
	SetRequireConversationResolution(val interface{})
	RequireConversationResolutionInput() interface{}
	RequiredLinearHistory() interface{}
	SetRequiredLinearHistory(val interface{})
	RequiredLinearHistoryInput() interface{}
	RequiredPullRequestReviews() BranchProtectionRequiredPullRequestReviewsList
	RequiredPullRequestReviewsInput() interface{}
	RequiredStatusChecks() BranchProtectionRequiredStatusChecksList
	RequiredStatusChecksInput() interface{}
	RequireSignedCommits() interface{}
	SetRequireSignedCommits(val interface{})
	RequireSignedCommitsInput() interface{}
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
	PutRequiredPullRequestReviews(value interface{})
	PutRequiredStatusChecks(value interface{})
	ResetAllowsDeletions()
	ResetAllowsForcePushes()
	ResetBlocksCreations()
	ResetEnforceAdmins()
	ResetId()
	ResetLockBranch()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPushRestrictions()
	ResetRequireConversationResolution()
	ResetRequiredLinearHistory()
	ResetRequiredPullRequestReviews()
	ResetRequiredStatusChecks()
	ResetRequireSignedCommits()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BranchProtection
type jsiiProxy_BranchProtection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BranchProtection) AllowsDeletions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowsDeletions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowsDeletionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowsDeletionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowsForcePushes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowsForcePushes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) AllowsForcePushesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowsForcePushesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) BlocksCreations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blocksCreations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) BlocksCreationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"blocksCreationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) EnforceAdmins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) EnforceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) LockBranch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockBranch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) LockBranchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"lockBranchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Pattern() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) PatternInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) PushRestrictions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pushRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) PushRestrictionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pushRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RepositoryId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RepositoryIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequireConversationResolution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireConversationResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequireConversationResolutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireConversationResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequiredLinearHistory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredLinearHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequiredLinearHistoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredLinearHistoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequiredPullRequestReviews() BranchProtectionRequiredPullRequestReviewsList {
	var returns BranchProtectionRequiredPullRequestReviewsList
	_jsii_.Get(
		j,
		"requiredPullRequestReviews",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequiredPullRequestReviewsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredPullRequestReviewsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequiredStatusChecks() BranchProtectionRequiredStatusChecksList {
	var returns BranchProtectionRequiredStatusChecksList
	_jsii_.Get(
		j,
		"requiredStatusChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequiredStatusChecksInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredStatusChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequireSignedCommits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSignedCommits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) RequireSignedCommitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSignedCommitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/github/r/branch_protection github_branch_protection} Resource.
func NewBranchProtection(scope constructs.Construct, id *string, config *BranchProtectionConfig) BranchProtection {
	_init_.Initialize()

	if err := validateNewBranchProtectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BranchProtection{}

	_jsii_.Create(
		"@cdktf/provider-github.branchProtection.BranchProtection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/github/r/branch_protection github_branch_protection} Resource.
func NewBranchProtection_Override(b BranchProtection, scope constructs.Construct, id *string, config *BranchProtectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.branchProtection.BranchProtection",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BranchProtection)SetAllowsDeletions(val interface{}) {
	if err := j.validateSetAllowsDeletionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowsDeletions",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetAllowsForcePushes(val interface{}) {
	if err := j.validateSetAllowsForcePushesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowsForcePushes",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetBlocksCreations(val interface{}) {
	if err := j.validateSetBlocksCreationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"blocksCreations",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetEnforceAdmins(val interface{}) {
	if err := j.validateSetEnforceAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceAdmins",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetLockBranch(val interface{}) {
	if err := j.validateSetLockBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lockBranch",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetPattern(val *string) {
	if err := j.validateSetPatternParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pattern",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetPushRestrictions(val *[]*string) {
	if err := j.validateSetPushRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushRestrictions",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetRepositoryId(val *string) {
	if err := j.validateSetRepositoryIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repositoryId",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetRequireConversationResolution(val interface{}) {
	if err := j.validateSetRequireConversationResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireConversationResolution",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetRequiredLinearHistory(val interface{}) {
	if err := j.validateSetRequiredLinearHistoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredLinearHistory",
		val,
	)
}

func (j *jsiiProxy_BranchProtection)SetRequireSignedCommits(val interface{}) {
	if err := j.validateSetRequireSignedCommitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSignedCommits",
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
func BranchProtection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBranchProtection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.branchProtection.BranchProtection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BranchProtection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBranchProtection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.branchProtection.BranchProtection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BranchProtection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBranchProtection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.branchProtection.BranchProtection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BranchProtection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.branchProtection.BranchProtection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BranchProtection) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BranchProtection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BranchProtection) PutRequiredPullRequestReviews(value interface{}) {
	if err := b.validatePutRequiredPullRequestReviewsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRequiredPullRequestReviews",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BranchProtection) PutRequiredStatusChecks(value interface{}) {
	if err := b.validatePutRequiredStatusChecksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRequiredStatusChecks",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BranchProtection) ResetAllowsDeletions() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowsDeletions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetAllowsForcePushes() {
	_jsii_.InvokeVoid(
		b,
		"resetAllowsForcePushes",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetBlocksCreations() {
	_jsii_.InvokeVoid(
		b,
		"resetBlocksCreations",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetEnforceAdmins() {
	_jsii_.InvokeVoid(
		b,
		"resetEnforceAdmins",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetLockBranch() {
	_jsii_.InvokeVoid(
		b,
		"resetLockBranch",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetPushRestrictions() {
	_jsii_.InvokeVoid(
		b,
		"resetPushRestrictions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetRequireConversationResolution() {
	_jsii_.InvokeVoid(
		b,
		"resetRequireConversationResolution",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetRequiredLinearHistory() {
	_jsii_.InvokeVoid(
		b,
		"resetRequiredLinearHistory",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetRequiredPullRequestReviews() {
	_jsii_.InvokeVoid(
		b,
		"resetRequiredPullRequestReviews",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetRequiredStatusChecks() {
	_jsii_.InvokeVoid(
		b,
		"resetRequiredStatusChecks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) ResetRequireSignedCommits() {
	_jsii_.InvokeVoid(
		b,
		"resetRequireSignedCommits",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

