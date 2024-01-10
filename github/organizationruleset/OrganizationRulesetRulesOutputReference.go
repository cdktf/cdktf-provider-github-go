// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v13/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v13/organizationruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationRulesetRulesOutputReference interface {
	cdktf.ComplexObject
	BranchNamePattern() OrganizationRulesetRulesBranchNamePatternOutputReference
	BranchNamePatternInput() *OrganizationRulesetRulesBranchNamePattern
	CommitAuthorEmailPattern() OrganizationRulesetRulesCommitAuthorEmailPatternOutputReference
	CommitAuthorEmailPatternInput() *OrganizationRulesetRulesCommitAuthorEmailPattern
	CommitMessagePattern() OrganizationRulesetRulesCommitMessagePatternOutputReference
	CommitMessagePatternInput() *OrganizationRulesetRulesCommitMessagePattern
	CommitterEmailPattern() OrganizationRulesetRulesCommitterEmailPatternOutputReference
	CommitterEmailPatternInput() *OrganizationRulesetRulesCommitterEmailPattern
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
	Creation() interface{}
	SetCreation(val interface{})
	CreationInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Deletion() interface{}
	SetDeletion(val interface{})
	DeletionInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *OrganizationRulesetRules
	SetInternalValue(val *OrganizationRulesetRules)
	NonFastForward() interface{}
	SetNonFastForward(val interface{})
	NonFastForwardInput() interface{}
	PullRequest() OrganizationRulesetRulesPullRequestOutputReference
	PullRequestInput() *OrganizationRulesetRulesPullRequest
	RequiredLinearHistory() interface{}
	SetRequiredLinearHistory(val interface{})
	RequiredLinearHistoryInput() interface{}
	RequiredSignatures() interface{}
	SetRequiredSignatures(val interface{})
	RequiredSignaturesInput() interface{}
	RequiredStatusChecks() OrganizationRulesetRulesRequiredStatusChecksOutputReference
	RequiredStatusChecksInput() *OrganizationRulesetRulesRequiredStatusChecks
	TagNamePattern() OrganizationRulesetRulesTagNamePatternOutputReference
	TagNamePatternInput() *OrganizationRulesetRulesTagNamePattern
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Update() interface{}
	SetUpdate(val interface{})
	UpdateInput() interface{}
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
	PutBranchNamePattern(value *OrganizationRulesetRulesBranchNamePattern)
	PutCommitAuthorEmailPattern(value *OrganizationRulesetRulesCommitAuthorEmailPattern)
	PutCommitMessagePattern(value *OrganizationRulesetRulesCommitMessagePattern)
	PutCommitterEmailPattern(value *OrganizationRulesetRulesCommitterEmailPattern)
	PutPullRequest(value *OrganizationRulesetRulesPullRequest)
	PutRequiredStatusChecks(value *OrganizationRulesetRulesRequiredStatusChecks)
	PutTagNamePattern(value *OrganizationRulesetRulesTagNamePattern)
	ResetBranchNamePattern()
	ResetCommitAuthorEmailPattern()
	ResetCommitMessagePattern()
	ResetCommitterEmailPattern()
	ResetCreation()
	ResetDeletion()
	ResetNonFastForward()
	ResetPullRequest()
	ResetRequiredLinearHistory()
	ResetRequiredSignatures()
	ResetRequiredStatusChecks()
	ResetTagNamePattern()
	ResetUpdate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrganizationRulesetRulesOutputReference
