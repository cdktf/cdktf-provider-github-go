// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datagithuborganization

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-github.dataGithubOrganization.DataGithubOrganization",
		reflect.TypeOf((*DataGithubOrganization)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurityEnabledForNewRepositories", GoGetter: "AdvancedSecurityEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRepositoryPermission", GoGetter: "DefaultRepositoryPermission"},
			_jsii_.MemberProperty{JsiiProperty: "dependabotAlertsEnabledForNewRepositories", GoGetter: "DependabotAlertsEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "dependabotSecurityUpdatesEnabledForNewRepositories", GoGetter: "DependabotSecurityUpdatesEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyGraphEnabledForNewRepositories", GoGetter: "DependencyGraphEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
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
			_jsii_.MemberProperty{JsiiProperty: "login", GoGetter: "Login"},
			_jsii_.MemberProperty{JsiiProperty: "members", GoGetter: "Members"},
			_jsii_.MemberProperty{JsiiProperty: "membersAllowedRepositoryCreationType", GoGetter: "MembersAllowedRepositoryCreationType"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreateInternalRepositories", GoGetter: "MembersCanCreateInternalRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePages", GoGetter: "MembersCanCreatePages"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePrivatePages", GoGetter: "MembersCanCreatePrivatePages"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePrivateRepositories", GoGetter: "MembersCanCreatePrivateRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePublicPages", GoGetter: "MembersCanCreatePublicPages"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePublicRepositories", GoGetter: "MembersCanCreatePublicRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreateRepositories", GoGetter: "MembersCanCreateRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanForkPrivateRepositories", GoGetter: "MembersCanForkPrivateRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "nodeId", GoGetter: "NodeId"},
			_jsii_.MemberProperty{JsiiProperty: "orgname", GoGetter: "Orgname"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "plan", GoGetter: "Plan"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "repositories", GoGetter: "Repositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanningEnabledForNewRepositories", GoGetter: "SecretScanningEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanningPushProtectionEnabledForNewRepositories", GoGetter: "SecretScanningPushProtectionEnabledForNewRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "twoFactorRequirementEnabled", GoGetter: "TwoFactorRequirementEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
			_jsii_.MemberProperty{JsiiProperty: "webCommitSignoffRequired", GoGetter: "WebCommitSignoffRequired"},
		},
		func() interface{} {
			j := jsiiProxy_DataGithubOrganization{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.dataGithubOrganization.DataGithubOrganizationConfig",
		reflect.TypeOf((*DataGithubOrganizationConfig)(nil)).Elem(),
	)
}
