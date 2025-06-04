// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v15/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v15/organizationruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRulesetRulesPullRequestOutputReference interface {
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
	InternalValue() *OrganizationRulesetRulesPullRequest
	SetInternalValue(val *OrganizationRulesetRulesPullRequest)
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

// The jsii proxy struct for OrganizationRulesetRulesPullRequestOutputReference
type jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) DismissStaleReviewsOnPush() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dismissStaleReviewsOnPush",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) DismissStaleReviewsOnPushInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dismissStaleReviewsOnPushInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) InternalValue() *OrganizationRulesetRulesPullRequest {
	var returns *OrganizationRulesetRulesPullRequest
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) RequireCodeOwnerReview() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCodeOwnerReview",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) RequireCodeOwnerReviewInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCodeOwnerReviewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) RequiredApprovingReviewCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovingReviewCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) RequiredApprovingReviewCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovingReviewCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) RequiredReviewThreadResolution() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredReviewThreadResolution",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) RequiredReviewThreadResolutionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredReviewThreadResolutionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) RequireLastPushApproval() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLastPushApproval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) RequireLastPushApprovalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireLastPushApprovalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrganizationRulesetRulesPullRequestOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrganizationRulesetRulesPullRequestOutputReference {
	_init_.Initialize()

	if err := validateNewOrganizationRulesetRulesPullRequestOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetRulesPullRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrganizationRulesetRulesPullRequestOutputReference_Override(o OrganizationRulesetRulesPullRequestOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetRulesPullRequestOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetDismissStaleReviewsOnPush(val interface{}) {
	if err := j.validateSetDismissStaleReviewsOnPushParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dismissStaleReviewsOnPush",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetInternalValue(val *OrganizationRulesetRulesPullRequest) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetRequireCodeOwnerReview(val interface{}) {
	if err := j.validateSetRequireCodeOwnerReviewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCodeOwnerReview",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetRequiredApprovingReviewCount(val *float64) {
	if err := j.validateSetRequiredApprovingReviewCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredApprovingReviewCount",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetRequiredReviewThreadResolution(val interface{}) {
	if err := j.validateSetRequiredReviewThreadResolutionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredReviewThreadResolution",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetRequireLastPushApproval(val interface{}) {
	if err := j.validateSetRequireLastPushApprovalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireLastPushApproval",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) ResetDismissStaleReviewsOnPush() {
	_jsii_.InvokeVoid(
		o,
		"resetDismissStaleReviewsOnPush",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) ResetRequireCodeOwnerReview() {
	_jsii_.InvokeVoid(
		o,
		"resetRequireCodeOwnerReview",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) ResetRequiredApprovingReviewCount() {
	_jsii_.InvokeVoid(
		o,
		"resetRequiredApprovingReviewCount",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) ResetRequiredReviewThreadResolution() {
	_jsii_.InvokeVoid(
		o,
		"resetRequiredReviewThreadResolution",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) ResetRequireLastPushApproval() {
	_jsii_.InvokeVoid(
		o,
		"resetRequireLastPushApproval",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OrganizationRulesetRulesPullRequestOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

