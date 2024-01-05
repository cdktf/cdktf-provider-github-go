// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package branchprotection

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-github.branchProtection.BranchProtection",
		reflect.TypeOf((*BranchProtection)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowsDeletions", GoGetter: "AllowsDeletions"},
			_jsii_.MemberProperty{JsiiProperty: "allowsDeletionsInput", GoGetter: "AllowsDeletionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowsForcePushes", GoGetter: "AllowsForcePushes"},
			_jsii_.MemberProperty{JsiiProperty: "allowsForcePushesInput", GoGetter: "AllowsForcePushesInput"},
			_jsii_.MemberProperty{JsiiProperty: "blocksCreations", GoGetter: "BlocksCreations"},
			_jsii_.MemberProperty{JsiiProperty: "blocksCreationsInput", GoGetter: "BlocksCreationsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enforceAdmins", GoGetter: "EnforceAdmins"},
			_jsii_.MemberProperty{JsiiProperty: "enforceAdminsInput", GoGetter: "EnforceAdminsInput"},
			_jsii_.MemberProperty{JsiiProperty: "forcePushBypassers", GoGetter: "ForcePushBypassers"},
			_jsii_.MemberProperty{JsiiProperty: "forcePushBypassersInput", GoGetter: "ForcePushBypassersInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "lockBranch", GoGetter: "LockBranch"},
			_jsii_.MemberProperty{JsiiProperty: "lockBranchInput", GoGetter: "LockBranchInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pattern", GoGetter: "Pattern"},
			_jsii_.MemberProperty{JsiiProperty: "patternInput", GoGetter: "PatternInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "pushRestrictions", GoGetter: "PushRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "pushRestrictionsInput", GoGetter: "PushRestrictionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putRequiredPullRequestReviews", GoMethod: "PutRequiredPullRequestReviews"},
			_jsii_.MemberMethod{JsiiMethod: "putRequiredStatusChecks", GoMethod: "PutRequiredStatusChecks"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryId", GoGetter: "RepositoryId"},
			_jsii_.MemberProperty{JsiiProperty: "repositoryIdInput", GoGetter: "RepositoryIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireConversationResolution", GoGetter: "RequireConversationResolution"},
			_jsii_.MemberProperty{JsiiProperty: "requireConversationResolutionInput", GoGetter: "RequireConversationResolutionInput"},
			_jsii_.MemberProperty{JsiiProperty: "requiredLinearHistory", GoGetter: "RequiredLinearHistory"},
			_jsii_.MemberProperty{JsiiProperty: "requiredLinearHistoryInput", GoGetter: "RequiredLinearHistoryInput"},
			_jsii_.MemberProperty{JsiiProperty: "requiredPullRequestReviews", GoGetter: "RequiredPullRequestReviews"},
			_jsii_.MemberProperty{JsiiProperty: "requiredPullRequestReviewsInput", GoGetter: "RequiredPullRequestReviewsInput"},
			_jsii_.MemberProperty{JsiiProperty: "requiredStatusChecks", GoGetter: "RequiredStatusChecks"},
			_jsii_.MemberProperty{JsiiProperty: "requiredStatusChecksInput", GoGetter: "RequiredStatusChecksInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireSignedCommits", GoGetter: "RequireSignedCommits"},
			_jsii_.MemberProperty{JsiiProperty: "requireSignedCommitsInput", GoGetter: "RequireSignedCommitsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowsDeletions", GoMethod: "ResetAllowsDeletions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowsForcePushes", GoMethod: "ResetAllowsForcePushes"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlocksCreations", GoMethod: "ResetBlocksCreations"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnforceAdmins", GoMethod: "ResetEnforceAdmins"},
			_jsii_.MemberMethod{JsiiMethod: "resetForcePushBypassers", GoMethod: "ResetForcePushBypassers"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLockBranch", GoMethod: "ResetLockBranch"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPushRestrictions", GoMethod: "ResetPushRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireConversationResolution", GoMethod: "ResetRequireConversationResolution"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiredLinearHistory", GoMethod: "ResetRequiredLinearHistory"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiredPullRequestReviews", GoMethod: "ResetRequiredPullRequestReviews"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiredStatusChecks", GoMethod: "ResetRequiredStatusChecks"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireSignedCommits", GoMethod: "ResetRequireSignedCommits"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_BranchProtection{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.branchProtection.BranchProtectionConfig",
		reflect.TypeOf((*BranchProtectionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.branchProtection.BranchProtectionRequiredPullRequestReviews",
		reflect.TypeOf((*BranchProtectionRequiredPullRequestReviews)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.branchProtection.BranchProtectionRequiredPullRequestReviewsList",
		reflect.TypeOf((*BranchProtectionRequiredPullRequestReviewsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_BranchProtectionRequiredPullRequestReviewsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.branchProtection.BranchProtectionRequiredPullRequestReviewsOutputReference",
		reflect.TypeOf((*BranchProtectionRequiredPullRequestReviewsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dismissalRestrictions", GoGetter: "DismissalRestrictions"},
			_jsii_.MemberProperty{JsiiProperty: "dismissalRestrictionsInput", GoGetter: "DismissalRestrictionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "pullRequestBypassers", GoGetter: "PullRequestBypassers"},
			_jsii_.MemberProperty{JsiiProperty: "pullRequestBypassersInput", GoGetter: "PullRequestBypassersInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireCodeOwnerReviews", GoGetter: "RequireCodeOwnerReviews"},
			_jsii_.MemberProperty{JsiiProperty: "requireCodeOwnerReviewsInput", GoGetter: "RequireCodeOwnerReviewsInput"},
			_jsii_.MemberProperty{JsiiProperty: "requiredApprovingReviewCount", GoGetter: "RequiredApprovingReviewCount"},
			_jsii_.MemberProperty{JsiiProperty: "requiredApprovingReviewCountInput", GoGetter: "RequiredApprovingReviewCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "requireLastPushApproval", GoGetter: "RequireLastPushApproval"},
			_jsii_.MemberProperty{JsiiProperty: "requireLastPushApprovalInput", GoGetter: "RequireLastPushApprovalInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDismissalRestrictions", GoMethod: "ResetDismissalRestrictions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDismissStaleReviews", GoMethod: "ResetDismissStaleReviews"},
			_jsii_.MemberMethod{JsiiMethod: "resetPullRequestBypassers", GoMethod: "ResetPullRequestBypassers"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireCodeOwnerReviews", GoMethod: "ResetRequireCodeOwnerReviews"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequiredApprovingReviewCount", GoMethod: "ResetRequiredApprovingReviewCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequireLastPushApproval", GoMethod: "ResetRequireLastPushApproval"},
			_jsii_.MemberMethod{JsiiMethod: "resetRestrictDismissals", GoMethod: "ResetRestrictDismissals"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "restrictDismissals", GoGetter: "RestrictDismissals"},
			_jsii_.MemberProperty{JsiiProperty: "restrictDismissalsInput", GoGetter: "RestrictDismissalsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BranchProtectionRequiredPullRequestReviewsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.branchProtection.BranchProtectionRequiredStatusChecks",
		reflect.TypeOf((*BranchProtectionRequiredStatusChecks)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.branchProtection.BranchProtectionRequiredStatusChecksList",
		reflect.TypeOf((*BranchProtectionRequiredStatusChecksList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_BranchProtectionRequiredStatusChecksList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-github.branchProtection.BranchProtectionRequiredStatusChecksOutputReference",
		reflect.TypeOf((*BranchProtectionRequiredStatusChecksOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetContexts", GoMethod: "ResetContexts"},
			_jsii_.MemberMethod{JsiiMethod: "resetStrict", GoMethod: "ResetStrict"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "strict", GoGetter: "Strict"},
			_jsii_.MemberProperty{JsiiProperty: "strictInput", GoGetter: "StrictInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_BranchProtectionRequiredStatusChecksOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
