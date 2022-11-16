//go:build no_runtime_type_checking

package provider

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GithubProvider) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (g *jsiiProxy_GithubProvider) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func validateGithubProvider_IsConstructParameters(x interface{}) error {
	return nil
}

func validateGithubProvider_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateGithubProvider_IsTerraformProviderParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_GithubProvider) validateSetAppAuthParameters(val *GithubProviderAppAuth) error {
	return nil
}

func (j *jsiiProxy_GithubProvider) validateSetInsecureParameters(val interface{}) error {
	return nil
}

func validateNewGithubProviderParameters(scope constructs.Construct, id *string, config *GithubProviderConfig) error {
	return nil
}

