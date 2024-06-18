// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterpriseactionspermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v14/enterpriseactionspermissions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/enterprise_actions_permissions github_enterprise_actions_permissions}.
type EnterpriseActionsPermissions interface {
	cdktf.TerraformResource
	AllowedActions() *string
	SetAllowedActions(val *string)
	AllowedActionsConfig() EnterpriseActionsPermissionsAllowedActionsConfigOutputReference
	AllowedActionsConfigInput() *EnterpriseActionsPermissionsAllowedActionsConfig
	AllowedActionsInput() *string
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
	EnabledOrganizations() *string
	SetEnabledOrganizations(val *string)
	EnabledOrganizationsConfig() EnterpriseActionsPermissionsEnabledOrganizationsConfigOutputReference
	EnabledOrganizationsConfigInput() *EnterpriseActionsPermissionsEnabledOrganizationsConfig
	EnabledOrganizationsInput() *string
	EnterpriseId() *string
	SetEnterpriseId(val *string)
	EnterpriseIdInput() *string
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
	PutAllowedActionsConfig(value *EnterpriseActionsPermissionsAllowedActionsConfig)
	PutEnabledOrganizationsConfig(value *EnterpriseActionsPermissionsEnabledOrganizationsConfig)
	ResetAllowedActions()
	ResetAllowedActionsConfig()
	ResetEnabledOrganizationsConfig()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
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

// The jsii proxy struct for EnterpriseActionsPermissions
type jsiiProxy_EnterpriseActionsPermissions struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EnterpriseActionsPermissions) AllowedActions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) AllowedActionsConfig() EnterpriseActionsPermissionsAllowedActionsConfigOutputReference {
	var returns EnterpriseActionsPermissionsAllowedActionsConfigOutputReference
	_jsii_.Get(
		j,
		"allowedActionsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) AllowedActionsConfigInput() *EnterpriseActionsPermissionsAllowedActionsConfig {
	var returns *EnterpriseActionsPermissionsAllowedActionsConfig
	_jsii_.Get(
		j,
		"allowedActionsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) AllowedActionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) EnabledOrganizations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledOrganizations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) EnabledOrganizationsConfig() EnterpriseActionsPermissionsEnabledOrganizationsConfigOutputReference {
	var returns EnterpriseActionsPermissionsEnabledOrganizationsConfigOutputReference
	_jsii_.Get(
		j,
		"enabledOrganizationsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) EnabledOrganizationsConfigInput() *EnterpriseActionsPermissionsEnabledOrganizationsConfig {
	var returns *EnterpriseActionsPermissionsEnabledOrganizationsConfig
	_jsii_.Get(
		j,
		"enabledOrganizationsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) EnabledOrganizationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledOrganizationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) EnterpriseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) EnterpriseIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/enterprise_actions_permissions github_enterprise_actions_permissions} Resource.
func NewEnterpriseActionsPermissions(scope constructs.Construct, id *string, config *EnterpriseActionsPermissionsConfig) EnterpriseActionsPermissions {
	_init_.Initialize()

	if err := validateNewEnterpriseActionsPermissionsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseActionsPermissions{}

	_jsii_.Create(
		"@cdktf/provider-github.enterpriseActionsPermissions.EnterpriseActionsPermissions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/enterprise_actions_permissions github_enterprise_actions_permissions} Resource.
func NewEnterpriseActionsPermissions_Override(e EnterpriseActionsPermissions, scope constructs.Construct, id *string, config *EnterpriseActionsPermissionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.enterpriseActionsPermissions.EnterpriseActionsPermissions",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetAllowedActions(val *string) {
	if err := j.validateSetAllowedActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedActions",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetEnabledOrganizations(val *string) {
	if err := j.validateSetEnabledOrganizationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledOrganizations",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetEnterpriseId(val *string) {
	if err := j.validateSetEnterpriseIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterpriseId",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissions)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

// Generates CDKTF code for importing a EnterpriseActionsPermissions resource upon running "cdktf plan <stack-name>".
func EnterpriseActionsPermissions_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEnterpriseActionsPermissions_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseActionsPermissions.EnterpriseActionsPermissions",
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
func EnterpriseActionsPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseActionsPermissions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseActionsPermissions.EnterpriseActionsPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseActionsPermissions_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseActionsPermissions_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseActionsPermissions.EnterpriseActionsPermissions",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseActionsPermissions_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseActionsPermissions_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseActionsPermissions.EnterpriseActionsPermissions",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EnterpriseActionsPermissions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.enterpriseActionsPermissions.EnterpriseActionsPermissions",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissions) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseActionsPermissions) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) PutAllowedActionsConfig(value *EnterpriseActionsPermissionsAllowedActionsConfig) {
	if err := e.validatePutAllowedActionsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putAllowedActionsConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) PutEnabledOrganizationsConfig(value *EnterpriseActionsPermissionsEnabledOrganizationsConfig) {
	if err := e.validatePutEnabledOrganizationsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putEnabledOrganizationsConfig",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ResetAllowedActions() {
	_jsii_.InvokeVoid(
		e,
		"resetAllowedActions",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ResetAllowedActionsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetAllowedActionsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ResetEnabledOrganizationsConfig() {
	_jsii_.InvokeVoid(
		e,
		"resetEnabledOrganizationsConfig",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissions) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissions) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

