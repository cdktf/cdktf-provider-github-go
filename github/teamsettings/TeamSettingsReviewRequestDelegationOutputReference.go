// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package teamsettings

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v12/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v12/teamsettings/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type TeamSettingsReviewRequestDelegationOutputReference interface {
	cdktf.ComplexObject
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
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
	InternalValue() *TeamSettingsReviewRequestDelegation
	SetInternalValue(val *TeamSettingsReviewRequestDelegation)
	MemberCount() *float64
	SetMemberCount(val *float64)
	MemberCountInput() *float64
	Notify() interface{}
	SetNotify(val interface{})
	NotifyInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	ResetAlgorithm()
	ResetMemberCount()
	ResetNotify()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for TeamSettingsReviewRequestDelegationOutputReference
type jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) InternalValue() *TeamSettingsReviewRequestDelegation {
	var returns *TeamSettingsReviewRequestDelegation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) MemberCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memberCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) MemberCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"memberCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) Notify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) NotifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"notifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewTeamSettingsReviewRequestDelegationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) TeamSettingsReviewRequestDelegationOutputReference {
	_init_.Initialize()

	if err := validateNewTeamSettingsReviewRequestDelegationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.teamSettings.TeamSettingsReviewRequestDelegationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewTeamSettingsReviewRequestDelegationOutputReference_Override(t TeamSettingsReviewRequestDelegationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.teamSettings.TeamSettingsReviewRequestDelegationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		t,
	)
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference)SetInternalValue(val *TeamSettingsReviewRequestDelegation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference)SetMemberCount(val *float64) {
	if err := j.validateSetMemberCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"memberCount",
		val,
	)
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference)SetNotify(val interface{}) {
	if err := j.validateSetNotifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"notify",
		val,
	)
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		t,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) ResetMemberCount() {
	_jsii_.InvokeVoid(
		t,
		"resetMemberCount",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) ResetNotify() {
	_jsii_.InvokeVoid(
		t,
		"resetNotify",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := t.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		t,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TeamSettingsReviewRequestDelegationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

