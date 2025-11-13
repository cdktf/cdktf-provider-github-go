// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package issuelabels

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_IssueLabelsLabelList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (i *jsiiProxy_IssueLabelsLabelList) validateGetParameters(index *float64) error {
	return nil
}

func (i *jsiiProxy_IssueLabelsLabelList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_IssueLabelsLabelList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_IssueLabelsLabelList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_IssueLabelsLabelList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_IssueLabelsLabelList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewIssueLabelsLabelListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