type jsiiProxy_OrganizationRulesetRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) BranchNamePattern() OrganizationRulesetRulesBranchNamePatternOutputReference {
	var returns OrganizationRulesetRulesBranchNamePatternOutputReference
	_jsii_.Get(
		j,
		"branchNamePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) BranchNamePatternInput() *OrganizationRulesetRulesBranchNamePattern {
	var returns *OrganizationRulesetRulesBranchNamePattern
	_jsii_.Get(
		j,
		"branchNamePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) CommitAuthorEmailPattern() OrganizationRulesetRulesCommitAuthorEmailPatternOutputReference {
	var returns OrganizationRulesetRulesCommitAuthorEmailPatternOutputReference
	_jsii_.Get(
		j,
		"commitAuthorEmailPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) CommitAuthorEmailPatternInput() *OrganizationRulesetRulesCommitAuthorEmailPattern {
	var returns *OrganizationRulesetRulesCommitAuthorEmailPattern
	_jsii_.Get(
		j,
		"commitAuthorEmailPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) CommitMessagePattern() OrganizationRulesetRulesCommitMessagePatternOutputReference {
	var returns OrganizationRulesetRulesCommitMessagePatternOutputReference
	_jsii_.Get(
		j,
		"commitMessagePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) CommitMessagePatternInput() *OrganizationRulesetRulesCommitMessagePattern {
	var returns *OrganizationRulesetRulesCommitMessagePattern
	_jsii_.Get(
		j,
		"commitMessagePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) CommitterEmailPattern() OrganizationRulesetRulesCommitterEmailPatternOutputReference {
	var returns OrganizationRulesetRulesCommitterEmailPatternOutputReference
	_jsii_.Get(
		j,
		"committerEmailPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) CommitterEmailPatternInput() *OrganizationRulesetRulesCommitterEmailPattern {
	var returns *OrganizationRulesetRulesCommitterEmailPattern
	_jsii_.Get(
		j,
		"committerEmailPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) Creation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) CreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) Deletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) DeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) InternalValue() *OrganizationRulesetRules {
	var returns *OrganizationRulesetRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) NonFastForward() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonFastForward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) NonFastForwardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonFastForwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) PullRequest() OrganizationRulesetRulesPullRequestOutputReference {
	var returns OrganizationRulesetRulesPullRequestOutputReference
	_jsii_.Get(
		j,
		"pullRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) PullRequestInput() *OrganizationRulesetRulesPullRequest {
	var returns *OrganizationRulesetRulesPullRequest
	_jsii_.Get(
		j,
		"pullRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) RequiredLinearHistory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredLinearHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) RequiredLinearHistoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredLinearHistoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) RequiredSignatures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredSignatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) RequiredSignaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredSignaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) RequiredStatusChecks() OrganizationRulesetRulesRequiredStatusChecksOutputReference {
	var returns OrganizationRulesetRulesRequiredStatusChecksOutputReference
	_jsii_.Get(
		j,
		"requiredStatusChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) RequiredStatusChecksInput() *OrganizationRulesetRulesRequiredStatusChecks {
	var returns *OrganizationRulesetRulesRequiredStatusChecks
	_jsii_.Get(
		j,
		"requiredStatusChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) TagNamePattern() OrganizationRulesetRulesTagNamePatternOutputReference {
	var returns OrganizationRulesetRulesTagNamePatternOutputReference
	_jsii_.Get(
		j,
		"tagNamePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) TagNamePatternInput() *OrganizationRulesetRulesTagNamePattern {
	var returns *OrganizationRulesetRulesTagNamePattern
	_jsii_.Get(
		j,
		"tagNamePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) Update() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference) UpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewOrganizationRulesetRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrganizationRulesetRulesOutputReference {
	_init_.Initialize()

	if err := validateNewOrganizationRulesetRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrganizationRulesetRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrganizationRulesetRulesOutputReference_Override(o OrganizationRulesetRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.organizationRuleset.OrganizationRulesetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetCreation(val interface{}) {
	if err := j.validateSetCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creation",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetDeletion(val interface{}) {
	if err := j.validateSetDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletion",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetInternalValue(val *OrganizationRulesetRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetNonFastForward(val interface{}) {
	if err := j.validateSetNonFastForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonFastForward",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetRequiredLinearHistory(val interface{}) {
	if err := j.validateSetRequiredLinearHistoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredLinearHistory",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetRequiredSignatures(val interface{}) {
	if err := j.validateSetRequiredSignaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredSignatures",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_OrganizationRulesetRulesOutputReference)SetUpdate(val interface{}) {
	if err := j.validateSetUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) PutBranchNamePattern(value *OrganizationRulesetRulesBranchNamePattern) {
	if err := o.validatePutBranchNamePatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putBranchNamePattern",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) PutCommitAuthorEmailPattern(value *OrganizationRulesetRulesCommitAuthorEmailPattern) {
	if err := o.validatePutCommitAuthorEmailPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCommitAuthorEmailPattern",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) PutCommitMessagePattern(value *OrganizationRulesetRulesCommitMessagePattern) {
	if err := o.validatePutCommitMessagePatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCommitMessagePattern",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) PutCommitterEmailPattern(value *OrganizationRulesetRulesCommitterEmailPattern) {
	if err := o.validatePutCommitterEmailPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putCommitterEmailPattern",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) PutPullRequest(value *OrganizationRulesetRulesPullRequest) {
	if err := o.validatePutPullRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putPullRequest",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) PutRequiredStatusChecks(value *OrganizationRulesetRulesRequiredStatusChecks) {
	if err := o.validatePutRequiredStatusChecksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putRequiredStatusChecks",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) PutTagNamePattern(value *OrganizationRulesetRulesTagNamePattern) {
	if err := o.validatePutTagNamePatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putTagNamePattern",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetBranchNamePattern() {
	_jsii_.InvokeVoid(
		o,
		"resetBranchNamePattern",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetCommitAuthorEmailPattern() {
	_jsii_.InvokeVoid(
		o,
		"resetCommitAuthorEmailPattern",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetCommitMessagePattern() {
	_jsii_.InvokeVoid(
		o,
		"resetCommitMessagePattern",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetCommitterEmailPattern() {
	_jsii_.InvokeVoid(
		o,
		"resetCommitterEmailPattern",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetCreation() {
	_jsii_.InvokeVoid(
		o,
		"resetCreation",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetDeletion() {
	_jsii_.InvokeVoid(
		o,
		"resetDeletion",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetNonFastForward() {
	_jsii_.InvokeVoid(
		o,
		"resetNonFastForward",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetPullRequest() {
	_jsii_.InvokeVoid(
		o,
		"resetPullRequest",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetRequiredLinearHistory() {
	_jsii_.InvokeVoid(
		o,
		"resetRequiredLinearHistory",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetRequiredSignatures() {
	_jsii_.InvokeVoid(
		o,
		"resetRequiredSignatures",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetRequiredStatusChecks() {
	_jsii_.InvokeVoid(
		o,
		"resetRequiredStatusChecks",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetTagNamePattern() {
	_jsii_.InvokeVoid(
		o,
		"resetTagNamePattern",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		o,
		"resetUpdate",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (o *jsiiProxy_OrganizationRulesetRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

