// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v11/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v11/organizationruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRulesetConditionsRepositoryNameOutputReference interface {
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
	Exclude() *[]*string
	SetExclude(val *[]*string)
	ExcludeInput() *[]*string
	// Experimental.
	Fqn() *string
	Inlcude() *[]*string
	SetInlcude(val *[]*string)
	InlcudeInput() *[]*string
	InternalValue() *OrganizationRulesetConditionsRepositoryName
	SetInternalValue(val *OrganizationRulesetConditionsRepositoryName)
	Protected() interface{}
	SetProtected(val interface{})
	ProtectedInput() interface{}
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
	ResetProtected()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrganizationRulesetConditionsRepositoryNameOutputReference
type jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) Exclude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exclude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) ExcludeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) Inlcude() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inlcude",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) InlcudeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inlcudeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) InternalValue() *OrganizationRulesetConditionsRepositoryName {
	var returns *OrganizationRulesetConditionsRepositoryName
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) Protected() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protected",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) ProtectedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrganizationRulesetConditionsRepositoryNameOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrganizationRulesetConditionsRepositoryNameOutputReference {
	_init_.Initialize()

	if err := validateNewOrganizationRulesetConditionsRepositoryNameOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetConditionsRepositoryNameOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrganizationRulesetConditionsRepositoryNameOutputReference_Override(o OrganizationRulesetConditionsRepositoryNameOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetConditionsRepositoryNameOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference)SetExclude(val *[]*string) {
	if err := j.validateSetExcludeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exclude",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference)SetInlcude(val *[]*string) {
	if err := j.validateSetInlcudeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inlcude",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference)SetInternalValue(val *OrganizationRulesetConditionsRepositoryName) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference)SetProtected(val interface{}) {
	if err := j.validateSetProtectedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protected",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) ResetProtected() {
	_jsii_.InvokeVoid(
		o,
		"resetProtected",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OrganizationRulesetConditionsRepositoryNameOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

