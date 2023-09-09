// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package repositoryruleset

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RepositoryRulesetBypassActorsList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RepositoryRulesetBypassActorsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RepositoryRulesetBypassActorsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RepositoryRulesetBypassActorsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RepositoryRulesetBypassActorsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RepositoryRulesetBypassActorsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRepositoryRulesetBypassActorsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

