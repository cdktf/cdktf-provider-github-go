// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisesecurityanalysissettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v15/enterprisesecurityanalysissettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings github_enterprise_security_analysis_settings}.
type EnterpriseSecurityAnalysisSettings interface {
	cdktf.TerraformResource
	AdvancedSecurityEnabledForNewRepositories() interface{}
	SetAdvancedSecurityEnabledForNewRepositories(val interface{})
	AdvancedSecurityEnabledForNewRepositoriesInput() interface{}
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
	EnterpriseSlug() *string
	SetEnterpriseSlug(val *string)
	EnterpriseSlugInput() *string
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
	SecretScanningEnabledForNewRepositories() interface{}
	SetSecretScanningEnabledForNewRepositories(val interface{})
	SecretScanningEnabledForNewRepositoriesInput() interface{}
	SecretScanningPushProtectionCustomLink() *string
	SetSecretScanningPushProtectionCustomLink(val *string)
	SecretScanningPushProtectionCustomLinkInput() *string
	SecretScanningPushProtectionEnabledForNewRepositories() interface{}
	SetSecretScanningPushProtectionEnabledForNewRepositories(val interface{})
	SecretScanningPushProtectionEnabledForNewRepositoriesInput() interface{}
	SecretScanningValidityChecksEnabled() interface{}
	SetSecretScanningValidityChecksEnabled(val interface{})
	SecretScanningValidityChecksEnabledInput() interface{}
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
	ResetAdvancedSecurityEnabledForNewRepositories()
	ResetId()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecretScanningEnabledForNewRepositories()
	ResetSecretScanningPushProtectionCustomLink()
	ResetSecretScanningPushProtectionEnabledForNewRepositories()
	ResetSecretScanningValidityChecksEnabled()
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

// The jsii proxy struct for EnterpriseSecurityAnalysisSettings
type jsiiProxy_EnterpriseSecurityAnalysisSettings struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) AdvancedSecurityEnabledForNewRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedSecurityEnabledForNewRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) AdvancedSecurityEnabledForNewRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"advancedSecurityEnabledForNewRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) EnterpriseSlug() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseSlug",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) EnterpriseSlugInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"enterpriseSlugInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) SecretScanningEnabledForNewRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningEnabledForNewRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) SecretScanningEnabledForNewRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningEnabledForNewRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) SecretScanningPushProtectionCustomLink() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretScanningPushProtectionCustomLink",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) SecretScanningPushProtectionCustomLinkInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretScanningPushProtectionCustomLinkInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) SecretScanningPushProtectionEnabledForNewRepositories() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningPushProtectionEnabledForNewRepositories",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) SecretScanningPushProtectionEnabledForNewRepositoriesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningPushProtectionEnabledForNewRepositoriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) SecretScanningValidityChecksEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningValidityChecksEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) SecretScanningValidityChecksEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretScanningValidityChecksEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings github_enterprise_security_analysis_settings} Resource.
func NewEnterpriseSecurityAnalysisSettings(scope constructs.Construct, id *string, config *EnterpriseSecurityAnalysisSettingsConfig) EnterpriseSecurityAnalysisSettings {
	_init_.Initialize()

	if err := validateNewEnterpriseSecurityAnalysisSettingsParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseSecurityAnalysisSettings{}

	_jsii_.Create(
		"@cdktf/provider-github.enterpriseSecurityAnalysisSettings.EnterpriseSecurityAnalysisSettings",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings github_enterprise_security_analysis_settings} Resource.
func NewEnterpriseSecurityAnalysisSettings_Override(e EnterpriseSecurityAnalysisSettings, scope constructs.Construct, id *string, config *EnterpriseSecurityAnalysisSettingsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.enterpriseSecurityAnalysisSettings.EnterpriseSecurityAnalysisSettings",
		[]interface{}{scope, id, config},
		e,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetAdvancedSecurityEnabledForNewRepositories(val interface{}) {
	if err := j.validateSetAdvancedSecurityEnabledForNewRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"advancedSecurityEnabledForNewRepositories",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetEnterpriseSlug(val *string) {
	if err := j.validateSetEnterpriseSlugParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enterpriseSlug",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetSecretScanningEnabledForNewRepositories(val interface{}) {
	if err := j.validateSetSecretScanningEnabledForNewRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretScanningEnabledForNewRepositories",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetSecretScanningPushProtectionCustomLink(val *string) {
	if err := j.validateSetSecretScanningPushProtectionCustomLinkParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretScanningPushProtectionCustomLink",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetSecretScanningPushProtectionEnabledForNewRepositories(val interface{}) {
	if err := j.validateSetSecretScanningPushProtectionEnabledForNewRepositoriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretScanningPushProtectionEnabledForNewRepositories",
		val,
	)
}

func (j *jsiiProxy_EnterpriseSecurityAnalysisSettings)SetSecretScanningValidityChecksEnabled(val interface{}) {
	if err := j.validateSetSecretScanningValidityChecksEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretScanningValidityChecksEnabled",
		val,
	)
}

// Generates CDKTF code for importing a EnterpriseSecurityAnalysisSettings resource upon running "cdktf plan <stack-name>".
func EnterpriseSecurityAnalysisSettings_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateEnterpriseSecurityAnalysisSettings_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseSecurityAnalysisSettings.EnterpriseSecurityAnalysisSettings",
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
func EnterpriseSecurityAnalysisSettings_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseSecurityAnalysisSettings_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseSecurityAnalysisSettings.EnterpriseSecurityAnalysisSettings",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseSecurityAnalysisSettings_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseSecurityAnalysisSettings_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseSecurityAnalysisSettings.EnterpriseSecurityAnalysisSettings",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func EnterpriseSecurityAnalysisSettings_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateEnterpriseSecurityAnalysisSettings_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.enterpriseSecurityAnalysisSettings.EnterpriseSecurityAnalysisSettings",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func EnterpriseSecurityAnalysisSettings_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.enterpriseSecurityAnalysisSettings.EnterpriseSecurityAnalysisSettings",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) AddMoveTarget(moveTarget *string) {
	if err := e.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := e.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) MoveFromId(id *string) {
	if err := e.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveFromId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) MoveTo(moveTarget *string, index interface{}) {
	if err := e.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) MoveToId(id *string) {
	if err := e.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"moveToId",
		[]interface{}{id},
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ResetAdvancedSecurityEnabledForNewRepositories() {
	_jsii_.InvokeVoid(
		e,
		"resetAdvancedSecurityEnabledForNewRepositories",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ResetId() {
	_jsii_.InvokeVoid(
		e,
		"resetId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ResetSecretScanningEnabledForNewRepositories() {
	_jsii_.InvokeVoid(
		e,
		"resetSecretScanningEnabledForNewRepositories",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ResetSecretScanningPushProtectionCustomLink() {
	_jsii_.InvokeVoid(
		e,
		"resetSecretScanningPushProtectionCustomLink",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ResetSecretScanningPushProtectionEnabledForNewRepositories() {
	_jsii_.InvokeVoid(
		e,
		"resetSecretScanningPushProtectionEnabledForNewRepositories",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ResetSecretScanningValidityChecksEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetSecretScanningValidityChecksEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseSecurityAnalysisSettings) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

