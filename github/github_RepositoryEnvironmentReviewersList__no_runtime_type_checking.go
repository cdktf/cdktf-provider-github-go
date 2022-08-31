//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt github Provider for Terraform CDK (cdktf)
package github

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_RepositoryEnvironmentReviewersList) validateGetParameters(index *float64) error {
	return nil
}

func (r *jsiiProxy_RepositoryEnvironmentReviewersList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_RepositoryEnvironmentReviewersList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_RepositoryEnvironmentReviewersList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_RepositoryEnvironmentReviewersList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_RepositoryEnvironmentReviewersList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewRepositoryEnvironmentReviewersListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

