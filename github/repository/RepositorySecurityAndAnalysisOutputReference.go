// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repository

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v13/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v13/repository/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositorySecurityAndAnalysisOutputReference interface {
	cdktf.ComplexObject
	AdvancedSecurity() RepositorySecurityAndAnalysisAdvancedSecurityOutputReference
	AdvancedSecurityInput() *RepositorySecurityAndAnalysisAdvancedSecurity
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
	InternalValue() *RepositorySecurityAndAnalysis
	SetInternalValue(val *RepositorySecurityAndAnalysis)
	SecretScanning() RepositorySecurityAndAnalysisSecretScanningOutputReference
	SecretScanningInput() *RepositorySecurityAndAnalysisSecretScanning
	SecretScanningPushProtection() RepositorySecurityAndAnalysisSecretScanningPushProtectionOutputReference
	SecretScanningPushProtectionInput() *RepositorySecurityAndAnalysisSecretScanningPushProtection
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
	PutAdvancedSecurity(value *RepositorySecurityAndAnalysisAdvancedSecurity)
	PutSecretScanning(value *RepositorySecurityAndAnalysisSecretScanning)
	PutSecretScanningPushProtection(value *RepositorySecurityAndAnalysisSecretScanningPushProtection)
	ResetAdvancedSecurity()
	ResetSecretScanning()
	ResetSecretScanningPushProtection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RepositorySecurityAndAnalysisOutputReference
type jsiiProxy_RepositorySecurityAndAnalysisOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) AdvancedSecurity() RepositorySecurityAndAnalysisAdvancedSecurityOutputReference {
	var returns RepositorySecurityAndAnalysisAdvancedSecurityOutputReference
	_jsii_.Get(
		j,
		"advancedSecurity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) AdvancedSecurityInput() *RepositorySecurityAndAnalysisAdvancedSecurity {
	var returns *RepositorySecurityAndAnalysisAdvancedSecurity
	_jsii_.Get(
		j,
		"advancedSecurityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) InternalValue() *RepositorySecurityAndAnalysis {
	var returns *RepositorySecurityAndAnalysis
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) SecretScanning() RepositorySecurityAndAnalysisSecretScanningOutputReference {
	var returns RepositorySecurityAndAnalysisSecretScanningOutputReference
	_jsii_.Get(
		j,
		"secretScanning",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) SecretScanningInput() *RepositorySecurityAndAnalysisSecretScanning {
	var returns *RepositorySecurityAndAnalysisSecretScanning
	_jsii_.Get(
		j,
		"secretScanningInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) SecretScanningPushProtection() RepositorySecurityAndAnalysisSecretScanningPushProtectionOutputReference {
	var returns RepositorySecurityAndAnalysisSecretScanningPushProtectionOutputReference
	_jsii_.Get(
		j,
		"secretScanningPushProtection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) SecretScanningPushProtectionInput() *RepositorySecurityAndAnalysisSecretScanningPushProtection {
	var returns *RepositorySecurityAndAnalysisSecretScanningPushProtection
	_jsii_.Get(
		j,
		"secretScanningPushProtectionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRepositorySecurityAndAnalysisOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) RepositorySecurityAndAnalysisOutputReference {
	_init_.Initialize()

	if err := validateNewRepositorySecurityAndAnalysisOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_RepositorySecurityAndAnalysisOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewRepositorySecurityAndAnalysisOutputReference_Override(r RepositorySecurityAndAnalysisOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.repository.RepositorySecurityAndAnalysisOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference)SetInternalValue(val *RepositorySecurityAndAnalysis) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RepositorySecurityAndAnalysisOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) PutAdvancedSecurity(value *RepositorySecurityAndAnalysisAdvancedSecurity) {
	if err := r.validatePutAdvancedSecurityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAdvancedSecurity",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) PutSecretScanning(value *RepositorySecurityAndAnalysisSecretScanning) {
	if err := r.validatePutSecretScanningParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSecretScanning",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) PutSecretScanningPushProtection(value *RepositorySecurityAndAnalysisSecretScanningPushProtection) {
	if err := r.validatePutSecretScanningPushProtectionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSecretScanningPushProtection",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) ResetAdvancedSecurity() {
	_jsii_.InvokeVoid(
		r,
		"resetAdvancedSecurity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) ResetSecretScanning() {
	_jsii_.InvokeVoid(
		r,
		"resetSecretScanning",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) ResetSecretScanningPushProtection() {
	_jsii_.InvokeVoid(
		r,
		"resetSecretScanningPushProtection",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_RepositorySecurityAndAnalysisOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

