// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v14/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v14/organizationruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList interface {
	cdktf.ComplexList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// The attribute on the parent resource this class is referencing.
	TerraformAttribute() *string
	SetTerraformAttribute(val *string)
	// The parent resource.
	TerraformResource() cdktf.IInterpolatingParent
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// whether the list is wrapping a set (will add tolist() to be able to access an item via an index).
	WrapsSet() *bool
	SetWrapsSet(val *bool)
	// Creating an iterator for this complex list.
	//
	// The list will be converted into a map with the mapKeyAttributeName as the key.
	// Experimental.
	AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList
type jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewOrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList {
	_init_.Initialize()

	if err := validateNewOrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolListParameters(terraformResource, terraformAttribute, wrapsSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList{}

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewOrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList_Override(o OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		o,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList)SetWrapsSet(val *bool) {
	if err := j.validateSetWrapsSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) AllWithMapKey(mapKeyAttributeName *string) cdktf.DynamicListTerraformIterator {
	if err := o.validateAllWithMapKeyParameters(mapKeyAttributeName); err != nil {
		panic(err)
	}
	var returns cdktf.DynamicListTerraformIterator

	_jsii_.Invoke(
		o,
		"allWithMapKey",
		[]interface{}{mapKeyAttributeName},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) Get(index *float64) OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolOutputReference {
	if err := o.validateGetParameters(index); err != nil {
		panic(err)
	}
	var returns OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolOutputReference

	_jsii_.Invoke(
		o,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OrganizationRulesetRulesRequiredCodeScanningRequiredCodeScanningToolList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
