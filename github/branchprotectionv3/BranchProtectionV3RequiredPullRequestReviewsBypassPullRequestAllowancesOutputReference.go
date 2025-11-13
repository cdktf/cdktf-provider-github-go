// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotectionv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v15/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v15/branchprotectionv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference interface {
	cdktf.ComplexObject
	Apps() *[]*string
	SetApps(val *[]*string)
	AppsInput() *[]*string
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
	InternalValue() *BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowances
	SetInternalValue(val *BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowances)
	Teams() *[]*string
	SetTeams(val *[]*string)
	TeamsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Users() *[]*string
	SetUsers(val *[]*string)
	UsersInput() *[]*string
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	ResetApps()
	ResetTeams()
	ResetUsers()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference
type jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) Apps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"apps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) AppsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"appsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) InternalValue() *BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowances {
	var returns *BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowances
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) Teams() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"teams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) TeamsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"teamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) Users() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"users",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) UsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"usersInput",
		&returns,
	)
	return returns
}


func NewBranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference {
	_init_.Initialize()

	if err := validateNewBranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference_Override(b BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference)SetApps(val *[]*string) {
	if err := j.validateSetAppsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apps",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference)SetInternalValue(val *BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowances) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference)SetTeams(val *[]*string) {
	if err := j.validateSetTeamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"teams",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference)SetUsers(val *[]*string) {
	if err := j.validateSetUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"users",
		val,
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) ResetApps() {
	_jsii_.InvokeVoid(
		b,
		"resetApps",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) ResetTeams() {
	_jsii_.InvokeVoid(
		b,
		"resetTeams",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) ResetUsers() {
	_jsii_.InvokeVoid(
		b,
		"resetUsers",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowancesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

