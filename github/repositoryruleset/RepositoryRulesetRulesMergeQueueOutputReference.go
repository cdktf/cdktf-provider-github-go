// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v14/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v14/repositoryruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryRulesetRulesMergeQueueOutputReference interface {
	cdktf.ComplexObject
	CheckResponseTimeoutMinutes() *float64
	SetCheckResponseTimeoutMinutes(val *float64)
	CheckResponseTimeoutMinutesInput() *float64
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
	GroupingStrategy() *string
	SetGroupingStrategy(val *string)
	GroupingStrategyInput() *string
	InternalValue() *RepositoryRulesetRulesMergeQueue
	SetInternalValue(val *RepositoryRulesetRulesMergeQueue)
	MaxEntriesToBuild() *float64
	SetMaxEntriesToBuild(val *float64)
	MaxEntriesToBuildInput() *float64
	MaxEntriesToMerge() *float64
	SetMaxEntriesToMerge(val *float64)
	MaxEntriesToMergeInput() *float64
	MergeMethod() *string
	SetMergeMethod(val *string)
	MergeMethodInput() *string
	MinEntriesToMerge() *float64
	SetMinEntriesToMerge(val *float64)
	MinEntriesToMergeInput() *float64
	MinEntriesToMergeWaitMinutes() *float64
	SetMinEntriesToMergeWaitMinutes(val *float64)
	MinEntriesToMergeWaitMinutesInput() *float64
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
	ResetCheckResponseTimeoutMinutes()
	ResetGroupingStrategy()
	ResetMaxEntriesToBuild()
	ResetMaxEntriesToMerge()
	ResetMergeMethod()
	ResetMinEntriesToMerge()
	ResetMinEntriesToMergeWaitMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RepositoryRulesetRulesMergeQueueOutputReference
type jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) CheckResponseTimeoutMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkResponseTimeoutMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) CheckResponseTimeoutMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"checkResponseTimeoutMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GroupingStrategy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingStrategy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GroupingStrategyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupingStrategyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) InternalValue() *RepositoryRulesetRulesMergeQueue {
	var returns *RepositoryRulesetRulesMergeQueue
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MaxEntriesToBuild() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEntriesToBuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MaxEntriesToBuildInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEntriesToBuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MaxEntriesToMerge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEntriesToMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MaxEntriesToMergeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxEntriesToMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MergeMethod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeMethod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MergeMethodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mergeMethodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MinEntriesToMerge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEntriesToMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MinEntriesToMergeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEntriesToMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MinEntriesToMergeWaitMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEntriesToMergeWaitMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) MinEntriesToMergeWaitMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"minEntriesToMergeWaitMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRepositoryRulesetRulesMergeQueueOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RepositoryRulesetRulesMergeQueueOutputReference {
	_init_.Initialize()

	if err := validateNewRepositoryRulesetRulesMergeQueueOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.repositoryRuleset.RepositoryRulesetRulesMergeQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRepositoryRulesetRulesMergeQueueOutputReference_Override(r RepositoryRulesetRulesMergeQueueOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.repositoryRuleset.RepositoryRulesetRulesMergeQueueOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetCheckResponseTimeoutMinutes(val *float64) {
	if err := j.validateSetCheckResponseTimeoutMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checkResponseTimeoutMinutes",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetGroupingStrategy(val *string) {
	if err := j.validateSetGroupingStrategyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupingStrategy",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetInternalValue(val *RepositoryRulesetRulesMergeQueue) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetMaxEntriesToBuild(val *float64) {
	if err := j.validateSetMaxEntriesToBuildParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxEntriesToBuild",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetMaxEntriesToMerge(val *float64) {
	if err := j.validateSetMaxEntriesToMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxEntriesToMerge",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetMergeMethod(val *string) {
	if err := j.validateSetMergeMethodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mergeMethod",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetMinEntriesToMerge(val *float64) {
	if err := j.validateSetMinEntriesToMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minEntriesToMerge",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetMinEntriesToMergeWaitMinutes(val *float64) {
	if err := j.validateSetMinEntriesToMergeWaitMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minEntriesToMergeWaitMinutes",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ResetCheckResponseTimeoutMinutes() {
	_jsii_.InvokeVoid(
		r,
		"resetCheckResponseTimeoutMinutes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ResetGroupingStrategy() {
	_jsii_.InvokeVoid(
		r,
		"resetGroupingStrategy",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ResetMaxEntriesToBuild() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxEntriesToBuild",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ResetMaxEntriesToMerge() {
	_jsii_.InvokeVoid(
		r,
		"resetMaxEntriesToMerge",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ResetMergeMethod() {
	_jsii_.InvokeVoid(
		r,
		"resetMergeMethod",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ResetMinEntriesToMerge() {
	_jsii_.InvokeVoid(
		r,
		"resetMinEntriesToMerge",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ResetMinEntriesToMergeWaitMinutes() {
	_jsii_.InvokeVoid(
		r,
		"resetMinEntriesToMergeWaitMinutes",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesMergeQueueOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

