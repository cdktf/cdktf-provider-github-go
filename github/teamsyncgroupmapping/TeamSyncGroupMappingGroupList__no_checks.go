//go:build no_runtime_type_checking

package teamsyncgroupmapping

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TeamSyncGroupMappingGroupList) validateGetParameters(index *float64) error {
	return nil
}

func (t *jsiiProxy_TeamSyncGroupMappingGroupList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_TeamSyncGroupMappingGroupList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_TeamSyncGroupMappingGroupList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_TeamSyncGroupMappingGroupList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_TeamSyncGroupMappingGroupList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewTeamSyncGroupMappingGroupListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}
