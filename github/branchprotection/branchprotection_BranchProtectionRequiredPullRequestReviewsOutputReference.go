package branchprotection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-github-go/github/v3/jsii"

	"github.com/hashicorp/cdktf-provider-github-go/github/v3/branchprotection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchProtectionRequiredPullRequestReviewsOutputReference interface {
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
	DismissalRestrictions() *[]*string
	SetDismissalRestrictions(val *[]*string)
	DismissalRestrictionsInput() *[]*string
	DismissStaleReviews() interface{}
	SetDismissStaleReviews(val interface{})
	DismissStaleReviewsInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PullRequestBypassers() *[]*string
	SetPullRequestBypassers(val *[]*string)
	PullRequestBypassersInput() *[]*string
	RequireCodeOwnerReviews() interface{}
	SetRequireCodeOwnerReviews(val interface{})
	RequireCodeOwnerReviewsInput() interface{}
	RequiredApprovingReviewCount() *float64
	SetRequiredApprovingReviewCount(val *float64)
	RequiredApprovingReviewCountInput() *float64
	RestrictDismissals() interface{}
	SetRestrictDismissals(val interface{})
	RestrictDismissalsInput() interface{}
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
	ResetDismissalRestrictions()
	ResetDismissStaleReviews()
	ResetPullRequestBypassers()
	ResetRequireCodeOwnerReviews()
	ResetRequiredApprovingReviewCount()
	ResetRestrictDismissals()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BranchProtectionRequiredPullRequestReviewsOutputReference
type jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) DismissalRestrictions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dismissalRestrictions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) DismissalRestrictionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dismissalRestrictionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) DismissStaleReviews() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dismissStaleReviews",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) DismissStaleReviewsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dismissStaleReviewsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) PullRequestBypassers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pullRequestBypassers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) PullRequestBypassersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"pullRequestBypassersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) RequireCodeOwnerReviews() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCodeOwnerReviews",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) RequireCodeOwnerReviewsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requireCodeOwnerReviewsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) RequiredApprovingReviewCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovingReviewCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) RequiredApprovingReviewCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requiredApprovingReviewCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) RestrictDismissals() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictDismissals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) RestrictDismissalsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictDismissalsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBranchProtectionRequiredPullRequestReviewsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BranchProtectionRequiredPullRequestReviewsOutputReference {
	_init_.Initialize()

	if err := validateNewBranchProtectionRequiredPullRequestReviewsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.branchProtection.BranchProtectionRequiredPullRequestReviewsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBranchProtectionRequiredPullRequestReviewsOutputReference_Override(b BranchProtectionRequiredPullRequestReviewsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.branchProtection.BranchProtectionRequiredPullRequestReviewsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetDismissalRestrictions(val *[]*string) {
	if err := j.validateSetDismissalRestrictionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dismissalRestrictions",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetDismissStaleReviews(val interface{}) {
	if err := j.validateSetDismissStaleReviewsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dismissStaleReviews",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetPullRequestBypassers(val *[]*string) {
	if err := j.validateSetPullRequestBypassersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pullRequestBypassers",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetRequireCodeOwnerReviews(val interface{}) {
	if err := j.validateSetRequireCodeOwnerReviewsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requireCodeOwnerReviews",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetRequiredApprovingReviewCount(val *float64) {
	if err := j.validateSetRequiredApprovingReviewCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requiredApprovingReviewCount",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetRestrictDismissals(val interface{}) {
	if err := j.validateSetRestrictDismissalsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictDismissals",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ResetDismissalRestrictions() {
	_jsii_.InvokeVoid(
		b,
		"resetDismissalRestrictions",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ResetDismissStaleReviews() {
	_jsii_.InvokeVoid(
		b,
		"resetDismissStaleReviews",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ResetPullRequestBypassers() {
	_jsii_.InvokeVoid(
		b,
		"resetPullRequestBypassers",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ResetRequireCodeOwnerReviews() {
	_jsii_.InvokeVoid(
		b,
		"resetRequireCodeOwnerReviews",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ResetRequiredApprovingReviewCount() {
	_jsii_.InvokeVoid(
		b,
		"resetRequiredApprovingReviewCount",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ResetRestrictDismissals() {
	_jsii_.InvokeVoid(
		b,
		"resetRestrictDismissals",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

