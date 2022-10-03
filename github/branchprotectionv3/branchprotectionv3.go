package branchprotectionv3

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3",
		reflect.TypeOf((*BranchProtectionV3)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "branch", GoGetter: "Branch"},
			_jsii_.MemberProperty{JsiiProperty: "branchInput", GoGetter: "BranchInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enforceAdmins", GoGetter: "EnforceAdmins"},
			_jsii_.MemberProperty{JsiiProperty: "enforceAdminsInput", GoGetter: "EnforceAdminsInput"},
			_jsii_.MemberProperty{JsiiProperty: "etag", GoGetter: "Etag"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putRequiredPullRequestReviews", GoMethod: "PutRequiredPullRequestReviews"},
			_jsii_.MemberMethod{JsiiMethod: "putRequiredStatusChecks", GoMethod: "PutRequiredStatusChecks"},
			_jsii_.MemberMethod{JsiiMethod: "putRestrictions", GoMethod: "PutRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "repository", GoGetter: "Repository"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryInput", GoGetter: "RepositoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireConversationResolution", GoGetter: "RequireConversationResolution"},
			_jsii_.MemberProperty{JsiiProperty: "requireConversationResolutionInput", GoGetter: "RequireConversationResolutionInput"},
			_jsii_.MemberProperty{JsiiProperty: "requiredPullRequestReviews", GoGetter: "RequiredPullRequestReviews"},
			_jsii_.MemberProperty{JsiiProperty: "requiredPullRequestReviewsInput", GoGetter: "RequiredPullRequestReviewsInput"},
			_jsii_.MemberProperty{JsiiProperty: "requiredStatusChecks", GoGetter: "RequiredStatusChecks"},
			_jsii_.MemberProperty{JsiiProperty: "requiredStatusChecksInput", GoGetter: "RequiredStatusChecksInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireSignedCommits", GoGetter: "RequireSignedCommits"},
			_jsii_.MemberProperty{JsiiProperty: "requireSignedCommitsInput", GoGetter: "RequireSignedCommitsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnforceAdmins", GoMethod: "ResetEnforceAdmins"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireConversationResolution", GoMethod: "ResetRequireConversationResolution"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiredPullRequestReviews", GoMethod: "ResetRequiredPullRequestReviews"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiredStatusChecks", GoMethod: "ResetRequiredStatusChecks"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireSignedCommits", GoMethod: "ResetRequireSignedCommits"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestrictions", GoMethod: "ResetRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "restrictions", GoGetter: "Restrictions"},
			_jsii_.MemberProperty{JsiiProperty: "restrictionsInput", GoGetter: "RestrictionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_BranchProtectionV3{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3Config",
		reflect.TypeOf((*BranchProtectionV3Config)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredPullRequestReviews",
		reflect.TypeOf((*BranchProtectionV3RequiredPullRequestReviews)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredPullRequestReviewsOutputReference",
		reflect.TypeOf((*BranchProtectionV3RequiredPullRequestReviewsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dismissalTeams", GoGetter: "DismissalTeams"},
			_jsii_.MemberProperty{JsiiProperty: "dismissalTeamsInput", GoGetter: "DismissalTeamsInput"},
			_jsii_.MemberProperty{JsiiProperty: "dismissalUsers", GoGetter: "DismissalUsers"},
			_jsii_.MemberProperty{JsiiProperty: "dismissalUsersInput", GoGetter: "DismissalUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "dismissStaleReviews", GoGetter: "DismissStaleReviews"},
			_jsii_.MemberProperty{JsiiProperty: "dismissStaleReviewsInput", GoGetter: "DismissStaleReviewsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeAdmins", GoGetter: "IncludeAdmins"},
			_jsii_.MemberProperty{JsiiProperty: "includeAdminsInput", GoGetter: "IncludeAdminsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "requireCodeOwnerReviews", GoGetter: "RequireCodeOwnerReviews"},
			_jsii_.MemberProperty{JsiiProperty: "requireCodeOwnerReviewsInput", GoGetter: "RequireCodeOwnerReviewsInput"},
			_jsii_.MemberProperty{JsiiProperty: "requiredApprovingReviewCount", GoGetter: "RequiredApprovingReviewCount"},
			_jsii_.MemberProperty{JsiiProperty: "requiredApprovingReviewCountInput", GoGetter: "RequiredApprovingReviewCountInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDismissalTeams", GoMethod: "ResetDismissalTeams"},
			_jsii_.MemberMethod{JsiiMethod: "resetDismissalUsers", GoMethod: "ResetDismissalUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resetDismissStaleReviews", GoMethod: "ResetDismissStaleReviews"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeAdmins", GoMethod: "ResetIncludeAdmins"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireCodeOwnerReviews", GoMethod: "ResetRequireCodeOwnerReviews"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiredApprovingReviewCount", GoMethod: "ResetRequiredApprovingReviewCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BranchProtectionV3RequiredPullRequestReviewsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredStatusChecks",
		reflect.TypeOf((*BranchProtectionV3RequiredStatusChecks)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RequiredStatusChecksOutputReference",
		reflect.TypeOf((*BranchProtectionV3RequiredStatusChecksOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "contexts", GoGetter: "Contexts"},
			_jsii_.MemberProperty{JsiiProperty: "contextsInput", GoGetter: "ContextsInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "includeAdmins", GoGetter: "IncludeAdmins"},
			_jsii_.MemberProperty{JsiiProperty: "includeAdminsInput", GoGetter: "IncludeAdminsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetContexts", GoMethod: "ResetContexts"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeAdmins", GoMethod: "ResetIncludeAdmins"},
			_jsii_.MemberMethod{JsiiMethod: "resetStrict", GoMethod: "ResetStrict"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "strict", GoGetter: "Strict"},
			_jsii_.MemberProperty{JsiiProperty: "strictInput", GoGetter: "StrictInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BranchProtectionV3RequiredStatusChecksOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3Restrictions",
		reflect.TypeOf((*BranchProtectionV3Restrictions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.branchProtectionV3.BranchProtectionV3RestrictionsOutputReference",
		reflect.TypeOf((*BranchProtectionV3RestrictionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "apps", GoGetter: "Apps"},
			_jsii_.MemberProperty{JsiiProperty: "appsInput", GoGetter: "AppsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetApps", GoMethod: "ResetApps"},
			_jsii_.MemberMethod{JsiiMethod: "resetTeams", GoMethod: "ResetTeams"},
			_jsii_.MemberMethod{JsiiMethod: "resetUsers", GoMethod: "ResetUsers"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "teams", GoGetter: "Teams"},
			_jsii_.MemberProperty{JsiiProperty: "teamsInput", GoGetter: "TeamsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
			_jsii_.MemberProperty{JsiiProperty: "usersInput", GoGetter: "UsersInput"},
		},
		func() interface{} {
			j := jsiiProxy_BranchProtectionV3RestrictionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
