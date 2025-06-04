// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package actionsorganizationpermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v15/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v15/actionsorganizationpermissions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ActionsOrganizationPermissionsAllowedActionsConfigOutputReference interface {
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
	InternalValue() *ActionsOrganizationPermissionsAllowedActionsConfig
	SetInternalValue(val *ActionsOrganizationPermissionsAllowedActionsConfig)
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

// The jsii proxy struct for ActionsOrganizationPermissionsAllowedActionsConfigOutputReference
type jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GithubOwnedAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOwnedAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GithubOwnedAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOwnedAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) InternalValue() *ActionsOrganizationPermissionsAllowedActionsConfig {
	var returns *ActionsOrganizationPermissionsAllowedActionsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) PatternsAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) PatternsAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) VerifiedAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifiedAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) VerifiedAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifiedAllowedInput",
		&returns,
	)
	return returns
}


func NewActionsOrganizationPermissionsAllowedActionsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ActionsOrganizationPermissionsAllowedActionsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewActionsOrganizationPermissionsAllowedActionsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.actionsOrganizationPermissions.ActionsOrganizationPermissionsAllowedActionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewActionsOrganizationPermissionsAllowedActionsConfigOutputReference_Override(a ActionsOrganizationPermissionsAllowedActionsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.actionsOrganizationPermissions.ActionsOrganizationPermissionsAllowedActionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference)SetGithubOwnedAllowed(val interface{}) {
	if err := j.validateSetGithubOwnedAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"githubOwnedAllowed",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference)SetInternalValue(val *ActionsOrganizationPermissionsAllowedActionsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference)SetPatternsAllowed(val *[]*string) {
	if err := j.validateSetPatternsAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternsAllowed",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference)SetVerifiedAllowed(val interface{}) {
	if err := j.validateSetVerifiedAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedAllowed",
		val,
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) ResetPatternsAllowed() {
	_jsii_.InvokeVoid(
		a,
		"resetPatternsAllowed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) ResetVerifiedAllowed() {
	_jsii_.InvokeVoid(
		a,
		"resetVerifiedAllowed",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ActionsOrganizationPermissionsAllowedActionsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

