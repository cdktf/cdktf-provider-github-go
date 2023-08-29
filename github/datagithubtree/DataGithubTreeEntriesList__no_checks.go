// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datagithubtree

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGithubTreeEntriesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGithubTreeEntriesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGithubTreeEntriesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGithubTreeEntriesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGithubTreeEntriesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGithubTreeEntriesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

