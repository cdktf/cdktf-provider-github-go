// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotectionv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v12/branchprotectionv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/branch_protection_v3 github_branch_protection_v3}.
type BranchProtectionV3 interface {
	cdktf.TerraformResource
	Branch() *string
	SetBranch(val *string)
	BranchInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnforceAdmins() interface{}
	SetEnforceAdmins(val interface{})
	EnforceAdminsInput() interface{}
	Etag() *string
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
	Repository() *string
	SetRepository(val *string)
	RepositoryInput() *string
	RequireConversationResolution() interface{}
	SetRequireConversationResolution(val interface{})
	RequireConversationResolutionInput() interface{}
	RequiredPullRequestReviews() BranchProtectionV3RequiredPullRequestReviewsOutputReference
	RequiredPullRequestReviewsInput() *BranchProtectionV3RequiredPullRequestReviews
	RequiredStatusChecks() BranchProtectionV3RequiredStatusChecksOutputReference
	RequiredStatusChecksInput() *BranchProtectionV3RequiredStatusChecks
	RequireSignedCommits() interface{}
	SetRequireSignedCommits(val interface{})
	RequireSignedCommitsInput() interface{}
	Restrictions() BranchProtectionV3RestrictionsOutputReference
	RestrictionsInput() *BranchProtectionV3Restrictions
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutRequiredPullRequestReviews(value *BranchProtectionV3RequiredPullRequestReviews)
	PutRequiredStatusChecks(value *BranchProtectionV3RequiredStatusChecks)
	PutRestrictions(value *BranchProtectionV3Restrictions)
	ResetEnforceAdmins()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRequireConversationResolution()
	ResetRequiredPullRequestReviews()
	ResetRequiredStatusChecks()
	ResetRequireSignedCommits()
	ResetRestrictions()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for BranchProtectionV3
type jsiiProxy_BranchProtectionV3 struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_BranchProtectionV3) Branch() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) BranchInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"branchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) EnforceAdmins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) EnforceAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enforceAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Repository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RepositoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"repositoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RequireConversationResolution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireConversationResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RequireConversationResolutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireConversationResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RequiredPullRequestReviews() BranchProtectionV3RequiredPullRequestReviewsOutputReference {
	var returns BranchProtectionV3RequiredPullRequestReviewsOutputReference
	_jsii_.Get(
		j,
		"requiredPullRequestReviews",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RequiredPullRequestReviewsInput() *BranchProtectionV3RequiredPullRequestReviews {
	var returns *BranchProtectionV3RequiredPullRequestReviews
	_jsii_.Get(
		j,
		"requiredPullRequestReviewsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RequiredStatusChecks() BranchProtectionV3RequiredStatusChecksOutputReference {
	var returns BranchProtectionV3RequiredStatusChecksOutputReference
	_jsii_.Get(
		j,
		"requiredStatusChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RequiredStatusChecksInput() *BranchProtectionV3RequiredStatusChecks {
	var returns *BranchProtectionV3RequiredStatusChecks
	_jsii_.Get(
		j,
		"requiredStatusChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RequireSignedCommits() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSignedCommits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RequireSignedCommitsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireSignedCommitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) Restrictions() BranchProtectionV3RestrictionsOutputReference {
	var returns BranchProtectionV3RestrictionsOutputReference
	_jsii_.Get(
		j,
		"restrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) RestrictionsInput() *BranchProtectionV3Restrictions {
	var returns *BranchProtectionV3Restrictions
	_jsii_.Get(
		j,
		"restrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/branch_protection_v3 github_branch_protection_v3} Resource.
func NewBranchProtectionV3(scope constructs.Construct, id *string, config *BranchProtectionV3Config) BranchProtectionV3 {
	_init_.Initialize()

	if err := validateNewBranchProtectionV3Parameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_BranchProtectionV3{}

	_jsii_.Create(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.43.0/docs/resources/branch_protection_v3 github_branch_protection_v3} Resource.
func NewBranchProtectionV3_Override(b BranchProtectionV3, scope constructs.Construct, id *string, config *BranchProtectionV3Config) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3",
		[]interface{}{scope, id, config},
		b,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetBranch(val *string) {
	if err := j.validateSetBranchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"branch",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetEnforceAdmins(val interface{}) {
	if err := j.validateSetEnforceAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enforceAdmins",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetRepository(val *string) {
	if err := j.validateSetRepositoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"repository",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetRequireConversationResolution(val interface{}) {
	if err := j.validateSetRequireConversationResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireConversationResolution",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3)SetRequireSignedCommits(val interface{}) {
	if err := j.validateSetRequireSignedCommitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireSignedCommits",
		val,
	)
}

// Generates CDKTF code for importing a BranchProtectionV3 resource upon running "cdktf plan <stack-name>".
func BranchProtectionV3_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateBranchProtectionV3_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func BranchProtectionV3_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBranchProtectionV3_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BranchProtectionV3_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBranchProtectionV3_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func BranchProtectionV3_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateBranchProtectionV3_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func BranchProtectionV3_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (b *jsiiProxy_BranchProtectionV3) AddMoveTarget(moveTarget *string) {
	if err := b.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (b *jsiiProxy_BranchProtectionV3) AddOverride(path *string, value interface{}) {
	if err := b.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (b *jsiiProxy_BranchProtectionV3) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BranchProtectionV3) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BranchProtectionV3) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BranchProtectionV3) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BranchProtectionV3) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BranchProtectionV3) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BranchProtectionV3) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BranchProtectionV3) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BranchProtectionV3) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BranchProtectionV3) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := b.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (b *jsiiProxy_BranchProtectionV3) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BranchProtectionV3) MoveFromId(id *string) {
	if err := b.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveFromId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BranchProtectionV3) MoveTo(moveTarget *string, index interface{}) {
	if err := b.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (b *jsiiProxy_BranchProtectionV3) MoveToId(id *string) {
	if err := b.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"moveToId",
		[]interface{}{id},
	)
}

func (b *jsiiProxy_BranchProtectionV3) OverrideLogicalId(newLogicalId *string) {
	if err := b.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (b *jsiiProxy_BranchProtectionV3) PutRequiredPullRequestReviews(value *BranchProtectionV3RequiredPullRequestReviews) {
	if err := b.validatePutRequiredPullRequestReviewsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRequiredPullRequestReviews",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BranchProtectionV3) PutRequiredStatusChecks(value *BranchProtectionV3RequiredStatusChecks) {
	if err := b.validatePutRequiredStatusChecksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRequiredStatusChecks",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BranchProtectionV3) PutRestrictions(value *BranchProtectionV3Restrictions) {
	if err := b.validatePutRestrictionsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRestrictions",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BranchProtectionV3) ResetEnforceAdmins() {
	_jsii_.InvokeVoid(
		b,
		"resetEnforceAdmins",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3) ResetId() {
	_jsii_.InvokeVoid(
		b,
		"resetId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		b,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3) ResetRequireConversationResolution() {
	_jsii_.InvokeVoid(
		b,
		"resetRequireConversationResolution",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3) ResetRequiredPullRequestReviews() {
	_jsii_.InvokeVoid(
		b,
		"resetRequiredPullRequestReviews",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3) ResetRequiredStatusChecks() {
	_jsii_.InvokeVoid(
		b,
		"resetRequiredStatusChecks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3) ResetRequireSignedCommits() {
	_jsii_.InvokeVoid(
		b,
		"resetRequireSignedCommits",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3) ResetRestrictions() {
	_jsii_.InvokeVoid(
		b,
		"resetRestrictions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

