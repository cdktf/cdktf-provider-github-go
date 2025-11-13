// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datagithubrepository

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGithubRepositoryPagesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (d *jsiiProxy_DataGithubRepositoryPagesList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGithubRepositoryPagesList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGithubRepositoryPagesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGithubRepositoryPagesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGithubRepositoryPagesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGithubRepositoryPagesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

