// Prebuilt github Provider for Terraform CDK (cdktf)
package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-github-go/github/v2/jsii"

	"github.com/hashicorp/cdktf-provider-github-go/github/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchProtectionRequiredPullRequestReviewsList interface {
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
	// Experimental.
	ComputeFqn() *string
	Get(index *float64) BranchProtectionRequiredPullRequestReviewsOutputReference
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BranchProtectionRequiredPullRequestReviewsList
type jsiiProxy_BranchProtectionRequiredPullRequestReviewsList struct {
	internal.Type__cdktfComplexList
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) WrapsSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"wrapsSet",
		&returns,
	)
	return returns
}


func NewBranchProtectionRequiredPullRequestReviewsList(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) BranchProtectionRequiredPullRequestReviewsList {
	_init_.Initialize()

	j := jsiiProxy_BranchProtectionRequiredPullRequestReviewsList{}

	_jsii_.Create(
		"@cdktf/provider-github.BranchProtectionRequiredPullRequestReviewsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		&j,
	)

	return &j
}

func NewBranchProtectionRequiredPullRequestReviewsList_Override(b BranchProtectionRequiredPullRequestReviewsList, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.BranchProtectionRequiredPullRequestReviewsList",
		[]interface{}{terraformResource, terraformAttribute, wrapsSet},
		b,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) SetInternalValue(val interface{}) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) SetWrapsSet(val *bool) {
	_jsii_.Set(
		j,
		"wrapsSet",
		val,
	)
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) Get(index *float64) BranchProtectionRequiredPullRequestReviewsOutputReference {
	var returns BranchProtectionRequiredPullRequestReviewsOutputReference

	_jsii_.Invoke(
		b,
		"get",
		[]interface{}{index},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

