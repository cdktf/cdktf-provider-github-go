//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package datagithubrepository

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGithubRepositoryPagesSourceList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGithubRepositoryPagesSourceList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGithubRepositoryPagesSourceList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGithubRepositoryPagesSourceList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGithubRepositoryPagesSourceList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGithubRepositoryPagesSourceListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
