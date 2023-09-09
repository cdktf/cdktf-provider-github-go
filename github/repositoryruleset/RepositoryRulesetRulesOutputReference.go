// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v11/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v11/repositoryruleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryRulesetRulesOutputReference interface {
	cdktf.ComplexObject
	BranchNamePattern() RepositoryRulesetRulesBranchNamePatternOutputReference
	BranchNamePatternInput() *RepositoryRulesetRulesBranchNamePattern
	CommitAuthorEmailPattern() RepositoryRulesetRulesCommitAuthorEmailPatternOutputReference
	CommitAuthorEmailPatternInput() *RepositoryRulesetRulesCommitAuthorEmailPattern
	CommitMessagePattern() RepositoryRulesetRulesCommitMessagePatternOutputReference
	CommitMessagePatternInput() *RepositoryRulesetRulesCommitMessagePattern
	CommitterEmailPattern() RepositoryRulesetRulesCommitterEmailPatternOutputReference
	CommitterEmailPatternInput() *RepositoryRulesetRulesCommitterEmailPattern
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
	InternalValue() *RepositoryRulesetRules
	SetInternalValue(val *RepositoryRulesetRules)
	NonFastForward() interface{}
	SetNonFastForward(val interface{})
	NonFastForwardInput() interface{}
	PullRequest() RepositoryRulesetRulesPullRequestOutputReference
	PullRequestInput() *RepositoryRulesetRulesPullRequest
	RequiredDeployments() RepositoryRulesetRulesRequiredDeploymentsOutputReference
	RequiredDeploymentsInput() *RepositoryRulesetRulesRequiredDeployments
	RequiredLinearHistory() interface{}
	SetRequiredLinearHistory(val interface{})
	RequiredLinearHistoryInput() interface{}
	RequiredSignatures() interface{}
	SetRequiredSignatures(val interface{})
	RequiredSignaturesInput() interface{}
	RequiredStatusChecks() RepositoryRulesetRulesRequiredStatusChecksOutputReference
	RequiredStatusChecksInput() *RepositoryRulesetRulesRequiredStatusChecks
	TagNamePattern() RepositoryRulesetRulesTagNamePatternOutputReference
	TagNamePatternInput() *RepositoryRulesetRulesTagNamePattern
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
	UpdateAllowsFetchAndMerge() interface{}
	SetUpdateAllowsFetchAndMerge(val interface{})
	UpdateAllowsFetchAndMergeInput() interface{}
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
	PutBranchNamePattern(value *RepositoryRulesetRulesBranchNamePattern)
	PutCommitAuthorEmailPattern(value *RepositoryRulesetRulesCommitAuthorEmailPattern)
	PutCommitMessagePattern(value *RepositoryRulesetRulesCommitMessagePattern)
	PutCommitterEmailPattern(value *RepositoryRulesetRulesCommitterEmailPattern)
	PutPullRequest(value *RepositoryRulesetRulesPullRequest)
	PutRequiredDeployments(value *RepositoryRulesetRulesRequiredDeployments)
	PutRequiredStatusChecks(value *RepositoryRulesetRulesRequiredStatusChecks)
	PutTagNamePattern(value *RepositoryRulesetRulesTagNamePattern)
	ResetBranchNamePattern()
	ResetCommitAuthorEmailPattern()
	ResetCommitMessagePattern()
	ResetCommitterEmailPattern()
	ResetCreation()
	ResetDeletion()
	ResetNonFastForward()
	ResetPullRequest()
	ResetRequiredDeployments()
	ResetRequiredLinearHistory()
	ResetRequiredSignatures()
	ResetRequiredStatusChecks()
	ResetTagNamePattern()
	ResetUpdate()
	ResetUpdateAllowsFetchAndMerge()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RepositoryRulesetRulesOutputReference
