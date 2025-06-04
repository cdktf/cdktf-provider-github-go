// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v15/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v15/organizationruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRulesetRulesRequiredStatusChecksOutputReference interface {
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
	InternalValue() *OrganizationRulesetRulesRequiredStatusChecks
	SetInternalValue(val *OrganizationRulesetRulesRequiredStatusChecks)
	RequiredCheck() OrganizationRulesetRulesRequiredStatusChecksRequiredCheckList
	RequiredCheckInput() interface{}
	StrictRequiredStatusChecksPolicy() interface{}
	SetStrictRequiredStatusChecksPolicy(val interface{})
	StrictRequiredStatusChecksPolicyInput() interface{}
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
	PutRequiredCheck(value interface{})
	ResetStrictRequiredStatusChecksPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrganizationRulesetRulesRequiredStatusChecksOutputReference
type jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) InternalValue() *OrganizationRulesetRulesRequiredStatusChecks {
	var returns *OrganizationRulesetRulesRequiredStatusChecks
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) RequiredCheck() OrganizationRulesetRulesRequiredStatusChecksRequiredCheckList {
	var returns OrganizationRulesetRulesRequiredStatusChecksRequiredCheckList
	_jsii_.Get(
		j,
		"requiredCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) RequiredCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) StrictRequiredStatusChecksPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictRequiredStatusChecksPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) StrictRequiredStatusChecksPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictRequiredStatusChecksPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrganizationRulesetRulesRequiredStatusChecksOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrganizationRulesetRulesRequiredStatusChecksOutputReference {
	_init_.Initialize()

	if err := validateNewOrganizationRulesetRulesRequiredStatusChecksOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetRulesRequiredStatusChecksOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrganizationRulesetRulesRequiredStatusChecksOutputReference_Override(o OrganizationRulesetRulesRequiredStatusChecksOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetRulesRequiredStatusChecksOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference)SetInternalValue(val *OrganizationRulesetRulesRequiredStatusChecks) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference)SetStrictRequiredStatusChecksPolicy(val interface{}) {
	if err := j.validateSetStrictRequiredStatusChecksPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strictRequiredStatusChecksPolicy",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) PutRequiredCheck(value interface{}) {
	if err := o.validatePutRequiredCheckParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRequiredCheck",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) ResetStrictRequiredStatusChecksPolicy() {
	_jsii_.InvokeVoid(
		o,
		"resetStrictRequiredStatusChecksPolicy",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredStatusChecksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

