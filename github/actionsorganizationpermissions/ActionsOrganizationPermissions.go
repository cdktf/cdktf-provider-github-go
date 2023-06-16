package actionsorganizationpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v10/actionsorganizationpermissions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/resources/actions_organization_permissions github_actions_organization_permissions}.
type ActionsOrganizationPermissions interface {
	cdktf.TerraformResource
	AllowedActions() *string
	SetAllowedActions(val *string)
	AllowedActionsConfig() ActionsOrganizationPermissionsAllowedActionsConfigOutputReference
	AllowedActionsConfigInput() *ActionsOrganizationPermissionsAllowedActionsConfig
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
	EnabledRepositories() *string
	SetEnabledRepositories(val *string)
	EnabledRepositoriesConfig() ActionsOrganizationPermissionsEnabledRepositoriesConfigOutputReference
	EnabledRepositoriesConfigInput() *ActionsOrganizationPermissionsEnabledRepositoriesConfig
	EnabledRepositoriesInput() *string
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
	PutAllowedActionsConfig(value *ActionsOrganizationPermissionsAllowedActionsConfig)
	PutEnabledRepositoriesConfig(value *ActionsOrganizationPermissionsEnabledRepositoriesConfig)
	ResetAllowedActions()
	ResetAllowedActionsConfig()
	ResetEnabledRepositoriesConfig()
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

// The jsii proxy struct for ActionsOrganizationPermissions
type jsiiProxy_ActionsOrganizationPermissions struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ActionsOrganizationPermissions) AllowedActions() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedActions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) AllowedActionsConfig() ActionsOrganizationPermissionsAllowedActionsConfigOutputReference {
	var returns ActionsOrganizationPermissionsAllowedActionsConfigOutputReference
	_jsii_.Get(
		j,
		"allowedActionsConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) AllowedActionsConfigInput() *ActionsOrganizationPermissionsAllowedActionsConfig {
	var returns *ActionsOrganizationPermissionsAllowedActionsConfig
	_jsii_.Get(
		j,
		"allowedActionsConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) AllowedActionsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"allowedActionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) EnabledRepositories() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) EnabledRepositoriesConfig() ActionsOrganizationPermissionsEnabledRepositoriesConfigOutputReference {
	var returns ActionsOrganizationPermissionsEnabledRepositoriesConfigOutputReference
	_jsii_.Get(
		j,
		"enabledRepositoriesConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) EnabledRepositoriesConfigInput() *ActionsOrganizationPermissionsEnabledRepositoriesConfig {
	var returns *ActionsOrganizationPermissionsEnabledRepositoriesConfig
	_jsii_.Get(
		j,
		"enabledRepositoriesConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) EnabledRepositoriesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enabledRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissions) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/resources/actions_organization_permissions github_actions_organization_permissions} Resource.
func NewActionsOrganizationPermissions(scope constructs.Construct, id *string, config *ActionsOrganizationPermissionsConfig) ActionsOrganizationPermissions {
	_init_.Initialize()

	if err := validateNewActionsOrganizationPermissionsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActionsOrganizationPermissions{}

	_jsii_.Create(
		"@cdktf/provider-github.actionsOrganizationPermissions.ActionsOrganizationPermissions",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.27.0/docs/resources/actions_organization_permissions github_actions_organization_permissions} Resource.
func NewActionsOrganizationPermissions_Override(a ActionsOrganizationPermissions, scope constructs.Construct, id *string, config *ActionsOrganizationPermissionsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.actionsOrganizationPermissions.ActionsOrganizationPermissions",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetAllowedActions(val *string) {
	if err := j.validateSetAllowedActionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedActions",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetEnabledRepositories(val *string) {
	if err := j.validateSetEnabledRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabledRepositories",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissions)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
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
func ActionsOrganizationPermissions_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActionsOrganizationPermissions_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.actionsOrganizationPermissions.ActionsOrganizationPermissions",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ActionsOrganizationPermissions_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActionsOrganizationPermissions_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.actionsOrganizationPermissions.ActionsOrganizationPermissions",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func ActionsOrganizationPermissions_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateActionsOrganizationPermissions_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.actionsOrganizationPermissions.ActionsOrganizationPermissions",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ActionsOrganizationPermissions_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.actionsOrganizationPermissions.ActionsOrganizationPermissions",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissions) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissions) PutAllowedActionsConfig(value *ActionsOrganizationPermissionsAllowedActionsConfig) {
	if err := a.validatePutAllowedActionsConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putAllowedActionsConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissions) PutEnabledRepositoriesConfig(value *ActionsOrganizationPermissionsEnabledRepositoriesConfig) {
	if err := a.validatePutEnabledRepositoriesConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putEnabledRepositoriesConfig",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissions) ResetAllowedActions() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedActions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissions) ResetAllowedActionsConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedActionsConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissions) ResetEnabledRepositoriesConfig() {
	_jsii_.InvokeVoid(
		a,
		"resetEnabledRepositoriesConfig",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissions) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissions) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissions) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissions) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

