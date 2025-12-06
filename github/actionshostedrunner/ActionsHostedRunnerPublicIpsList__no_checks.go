// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package actionshostedrunner

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_ActionsHostedRunnerPublicIpsList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (a *jsiiProxy_ActionsHostedRunnerPublicIpsList) validateGetParameters(index *float64) error {
	return nil
}

func (a *jsiiProxy_ActionsHostedRunnerPublicIpsList) validateResolveParameters(context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_ActionsHostedRunnerPublicIpsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_ActionsHostedRunnerPublicIpsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_ActionsHostedRunnerPublicIpsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewActionsHostedRunnerPublicIpsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

