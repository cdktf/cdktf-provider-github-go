// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package branchprotection

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

