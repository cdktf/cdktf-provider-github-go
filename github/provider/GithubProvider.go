// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-github-go/github/v11/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs github}.
type GithubProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AppAuth() *GithubProviderAppAuth
	SetAppAuth(val *GithubProviderAppAuth)
	AppAuthInput() *GithubProviderAppAuth
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Insecure() interface{}
	SetInsecure(val interface{})
	InsecureInput() interface{}
	// Experimental.
	MetaAttributes() *map[string]interface{}
	// The tree node.
	Node() constructs.Node
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	Owner() *string
	SetOwner(val *string)
	OwnerInput() *string
	ParallelRequests() interface{}
	SetParallelRequests(val interface{})
	ParallelRequestsInput() interface{}
	// Experimental.
	RawOverrides() interface{}
	ReadDelayMs() *float64
	SetReadDelayMs(val *float64)
	ReadDelayMsInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	WriteDelayMs() *float64
	SetWriteDelayMs(val *float64)
	WriteDelayMsInput() *float64
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAlias()
	ResetAppAuth()
	ResetBaseUrl()
	ResetInsecure()
	ResetOrganization()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetOwner()
	ResetParallelRequests()
	ResetReadDelayMs()
	ResetToken()
	ResetWriteDelayMs()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for GithubProvider
type jsiiProxy_GithubProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_GithubProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) AppAuth() *GithubProviderAppAuth {
	var returns *GithubProviderAppAuth
	_jsii_.Get(
		j,
		"appAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) AppAuthInput() *GithubProviderAppAuth {
	var returns *GithubProviderAppAuth
	_jsii_.Get(
		j,
		"appAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) Insecure() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) InsecureInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) Owner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"owner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) OwnerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ownerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) ParallelRequests() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parallelRequests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) ParallelRequestsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"parallelRequestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) ReadDelayMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) ReadDelayMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readDelayMsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) WriteDelayMs() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeDelayMs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GithubProvider) WriteDelayMsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"writeDelayMsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs github} Resource.
func NewGithubProvider(scope constructs.Construct, id *string, config *GithubProviderConfig) GithubProvider {
	_init_.Initialize()

	if err := validateNewGithubProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_GithubProvider{}

	_jsii_.Create(
		"@cdktf/provider-github.provider.GithubProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs github} Resource.
func NewGithubProvider_Override(g GithubProvider, scope constructs.Construct, id *string, config *GithubProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.provider.GithubProvider",
		[]interface{}{scope, id, config},
		g,
	)
}

func (j *jsiiProxy_GithubProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_GithubProvider)SetAppAuth(val *GithubProviderAppAuth) {
	if err := j.validateSetAppAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appAuth",
		val,
	)
}

func (j *jsiiProxy_GithubProvider)SetBaseUrl(val *string) {
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_GithubProvider)SetInsecure(val interface{}) {
	if err := j.validateSetInsecureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecure",
		val,
	)
}

func (j *jsiiProxy_GithubProvider)SetOrganization(val *string) {
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_GithubProvider)SetOwner(val *string) {
	_jsii_.Set(
		j,
		"owner",
		val,
	)
}

func (j *jsiiProxy_GithubProvider)SetParallelRequests(val interface{}) {
	if err := j.validateSetParallelRequestsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"parallelRequests",
		val,
	)
}

func (j *jsiiProxy_GithubProvider)SetReadDelayMs(val *float64) {
	_jsii_.Set(
		j,
		"readDelayMs",
		val,
	)
}

func (j *jsiiProxy_GithubProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_GithubProvider)SetWriteDelayMs(val *float64) {
	_jsii_.Set(
		j,
		"writeDelayMs",
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
func GithubProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGithubProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.provider.GithubProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GithubProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGithubProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.provider.GithubProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func GithubProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateGithubProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-github.provider.GithubProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func GithubProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-github.provider.GithubProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (g *jsiiProxy_GithubProvider) AddOverride(path *string, value interface{}) {
	if err := g.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (g *jsiiProxy_GithubProvider) OverrideLogicalId(newLogicalId *string) {
	if err := g.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (g *jsiiProxy_GithubProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		g,
		"resetAlias",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetAppAuth() {
	_jsii_.InvokeVoid(
		g,
		"resetAppAuth",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		g,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetInsecure() {
	_jsii_.InvokeVoid(
		g,
		"resetInsecure",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetOrganization() {
	_jsii_.InvokeVoid(
		g,
		"resetOrganization",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		g,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetOwner() {
	_jsii_.InvokeVoid(
		g,
		"resetOwner",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetParallelRequests() {
	_jsii_.InvokeVoid(
		g,
		"resetParallelRequests",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetReadDelayMs() {
	_jsii_.InvokeVoid(
		g,
		"resetReadDelayMs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetToken() {
	_jsii_.InvokeVoid(
		g,
		"resetToken",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) ResetWriteDelayMs() {
	_jsii_.InvokeVoid(
		g,
		"resetWriteDelayMs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GithubProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GithubProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GithubProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GithubProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		g,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

