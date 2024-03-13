// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterpriseactionspermissions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v14/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v14/enterpriseactionspermissions/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseActionsPermissionsAllowedActionsConfigOutputReference interface {
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
	InternalValue() *EnterpriseActionsPermissionsAllowedActionsConfig
	SetInternalValue(val *EnterpriseActionsPermissionsAllowedActionsConfig)
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

// The jsii proxy struct for EnterpriseActionsPermissionsAllowedActionsConfigOutputReference
type jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GithubOwnedAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOwnedAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GithubOwnedAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"githubOwnedAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) InternalValue() *EnterpriseActionsPermissionsAllowedActionsConfig {
	var returns *EnterpriseActionsPermissionsAllowedActionsConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) PatternsAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) PatternsAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"patternsAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) VerifiedAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifiedAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) VerifiedAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifiedAllowedInput",
		&returns,
	)
	return returns
}


func NewEnterpriseActionsPermissionsAllowedActionsConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EnterpriseActionsPermissionsAllowedActionsConfigOutputReference {
	_init_.Initialize()

	if err := validateNewEnterpriseActionsPermissionsAllowedActionsConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.enterpriseActionsPermissions.EnterpriseActionsPermissionsAllowedActionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEnterpriseActionsPermissionsAllowedActionsConfigOutputReference_Override(e EnterpriseActionsPermissionsAllowedActionsConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.enterpriseActionsPermissions.EnterpriseActionsPermissionsAllowedActionsConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference)SetGithubOwnedAllowed(val interface{}) {
	if err := j.validateSetGithubOwnedAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"githubOwnedAllowed",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference)SetInternalValue(val *EnterpriseActionsPermissionsAllowedActionsConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference)SetPatternsAllowed(val *[]*string) {
	if err := j.validateSetPatternsAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patternsAllowed",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference)SetVerifiedAllowed(val interface{}) {
	if err := j.validateSetVerifiedAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifiedAllowed",
		val,
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) ResetPatternsAllowed() {
	_jsii_.InvokeVoid(
		e,
		"resetPatternsAllowed",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) ResetVerifiedAllowed() {
	_jsii_.InvokeVoid(
		e,
		"resetVerifiedAllowed",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EnterpriseActionsPermissionsAllowedActionsConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

