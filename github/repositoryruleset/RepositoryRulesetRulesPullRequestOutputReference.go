// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v12/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v12/repositoryruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryRulesetRulesPullRequestOutputReference interface {
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
	DismissStaleReviewsOnPush() interface{}
	SetDismissStaleReviewsOnPush(val interface{})
	DismissStaleReviewsOnPushInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *RepositoryRulesetRulesPullRequest
	SetInternalValue(val *RepositoryRulesetRulesPullRequest)
	RequireCodeOwnerReview() interface{}
	SetRequireCodeOwnerReview(val interface{})
	RequireCodeOwnerReviewInput() interface{}
	RequiredApprovingReviewCount() *float64
	SetRequiredApprovingReviewCount(val *float64)
	RequiredApprovingReviewCountInput() *float64
	RequiredReviewThreadResolution() interface{}
	SetRequiredReviewThreadResolution(val interface{})
	RequiredReviewThreadResolutionInput() interface{}
	RequireLastPushApproval() interface{}
	SetRequireLastPushApproval(val interface{})
	RequireLastPushApprovalInput() interface{}
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
	ResetDismissStaleReviewsOnPush()
	ResetRequireCodeOwnerReview()
	ResetRequiredApprovingReviewCount()
	ResetRequiredReviewThreadResolution()
	ResetRequireLastPushApproval()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RepositoryRulesetRulesPullRequestOutputReference
type jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) DismissStaleReviewsOnPush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dismissStaleReviewsOnPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) DismissStaleReviewsOnPushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dismissStaleReviewsOnPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) InternalValue() *RepositoryRulesetRulesPullRequest {
	var returns *RepositoryRulesetRulesPullRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) RequireCodeOwnerReview() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCodeOwnerReview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) RequireCodeOwnerReviewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCodeOwnerReviewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) RequiredApprovingReviewCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovingReviewCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) RequiredApprovingReviewCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovingReviewCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) RequiredReviewThreadResolution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredReviewThreadResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) RequiredReviewThreadResolutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredReviewThreadResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) RequireLastPushApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLastPushApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) RequireLastPushApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLastPushApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRepositoryRulesetRulesPullRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RepositoryRulesetRulesPullRequestOutputReference {
	_init_.Initialize()

	if err := validateNewRepositoryRulesetRulesPullRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.repositoryRuleset.RepositoryRulesetRulesPullRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRepositoryRulesetRulesPullRequestOutputReference_Override(r RepositoryRulesetRulesPullRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.repositoryRuleset.RepositoryRulesetRulesPullRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetDismissStaleReviewsOnPush(val interface{}) {
	if err := j.validateSetDismissStaleReviewsOnPushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dismissStaleReviewsOnPush",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetInternalValue(val *RepositoryRulesetRulesPullRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetRequireCodeOwnerReview(val interface{}) {
	if err := j.validateSetRequireCodeOwnerReviewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCodeOwnerReview",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetRequiredApprovingReviewCount(val *float64) {
	if err := j.validateSetRequiredApprovingReviewCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredApprovingReviewCount",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetRequiredReviewThreadResolution(val interface{}) {
	if err := j.validateSetRequiredReviewThreadResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredReviewThreadResolution",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetRequireLastPushApproval(val interface{}) {
	if err := j.validateSetRequireLastPushApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireLastPushApproval",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) ResetDismissStaleReviewsOnPush() {
	_jsii_.InvokeVoid(
		r,
		"resetDismissStaleReviewsOnPush",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) ResetRequireCodeOwnerReview() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireCodeOwnerReview",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) ResetRequiredApprovingReviewCount() {
	_jsii_.InvokeVoid(
		r,
		"resetRequiredApprovingReviewCount",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) ResetRequiredReviewThreadResolution() {
	_jsii_.InvokeVoid(
		r,
		"resetRequiredReviewThreadResolution",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) ResetRequireLastPushApproval() {
	_jsii_.InvokeVoid(
		r,
		"resetRequireLastPushApproval",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RepositoryRulesetRulesPullRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

