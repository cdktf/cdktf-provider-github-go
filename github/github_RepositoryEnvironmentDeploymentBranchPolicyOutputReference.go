// Prebuilt github Provider for Terraform CDK (cdktf)
package github

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-github-go/github/v2/jsii"

	"github.com/hashicorp/cdktf-provider-github-go/github/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryEnvironmentDeploymentBranchPolicyOutputReference interface {
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
	CustomBranchPolicies() interface{}
	SetCustomBranchPolicies(val interface{})
	CustomBranchPoliciesInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *RepositoryEnvironmentDeploymentBranchPolicy
	SetInternalValue(val *RepositoryEnvironmentDeploymentBranchPolicy)
	ProtectedBranches() interface{}
	SetProtectedBranches(val interface{})
	ProtectedBranchesInput() interface{}
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RepositoryEnvironmentDeploymentBranchPolicyOutputReference
type jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) CustomBranchPolicies() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customBranchPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) CustomBranchPoliciesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"customBranchPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) InternalValue() *RepositoryEnvironmentDeploymentBranchPolicy {
	var returns *RepositoryEnvironmentDeploymentBranchPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) ProtectedBranches() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedBranches",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) ProtectedBranchesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"protectedBranchesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRepositoryEnvironmentDeploymentBranchPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RepositoryEnvironmentDeploymentBranchPolicyOutputReference {
	_init_.Initialize()

	j := jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.RepositoryEnvironmentDeploymentBranchPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRepositoryEnvironmentDeploymentBranchPolicyOutputReference_Override(r RepositoryEnvironmentDeploymentBranchPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.RepositoryEnvironmentDeploymentBranchPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) SetComplexObjectIndex(val interface{}) {
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) SetComplexObjectIsFromSet(val *bool) {
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) SetCustomBranchPolicies(val interface{}) {
	_jsii_.Set(
		j,
		"customBranchPolicies",
		val,
	)
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) SetInternalValue(val *RepositoryEnvironmentDeploymentBranchPolicy) {
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) SetProtectedBranches(val interface{}) {
	_jsii_.Set(
		j,
		"protectedBranches",
		val,
	)
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) SetTerraformAttribute(val *string) {
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositoryEnvironmentDeploymentBranchPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

