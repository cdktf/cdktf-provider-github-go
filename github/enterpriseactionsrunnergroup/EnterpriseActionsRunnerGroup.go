// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterpriseactionsrunnergroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v14/enterpriseactionsrunnergroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/enterprise_actions_runner_group github_enterprise_actions_runner_group}.
type EnterpriseActionsRunnerGroup interface {
	cdktf.TerraformResource
	AllowsPublicRepositories() interface{}
	SetAllowsPublicRepositories(val interface{})
	AllowsPublicRepositoriesInput() interface{}
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
	Default() cdktf.IResolvable
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnterpriseSlug() *string
	SetEnterpriseSlug(val *string)
	EnterpriseSlugInput() *string
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
	RestrictedToWorkflows() interface{}
	SetRestrictedToWorkflows(val interface{})
	RestrictedToWorkflowsInput() interface{}
	RunnersUrl() *string
	SelectedOrganizationIds() *[]*float64
	SetSelectedOrganizationIds(val *[]*float64)
	SelectedOrganizationIdsInput() *[]*float64
	SelectedOrganizationsUrl() *string
	SelectedWorkflows() *[]*string
	SetSelectedWorkflows(val *[]*string)
	SelectedWorkflowsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Visibility() *string
	SetVisibility(val *string)
	VisibilityInput() *string
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
	ResetAllowsPublicRepositories()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRestrictedToWorkflows()
	ResetSelectedOrganizationIds()
	ResetSelectedWorkflows()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for EnterpriseActionsRunnerGroup
type jsiiProxy_EnterpriseActionsRunnerGroup struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) AllowsPublicRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowsPublicRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) AllowsPublicRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowsPublicRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Default() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"default",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) EnterpriseSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) EnterpriseSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Etag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"etag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) RestrictedToWorkflows() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictedToWorkflows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) RestrictedToWorkflowsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictedToWorkflowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) RunnersUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runnersUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) SelectedOrganizationIds() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"selectedOrganizationIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) SelectedOrganizationIdsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"selectedOrganizationIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) SelectedOrganizationsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"selectedOrganizationsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) SelectedWorkflows() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selectedWorkflows",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) SelectedWorkflowsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"selectedWorkflowsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) Visibility() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibility",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup) VisibilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"visibilityInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/enterprise_actions_runner_group github_enterprise_actions_runner_group} Resource.
func NewEnterpriseActionsRunnerGroup(scope constructs.Construct, id *string, config *EnterpriseActionsRunnerGroupConfig) EnterpriseActionsRunnerGroup {
	_init_.Initialize()

	if err := validateNewEnterpriseActionsRunnerGroupParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseActionsRunnerGroup{}

	_jsii_.Create(
		"@cdktf/provider-github.enterpriseActionsRunnerGroup.EnterpriseActionsRunnerGroup",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/enterprise_actions_runner_group github_enterprise_actions_runner_group} Resource.
func NewEnterpriseActionsRunnerGroup_Override(e EnterpriseActionsRunnerGroup, scope constructs.Construct, id *string, config *EnterpriseActionsRunnerGroupConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.enterpriseActionsRunnerGroup.EnterpriseActionsRunnerGroup",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetAllowsPublicRepositories(val interface{}) {
	if err := j.validateSetAllowsPublicRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowsPublicRepositories",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetEnterpriseSlug(val *string) {
	if err := j.validateSetEnterpriseSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterpriseSlug",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetRestrictedToWorkflows(val interface{}) {
	if err := j.validateSetRestrictedToWorkflowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictedToWorkflows",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetSelectedOrganizationIds(val *[]*float64) {
	if err := j.validateSetSelectedOrganizationIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectedOrganizationIds",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetSelectedWorkflows(val *[]*string) {
	if err := j.validateSetSelectedWorkflowsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"selectedWorkflows",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsRunnerGroup)SetVisibility(val *string) {
	if err := j.validateSetVisibilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"visibility",
		val,
	)
}

// Generates CDKTF code for importing a EnterpriseActionsRunnerGroup resource upon running "cdktf plan <stack-name>".
func EnterpriseActionsRunnerGroup_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEnterpriseActionsRunnerGroup_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseActionsRunnerGroup.EnterpriseActionsRunnerGroup",
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
func EnterpriseActionsRunnerGroup_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseActionsRunnerGroup_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseActionsRunnerGroup.EnterpriseActionsRunnerGroup",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseActionsRunnerGroup_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseActionsRunnerGroup_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseActionsRunnerGroup.EnterpriseActionsRunnerGroup",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseActionsRunnerGroup_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseActionsRunnerGroup_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseActionsRunnerGroup.EnterpriseActionsRunnerGroup",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EnterpriseActionsRunnerGroup_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.enterpriseActionsRunnerGroup.EnterpriseActionsRunnerGroup",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ResetAllowsPublicRepositories() {
	_jsii_.InvokeVoid(
		e,
		"resetAllowsPublicRepositories",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ResetRestrictedToWorkflows() {
	_jsii_.InvokeVoid(
		e,
		"resetRestrictedToWorkflows",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ResetSelectedOrganizationIds() {
	_jsii_.InvokeVoid(
		e,
		"resetSelectedOrganizationIds",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ResetSelectedWorkflows() {
	_jsii_.InvokeVoid(
		e,
		"resetSelectedWorkflows",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsRunnerGroup) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

