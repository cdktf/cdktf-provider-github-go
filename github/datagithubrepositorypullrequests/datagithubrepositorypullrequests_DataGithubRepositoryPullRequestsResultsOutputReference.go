package datagithubrepositorypullrequests

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-github-go/github/v3/jsii"

	"github.com/hashicorp/cdktf-provider-github-go/github/v3/datagithubrepositorypullrequests/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataGithubRepositoryPullRequestsResultsOutputReference interface {
	cdktf.ComplexObject
	BaseRef() *string
	BaseSha() *string
	Body() *string
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
	Draft() cdktf.IResolvable
	// Experimental.
	Fqn() *string
	HeadOwner() *string
	HeadRef() *string
	HeadRepository() *string
	HeadSha() *string
	InternalValue() *DataGithubRepositoryPullRequestsResults
	SetInternalValue(val *DataGithubRepositoryPullRequestsResults)
	Labels() *[]*string
	MaintainerCanModify() cdktf.IResolvable
	Number() *float64
	OpenedAt() *float64
	OpenedBy() *string
	State() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Title() *string
	UpdatedAt() *float64
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

// The jsii proxy struct for DataGithubRepositoryPullRequestsResultsOutputReference
type jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) BaseRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) BaseSha() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseSha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) Body() *string {
	var returns *string
	_jsii_.Get(
		j,
		"body",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) Draft() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"draft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) HeadOwner() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headOwner",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) HeadRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) HeadRepository() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headRepository",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) HeadSha() *string {
	var returns *string
	_jsii_.Get(
		j,
		"headSha",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) InternalValue() *DataGithubRepositoryPullRequestsResults {
	var returns *DataGithubRepositoryPullRequestsResults
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) Labels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) MaintainerCanModify() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"maintainerCanModify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) Number() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"number",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) OpenedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"openedAt",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) OpenedBy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"openedBy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) State() *string {
	var returns *string
	_jsii_.Get(
		j,
		"state",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) Title() *string {
	var returns *string
	_jsii_.Get(
		j,
		"title",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) UpdatedAt() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"updatedAt",
		&returns,
	)
	return returns
}


func NewDataGithubRepositoryPullRequestsResultsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataGithubRepositoryPullRequestsResultsOutputReference {
	_init_.Initialize()

	if err := validateNewDataGithubRepositoryPullRequestsResultsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-github.dataGithubRepositoryPullRequests.DataGithubRepositoryPullRequestsResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataGithubRepositoryPullRequestsResultsOutputReference_Override(d DataGithubRepositoryPullRequestsResultsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-github.dataGithubRepositoryPullRequests.DataGithubRepositoryPullRequestsResultsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference)SetInternalValue(val *DataGithubRepositoryPullRequestsResults) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

