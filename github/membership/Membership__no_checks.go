// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package membership

// Building without runtime type checking enabled, so all the below just return nil

func (m *jsiiProxy_Membership) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (m *jsiiProxy_Membership) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateImportFromParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (m *jsiiProxy_Membership) validateMoveToIdParameters(id *string) error {
	return nil
}

func (m *jsiiProxy_Membership) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateMembership_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateMembership_IsConstructParameters(x interface{}) error {
	return nil
}

func validateMembership_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateMembership_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Membership) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Membership) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Membership) validateSetDowngradeOnDestroyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Membership) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Membership) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Membership) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Membership) validateSetRoleParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Membership) validateSetUsernameParameters(val *string) error {
	return nil
}

func validateNewMembershipParameters(scope constructs.Construct, id *string, config *MembershipConfig) error {
	return nil
}

