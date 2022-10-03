//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package branchprotection

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BranchProtectionRequiredStatusChecksList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BranchProtectionRequiredStatusChecksList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRequiredStatusChecksList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRequiredStatusChecksList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRequiredStatusChecksList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRequiredStatusChecksList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBranchProtectionRequiredStatusChecksListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

