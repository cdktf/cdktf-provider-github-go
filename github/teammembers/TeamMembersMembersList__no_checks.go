// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package teammembers

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TeamMembersMembersList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TeamMembersMembersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TeamMembersMembersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TeamMembersMembersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TeamMembersMembersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TeamMembersMembersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTeamMembersMembersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

