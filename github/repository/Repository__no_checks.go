// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

//go:build no_runtime_type_checking

package repository

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Repository) validateAddMoveTargetParameters(moveTarget *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateAddOverrideParameters(path *string, value interface{}) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetAnyMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetBooleanAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetBooleanMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetNumberAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetNumberListAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetNumberMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetStringAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateGetStringMapAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateImportFromParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateInterpolationForAttributeParameters(terraformAttribute *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateMoveFromIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateMoveToParameters(moveTarget *string, index interface{}) error {
	return nil
}

func (r *jsiiProxy_Repository) validateMoveToIdParameters(id *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validateOverrideLogicalIdParameters(newLogicalId *string) error {
	return nil
}

func (r *jsiiProxy_Repository) validatePutPagesParameters(value *RepositoryPages) error {
	return nil
}

func (r *jsiiProxy_Repository) validatePutSecurityAndAnalysisParameters(value *RepositorySecurityAndAnalysis) error {
	return nil
}

func (r *jsiiProxy_Repository) validatePutTemplateParameters(value *RepositoryTemplate) error {
	return nil
}

func validateRepository_GenerateConfigForImportParameters(scope constructs.Construct, importToId *string, importFromId *string) error {
	return nil
}

func validateRepository_IsConstructParameters(x interface{}) error {
	return nil
}

func validateRepository_IsTerraformElementParameters(x interface{}) error {
	return nil
}

func validateRepository_IsTerraformResourceParameters(x interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetAllowAutoMergeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetAllowMergeCommitParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetAllowRebaseMergeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetAllowSquashMergeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetAllowUpdateBranchParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetArchivedParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetArchiveOnDestroyParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetAutoInitParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetConnectionParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetCountParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetDefaultBranchParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetDeleteBranchOnMergeParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetDescriptionParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetGitignoreTemplateParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetHasDiscussionsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetHasDownloadsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetHasIssuesParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetHasProjectsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetHasWikiParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetHomepageUrlParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetIdParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetIgnoreVulnerabilityAlertsDuringReadParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetIsTemplateParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetLicenseTemplateParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetLifecycleParameters(val *cdktf.TerraformResourceLifecycle) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetMergeCommitMessageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetMergeCommitTitleParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetNameParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetPrivateParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetProvisionersParameters(val *[]interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetSquashMergeCommitMessageParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetSquashMergeCommitTitleParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetTopicsParameters(val *[]*string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetVisibilityParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetVulnerabilityAlertsParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_Repository) validateSetWebCommitSignoffRequiredParameters(val interface{}) error {
	return nil
}

func validateNewRepositoryParameters(scope constructs.Construct, id *string, config *RepositoryConfig) error {
	return nil
}