type jsiiProxy_RepositoryRulesetRulesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) BranchNamePattern() RepositoryRulesetRulesBranchNamePatternOutputReference {
	var returns RepositoryRulesetRulesBranchNamePatternOutputReference
	_jsii_.Get(
		j,
		"branchNamePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) BranchNamePatternInput() *RepositoryRulesetRulesBranchNamePattern {
	var returns *RepositoryRulesetRulesBranchNamePattern
	_jsii_.Get(
		j,
		"branchNamePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) CommitAuthorEmailPattern() RepositoryRulesetRulesCommitAuthorEmailPatternOutputReference {
	var returns RepositoryRulesetRulesCommitAuthorEmailPatternOutputReference
	_jsii_.Get(
		j,
		"commitAuthorEmailPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) CommitAuthorEmailPatternInput() *RepositoryRulesetRulesCommitAuthorEmailPattern {
	var returns *RepositoryRulesetRulesCommitAuthorEmailPattern
	_jsii_.Get(
		j,
		"commitAuthorEmailPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) CommitMessagePattern() RepositoryRulesetRulesCommitMessagePatternOutputReference {
	var returns RepositoryRulesetRulesCommitMessagePatternOutputReference
	_jsii_.Get(
		j,
		"commitMessagePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) CommitMessagePatternInput() *RepositoryRulesetRulesCommitMessagePattern {
	var returns *RepositoryRulesetRulesCommitMessagePattern
	_jsii_.Get(
		j,
		"commitMessagePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) CommitterEmailPattern() RepositoryRulesetRulesCommitterEmailPatternOutputReference {
	var returns RepositoryRulesetRulesCommitterEmailPatternOutputReference
	_jsii_.Get(
		j,
		"committerEmailPattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) CommitterEmailPatternInput() *RepositoryRulesetRulesCommitterEmailPattern {
	var returns *RepositoryRulesetRulesCommitterEmailPattern
	_jsii_.Get(
		j,
		"committerEmailPatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) Creation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) CreationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"creationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) Deletion() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) DeletionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"deletionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) InternalValue() *RepositoryRulesetRules {
	var returns *RepositoryRulesetRules
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) NonFastForward() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonFastForward",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) NonFastForwardInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"nonFastForwardInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) PullRequest() RepositoryRulesetRulesPullRequestOutputReference {
	var returns RepositoryRulesetRulesPullRequestOutputReference
	_jsii_.Get(
		j,
		"pullRequest",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) PullRequestInput() *RepositoryRulesetRulesPullRequest {
	var returns *RepositoryRulesetRulesPullRequest
	_jsii_.Get(
		j,
		"pullRequestInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) RequiredDeployments() RepositoryRulesetRulesRequiredDeploymentsOutputReference {
	var returns RepositoryRulesetRulesRequiredDeploymentsOutputReference
	_jsii_.Get(
		j,
		"requiredDeployments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) RequiredDeploymentsInput() *RepositoryRulesetRulesRequiredDeployments {
	var returns *RepositoryRulesetRulesRequiredDeployments
	_jsii_.Get(
		j,
		"requiredDeploymentsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) RequiredLinearHistory() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredLinearHistory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) RequiredLinearHistoryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredLinearHistoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) RequiredSignatures() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredSignatures",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) RequiredSignaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requiredSignaturesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) RequiredStatusChecks() RepositoryRulesetRulesRequiredStatusChecksOutputReference {
	var returns RepositoryRulesetRulesRequiredStatusChecksOutputReference
	_jsii_.Get(
		j,
		"requiredStatusChecks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) RequiredStatusChecksInput() *RepositoryRulesetRulesRequiredStatusChecks {
	var returns *RepositoryRulesetRulesRequiredStatusChecks
	_jsii_.Get(
		j,
		"requiredStatusChecksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) TagNamePattern() RepositoryRulesetRulesTagNamePatternOutputReference {
	var returns RepositoryRulesetRulesTagNamePatternOutputReference
	_jsii_.Get(
		j,
		"tagNamePattern",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) TagNamePatternInput() *RepositoryRulesetRulesTagNamePattern {
	var returns *RepositoryRulesetRulesTagNamePattern
	_jsii_.Get(
		j,
		"tagNamePatternInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) Update() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"update",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) UpdateAllowsFetchAndMerge() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateAllowsFetchAndMerge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) UpdateAllowsFetchAndMergeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateAllowsFetchAndMergeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference) UpdateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"updateInput",
		&returns,
	)
	return returns
}


func NewRepositoryRulesetRulesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RepositoryRulesetRulesOutputReference {
	_init_.Initialize()

	if err := validateNewRepositoryRulesetRulesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositoryRulesetRulesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.repositoryRuleset.RepositoryRulesetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRepositoryRulesetRulesOutputReference_Override(r RepositoryRulesetRulesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.repositoryRuleset.RepositoryRulesetRulesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetCreation(val interface{}) {
	if err := j.validateSetCreationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"creation",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetDeletion(val interface{}) {
	if err := j.validateSetDeletionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deletion",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetInternalValue(val *RepositoryRulesetRules) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetNonFastForward(val interface{}) {
	if err := j.validateSetNonFastForwardParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonFastForward",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetRequiredLinearHistory(val interface{}) {
	if err := j.validateSetRequiredLinearHistoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredLinearHistory",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetRequiredSignatures(val interface{}) {
	if err := j.validateSetRequiredSignaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredSignatures",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetUpdate(val interface{}) {
	if err := j.validateSetUpdateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"update",
		val,
	)
}

func (j *jsiiProxy_RepositoryRulesetRulesOutputReference)SetUpdateAllowsFetchAndMerge(val interface{}) {
	if err := j.validateSetUpdateAllowsFetchAndMergeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"updateAllowsFetchAndMerge",
		val,
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) PutBranchNamePattern(value *RepositoryRulesetRulesBranchNamePattern) {
	if err := r.validatePutBranchNamePatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putBranchNamePattern",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) PutCommitAuthorEmailPattern(value *RepositoryRulesetRulesCommitAuthorEmailPattern) {
	if err := r.validatePutCommitAuthorEmailPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCommitAuthorEmailPattern",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) PutCommitMessagePattern(value *RepositoryRulesetRulesCommitMessagePattern) {
	if err := r.validatePutCommitMessagePatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCommitMessagePattern",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) PutCommitterEmailPattern(value *RepositoryRulesetRulesCommitterEmailPattern) {
	if err := r.validatePutCommitterEmailPatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putCommitterEmailPattern",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) PutPullRequest(value *RepositoryRulesetRulesPullRequest) {
	if err := r.validatePutPullRequestParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putPullRequest",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) PutRequiredDeployments(value *RepositoryRulesetRulesRequiredDeployments) {
	if err := r.validatePutRequiredDeploymentsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRequiredDeployments",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) PutRequiredStatusChecks(value *RepositoryRulesetRulesRequiredStatusChecks) {
	if err := r.validatePutRequiredStatusChecksParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putRequiredStatusChecks",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) PutTagNamePattern(value *RepositoryRulesetRulesTagNamePattern) {
	if err := r.validatePutTagNamePatternParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putTagNamePattern",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetBranchNamePattern() {
	_jsii_.InvokeVoid(
		r,
		"resetBranchNamePattern",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetCommitAuthorEmailPattern() {
	_jsii_.InvokeVoid(
		r,
		"resetCommitAuthorEmailPattern",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetCommitMessagePattern() {
	_jsii_.InvokeVoid(
		r,
		"resetCommitMessagePattern",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetCommitterEmailPattern() {
	_jsii_.InvokeVoid(
		r,
		"resetCommitterEmailPattern",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetCreation() {
	_jsii_.InvokeVoid(
		r,
		"resetCreation",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetDeletion() {
	_jsii_.InvokeVoid(
		r,
		"resetDeletion",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetNonFastForward() {
	_jsii_.InvokeVoid(
		r,
		"resetNonFastForward",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetPullRequest() {
	_jsii_.InvokeVoid(
		r,
		"resetPullRequest",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetRequiredDeployments() {
	_jsii_.InvokeVoid(
		r,
		"resetRequiredDeployments",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetRequiredLinearHistory() {
	_jsii_.InvokeVoid(
		r,
		"resetRequiredLinearHistory",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetRequiredSignatures() {
	_jsii_.InvokeVoid(
		r,
		"resetRequiredSignatures",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetRequiredStatusChecks() {
	_jsii_.InvokeVoid(
		r,
		"resetRequiredStatusChecks",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetTagNamePattern() {
	_jsii_.InvokeVoid(
		r,
		"resetTagNamePattern",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetUpdate() {
	_jsii_.InvokeVoid(
		r,
		"resetUpdate",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ResetUpdateAllowsFetchAndMerge() {
	_jsii_.InvokeVoid(
		r,
		"resetUpdateAllowsFetchAndMerge",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RepositoryRulesetRulesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

