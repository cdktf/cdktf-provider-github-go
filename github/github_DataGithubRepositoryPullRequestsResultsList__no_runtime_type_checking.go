//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// Prebuilt github Provider for Terraform CDK (cdktf)
package github

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGithubRepositoryPullRequestsResultsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGithubRepositoryPullRequestsResultsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGithubRepositoryPullRequestsResultsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
