// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package actionsvariable

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ActionsVariable) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateImportFromParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateMoveToIdParameters(id *string) error {
	return nil
}

func (a *jsiiProxy_ActionsVariable) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateActionsVariable_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateActionsVariable_IsConstructParameters(x interface{}) error {
	return nil
}

func validateActionsVariable_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateActionsVariable_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_ActionsVariable) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ActionsVariable) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_ActionsVariable) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ActionsVariable) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_ActionsVariable) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_ActionsVariable) validateSetRepositoryParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ActionsVariable) validateSetValueParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ActionsVariable) validateSetVariableNameParameters(val *string) error {
	return nil
}

func validateNewActionsVariableParameters(scope constructs.Construct, id *string, config *ActionsVariableConfig) error {
	return nil
}

