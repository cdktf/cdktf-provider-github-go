package branchprotectionv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v3/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v3/branchprotectionv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchProtectionV3RequiredPullRequestReviewsOutputReference interface {
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
	DismissalTeams() *[]*string
	SetDismissalTeams(val *[]*string)
	DismissalTeamsInput() *[]*string
	DismissalUsers() *[]*string
	SetDismissalUsers(val *[]*string)
	DismissalUsersInput() *[]*string
	DismissStaleReviews() interface{}
	SetDismissStaleReviews(val interface{})
	DismissStaleReviewsInput() interface{}
	// Experimental.
	Fqn() *string
	IncludeAdmins() interface{}
	SetIncludeAdmins(val interface{})
	IncludeAdminsInput() interface{}
	InternalValue() *BranchProtectionV3RequiredPullRequestReviews
	SetInternalValue(val *BranchProtectionV3RequiredPullRequestReviews)
	RequireCodeOwnerReviews() interface{}
	SetRequireCodeOwnerReviews(val interface{})
	RequireCodeOwnerReviewsInput() interface{}
	RequiredApprovingReviewCount() *float64
	SetRequiredApprovingReviewCount(val *float64)
	RequiredApprovingReviewCountInput() *float64
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
	ResetDismissalTeams()
	ResetDismissalUsers()
	ResetDismissStaleReviews()
	ResetIncludeAdmins()
	ResetRequireCodeOwnerReviews()
	ResetRequiredApprovingReviewCount()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BranchProtectionV3RequiredPullRequestReviewsOutputReference
type jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) DismissalTeams() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dismissalTeams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) DismissalTeamsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dismissalTeamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) DismissalUsers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dismissalUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) DismissalUsersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dismissalUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) DismissStaleReviews() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dismissStaleReviews",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) DismissStaleReviewsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dismissStaleReviewsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) IncludeAdmins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) IncludeAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) InternalValue() *BranchProtectionV3RequiredPullRequestReviews {
	var returns *BranchProtectionV3RequiredPullRequestReviews
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) RequireCodeOwnerReviews() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCodeOwnerReviews",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) RequireCodeOwnerReviewsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCodeOwnerReviewsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) RequiredApprovingReviewCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovingReviewCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) RequiredApprovingReviewCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovingReviewCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBranchProtectionV3RequiredPullRequestReviewsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BranchProtectionV3RequiredPullRequestReviewsOutputReference {
	_init_.Initialize()

	if err := validateNewBranchProtectionV3RequiredPullRequestReviewsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredPullRequestReviewsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBranchProtectionV3RequiredPullRequestReviewsOutputReference_Override(b BranchProtectionV3RequiredPullRequestReviewsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredPullRequestReviewsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetDismissalTeams(val *[]*string) {
	if err := j.validateSetDismissalTeamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dismissalTeams",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetDismissalUsers(val *[]*string) {
	if err := j.validateSetDismissalUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dismissalUsers",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetDismissStaleReviews(val interface{}) {
	if err := j.validateSetDismissStaleReviewsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dismissStaleReviews",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetIncludeAdmins(val interface{}) {
	if err := j.validateSetIncludeAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeAdmins",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetInternalValue(val *BranchProtectionV3RequiredPullRequestReviews) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetRequireCodeOwnerReviews(val interface{}) {
	if err := j.validateSetRequireCodeOwnerReviewsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCodeOwnerReviews",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetRequiredApprovingReviewCount(val *float64) {
	if err := j.validateSetRequiredApprovingReviewCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredApprovingReviewCount",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ResetDismissalTeams() {
	_jsii_.InvokeVoid(
		b,
		"resetDismissalTeams",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ResetDismissalUsers() {
	_jsii_.InvokeVoid(
		b,
		"resetDismissalUsers",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ResetDismissStaleReviews() {
	_jsii_.InvokeVoid(
		b,
		"resetDismissStaleReviews",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ResetIncludeAdmins() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeAdmins",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ResetRequireCodeOwnerReviews() {
	_jsii_.InvokeVoid(
		b,
		"resetRequireCodeOwnerReviews",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ResetRequiredApprovingReviewCount() {
	_jsii_.InvokeVoid(
		b,
		"resetRequiredApprovingReviewCount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

