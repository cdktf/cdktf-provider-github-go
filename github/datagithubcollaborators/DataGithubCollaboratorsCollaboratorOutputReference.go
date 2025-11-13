// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagithubcollaborators

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-github-go/github/v15/jsii"

	"github.com/cdktf/cdktf-provider-github-go/github/v15/datagithubcollaborators/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubCollaboratorsCollaboratorOutputReference interface {
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
	EventsUrl() *string
	FollowersUrl() *string
	FollowingUrl() *string
	// Experimental.
	Fqn() *string
	GistsUrl() *string
	HtmlUrl() *string
	Id() *float64
	InternalValue() *DataGithubCollaboratorsCollaborator
	SetInternalValue(val *DataGithubCollaboratorsCollaborator)
	Login() *string
	OrganizationsUrl() *string
	Permission() *string
	ReceivedEventsUrl() *string
	ReposUrl() *string
	SiteAdmin() cdktf.IResolvable
	StarredUrl() *string
	SubscriptionsUrl() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Type() *string
	Url() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataGithubCollaboratorsCollaboratorOutputReference
type jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) EventsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"eventsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) FollowersUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"followersUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) FollowingUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"followingUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GistsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gistsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) HtmlUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"htmlUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) Id() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) InternalValue() *DataGithubCollaboratorsCollaborator {
	var returns *DataGithubCollaboratorsCollaborator
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) Login() *string {
	var returns *string
	_jsii_.Get(
		j,
		"login",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) OrganizationsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) Permission() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permission",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) ReceivedEventsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"receivedEventsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) ReposUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"reposUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) SiteAdmin() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"siteAdmin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) StarredUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"starredUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) SubscriptionsUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionsUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}


func NewDataGithubCollaboratorsCollaboratorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGithubCollaboratorsCollaboratorOutputReference {
	_init_.Initialize()

	if err := validateNewDataGithubCollaboratorsCollaboratorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.dataGithubCollaborators.DataGithubCollaboratorsCollaboratorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGithubCollaboratorsCollaboratorOutputReference_Override(d DataGithubCollaboratorsCollaboratorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.dataGithubCollaborators.DataGithubCollaboratorsCollaboratorOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference)SetInternalValue(val *DataGithubCollaboratorsCollaborator) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubCollaboratorsCollaboratorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

