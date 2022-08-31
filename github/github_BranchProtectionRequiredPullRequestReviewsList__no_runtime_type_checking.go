//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt github Provider for Terraform CDK (cdktf)
package github

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) validateGetParameters(index *float64) error {
	return nil
}

func (b *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_BranchProtectionRequiredPullRequestReviewsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewBranchProtectionRequiredPullRequestReviewsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

