// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package repositorycollaborators

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RepositoryCollaboratorsTeamList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (r *jsiiProxy_RepositoryCollaboratorsTeamList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RepositoryCollaboratorsTeamList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RepositoryCollaboratorsTeamList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RepositoryCollaboratorsTeamList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RepositoryCollaboratorsTeamList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RepositoryCollaboratorsTeamList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRepositoryCollaboratorsTeamListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

