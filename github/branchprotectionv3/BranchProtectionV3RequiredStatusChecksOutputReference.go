package branchprotectionv3

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v10/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v10/branchprotectionv3/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BranchProtectionV3RequiredStatusChecksOutputReference interface {
	cdktf.ComplexObject
	Checks() *[]*string
	SetChecks(val *[]*string)
	ChecksInput() *[]*string
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
	Contexts() *[]*string
	SetContexts(val *[]*string)
	ContextsInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IncludeAdmins() interface{}
	SetIncludeAdmins(val interface{})
	IncludeAdminsInput() interface{}
	InternalValue() *BranchProtectionV3RequiredStatusChecks
	SetInternalValue(val *BranchProtectionV3RequiredStatusChecks)
	Strict() interface{}
	SetStrict(val interface{})
	StrictInput() interface{}
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
	ResetChecks()
	ResetContexts()
	ResetIncludeAdmins()
	ResetStrict()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BranchProtectionV3RequiredStatusChecksOutputReference
type jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) Checks() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"checks",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ChecksInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"checksInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) Contexts() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contexts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ContextsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contextsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) IncludeAdmins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAdmins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) IncludeAdminsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includeAdminsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) InternalValue() *BranchProtectionV3RequiredStatusChecks {
	var returns *BranchProtectionV3RequiredStatusChecks
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) Strict() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strict",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) StrictInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"strictInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBranchProtectionV3RequiredStatusChecksOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) BranchProtectionV3RequiredStatusChecksOutputReference {
	_init_.Initialize()

	if err := validateNewBranchProtectionV3RequiredStatusChecksOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredStatusChecksOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewBranchProtectionV3RequiredStatusChecksOutputReference_Override(b BranchProtectionV3RequiredStatusChecksOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredStatusChecksOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		b,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference)SetChecks(val *[]*string) {
	if err := j.validateSetChecksParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"checks",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference)SetContexts(val *[]*string) {
	if err := j.validateSetContextsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contexts",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference)SetIncludeAdmins(val interface{}) {
	if err := j.validateSetIncludeAdminsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeAdmins",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference)SetInternalValue(val *BranchProtectionV3RequiredStatusChecks) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference)SetStrict(val interface{}) {
	if err := j.validateSetStrictParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"strict",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ResetChecks() {
	_jsii_.InvokeVoid(
		b,
		"resetChecks",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ResetContexts() {
	_jsii_.InvokeVoid(
		b,
		"resetContexts",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ResetIncludeAdmins() {
	_jsii_.InvokeVoid(
		b,
		"resetIncludeAdmins",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ResetStrict() {
	_jsii_.InvokeVoid(
		b,
		"resetStrict",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (b *jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

