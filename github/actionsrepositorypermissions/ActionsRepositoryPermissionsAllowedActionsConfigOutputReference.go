// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionsrepositorypermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v11/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v11/actionsrepositorypermissions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsRepositoryPermissionsAllowedActionsConfigOutputReference interface {
	cdktf.ComplexObject
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	GithubOwnedAllowed() interface{}
	SetGithubOwnedAllowed(val interface{})
	GithubOwnedAllowedInput() interface{}
	InternalValue() *ActionsRepositoryPermissionsAllowedActionsConfig
	SetInternalValue(val *ActionsRepositoryPermissionsAllowedActionsConfig)
	PatternsAllowed() *[]*string
	SetPatternsAllowed(val *[]*string)
	PatternsAllowedInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VerifiedAllowed() interface{}
	SetVerifiedAllowed(val interface{})
	VerifiedAllowedInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetPatternsAllowed()
	ResetVerifiedAllowed()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ActionsRepositoryPermissionsAllowedActionsConfigOutputReference
type jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GithubOwnedAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOwnedAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GithubOwnedAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOwnedAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) InternalValue() *ActionsRepositoryPermissionsAllowedActionsConfig {
	var returns *ActionsRepositoryPermissionsAllowedActionsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) PatternsAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) PatternsAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) VerifiedAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifiedAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) VerifiedAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifiedAllowedInput",
		&returns,
	)
	return returns
}


func NewActionsRepositoryPermissionsAllowedActionsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActionsRepositoryPermissionsAllowedActionsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewActionsRepositoryPermissionsAllowedActionsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.actionsRepositoryPermissions.ActionsRepositoryPermissionsAllowedActionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActionsRepositoryPermissionsAllowedActionsConfigOutputReference_Override(a ActionsRepositoryPermissionsAllowedActionsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.actionsRepositoryPermissions.ActionsRepositoryPermissionsAllowedActionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference)SetGithubOwnedAllowed(val interface{}) {
	if err := j.validateSetGithubOwnedAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"githubOwnedAllowed",
		val,
	)
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference)SetInternalValue(val *ActionsRepositoryPermissionsAllowedActionsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference)SetPatternsAllowed(val *[]*string) {
	if err := j.validateSetPatternsAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternsAllowed",
		val,
	)
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference)SetVerifiedAllowed(val interface{}) {
	if err := j.validateSetVerifiedAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedAllowed",
		val,
	)
}

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) ResetPatternsAllowed() {
	_jsii_.InvokeVoid(
		a,
		"resetPatternsAllowed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) ResetVerifiedAllowed() {
	_jsii_.InvokeVoid(
		a,
		"resetVerifiedAllowed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsRepositoryPermissionsAllowedActionsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

