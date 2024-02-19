// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package branchprotection

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BranchProtectionRestrictPushesList) validateAllWithMapKeyParameters(mapKeyAttributeName *string) error {
	return nil
}

func (b *jsiiProxy_BranchProtectionRestrictPushesList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BranchProtectionRestrictPushesList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRestrictPushesList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRestrictPushesList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRestrictPushesList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRestrictPushesList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBranchProtectionRestrictPushesListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

