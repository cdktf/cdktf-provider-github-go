//go:build no_runtime_type_checking

package datagithubrelease

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DataGithubReleaseAssetsList) validateGetParameters(index *float64) error {
	return nil
}

func (d *jsiiProxy_DataGithubReleaseAssetsList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_DataGithubReleaseAssetsList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_DataGithubReleaseAssetsList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_DataGithubReleaseAssetsList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewDataGithubReleaseAssetsListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

