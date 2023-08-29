// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package repositorycollaborators

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RepositoryCollaboratorsUserList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RepositoryCollaboratorsUserList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RepositoryCollaboratorsUserList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RepositoryCollaboratorsUserList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RepositoryCollaboratorsUserList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RepositoryCollaboratorsUserList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRepositoryCollaboratorsUserListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

