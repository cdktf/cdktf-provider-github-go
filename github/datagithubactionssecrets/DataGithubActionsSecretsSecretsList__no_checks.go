// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package datagithubactionssecrets

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGithubActionsSecretsSecretsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGithubActionsSecretsSecretsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGithubActionsSecretsSecretsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGithubActionsSecretsSecretsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGithubActionsSecretsSecretsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGithubActionsSecretsSecretsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

