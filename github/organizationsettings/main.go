// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationsettings

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-github.organizationSettings.OrganizationSettings",
		reflect.TypeOf((*OrganizationSettings)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurityEnabledForNewRepositories", GoGetter: "AdvancedSecurityEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSecurityEnabledForNewRepositoriesInput", GoGetter: "AdvancedSecurityEnabledForNewRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "billingEmail", GoGetter: "BillingEmail"},
			_jsii_.MemberProperty{JsiiProperty: "billingEmailInput", GoGetter: "BillingEmailInput"},
			_jsii_.MemberProperty{JsiiProperty: "blog", GoGetter: "Blog"},
			_jsii_.MemberProperty{JsiiProperty: "blogInput", GoGetter: "BlogInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "company", GoGetter: "Company"},
			_jsii_.MemberProperty{JsiiProperty: "companyInput", GoGetter: "CompanyInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRepositoryPermission", GoGetter: "DefaultRepositoryPermission"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRepositoryPermissionInput", GoGetter: "DefaultRepositoryPermissionInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependabotAlertsEnabledForNewRepositories", GoGetter: "DependabotAlertsEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "dependabotAlertsEnabledForNewRepositoriesInput", GoGetter: "DependabotAlertsEnabledForNewRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependabotSecurityUpdatesEnabledForNewRepositories", GoGetter: "DependabotSecurityUpdatesEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "dependabotSecurityUpdatesEnabledForNewRepositoriesInput", GoGetter: "DependabotSecurityUpdatesEnabledForNewRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyGraphEnabledForNewRepositories", GoGetter: "DependencyGraphEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "dependencyGraphEnabledForNewRepositoriesInput", GoGetter: "DependencyGraphEnabledForNewRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "email", GoGetter: "Email"},
			_jsii_.MemberProperty{JsiiProperty: "emailInput", GoGetter: "EmailInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hasOrganizationProjects", GoGetter: "HasOrganizationProjects"},
			_jsii_.MemberProperty{JsiiProperty: "hasOrganizationProjectsInput", GoGetter: "HasOrganizationProjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "hasRepositoryProjects", GoGetter: "HasRepositoryProjects"},
			_jsii_.MemberProperty{JsiiProperty: "hasRepositoryProjectsInput", GoGetter: "HasRepositoryProjectsInput"},
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreateInternalRepositories", GoGetter: "MembersCanCreateInternalRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreateInternalRepositoriesInput", GoGetter: "MembersCanCreateInternalRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePages", GoGetter: "MembersCanCreatePages"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePagesInput", GoGetter: "MembersCanCreatePagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePrivatePages", GoGetter: "MembersCanCreatePrivatePages"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePrivatePagesInput", GoGetter: "MembersCanCreatePrivatePagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePrivateRepositories", GoGetter: "MembersCanCreatePrivateRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePrivateRepositoriesInput", GoGetter: "MembersCanCreatePrivateRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePublicPages", GoGetter: "MembersCanCreatePublicPages"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePublicPagesInput", GoGetter: "MembersCanCreatePublicPagesInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePublicRepositories", GoGetter: "MembersCanCreatePublicRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreatePublicRepositoriesInput", GoGetter: "MembersCanCreatePublicRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreateRepositories", GoGetter: "MembersCanCreateRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanCreateRepositoriesInput", GoGetter: "MembersCanCreateRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanForkPrivateRepositories", GoGetter: "MembersCanForkPrivateRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "membersCanForkPrivateRepositoriesInput", GoGetter: "MembersCanForkPrivateRepositoriesInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdvancedSecurityEnabledForNewRepositories", GoMethod: "ResetAdvancedSecurityEnabledForNewRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetBlog", GoMethod: "ResetBlog"},
			_jsii_.MemberMethod{JsiiMethod: "resetCompany", GoMethod: "ResetCompany"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultRepositoryPermission", GoMethod: "ResetDefaultRepositoryPermission"},
			_jsii_.MemberMethod{JsiiMethod: "resetDependabotAlertsEnabledForNewRepositories", GoMethod: "ResetDependabotAlertsEnabledForNewRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetDependabotSecurityUpdatesEnabledForNewRepositories", GoMethod: "ResetDependabotSecurityUpdatesEnabledForNewRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetDependencyGraphEnabledForNewRepositories", GoMethod: "ResetDependencyGraphEnabledForNewRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmail", GoMethod: "ResetEmail"},
			_jsii_.MemberMethod{JsiiMethod: "resetHasOrganizationProjects", GoMethod: "ResetHasOrganizationProjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetHasRepositoryProjects", GoMethod: "ResetHasRepositoryProjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocation", GoMethod: "ResetLocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembersCanCreateInternalRepositories", GoMethod: "ResetMembersCanCreateInternalRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembersCanCreatePages", GoMethod: "ResetMembersCanCreatePages"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembersCanCreatePrivatePages", GoMethod: "ResetMembersCanCreatePrivatePages"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembersCanCreatePrivateRepositories", GoMethod: "ResetMembersCanCreatePrivateRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembersCanCreatePublicPages", GoMethod: "ResetMembersCanCreatePublicPages"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembersCanCreatePublicRepositories", GoMethod: "ResetMembersCanCreatePublicRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembersCanCreateRepositories", GoMethod: "ResetMembersCanCreateRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetMembersCanForkPrivateRepositories", GoMethod: "ResetMembersCanForkPrivateRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretScanningEnabledForNewRepositories", GoMethod: "ResetSecretScanningEnabledForNewRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretScanningPushProtectionEnabledForNewRepositories", GoMethod: "ResetSecretScanningPushProtectionEnabledForNewRepositories"},
			_jsii_.MemberMethod{JsiiMethod: "resetTwitterUsername", GoMethod: "ResetTwitterUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebCommitSignoffRequired", GoMethod: "ResetWebCommitSignoffRequired"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanningEnabledForNewRepositories", GoGetter: "SecretScanningEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanningEnabledForNewRepositoriesInput", GoGetter: "SecretScanningEnabledForNewRepositoriesInput"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanningPushProtectionEnabledForNewRepositories", GoGetter: "SecretScanningPushProtectionEnabledForNewRepositories"},
			_jsii_.MemberProperty{JsiiProperty: "secretScanningPushProtectionEnabledForNewRepositoriesInput", GoGetter: "SecretScanningPushProtectionEnabledForNewRepositoriesInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "twitterUsername", GoGetter: "TwitterUsername"},
			_jsii_.MemberProperty{JsiiProperty: "twitterUsernameInput", GoGetter: "TwitterUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "webCommitSignoffRequired", GoGetter: "WebCommitSignoffRequired"},
			_jsii_.MemberProperty{JsiiProperty: "webCommitSignoffRequiredInput", GoGetter: "WebCommitSignoffRequiredInput"},
		},
		func() interface{} {
			j := jsiiProxy_OrganizationSettings{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.organizationSettings.OrganizationSettingsConfig",
		reflect.TypeOf((*OrganizationSettingsConfig)(nil)).Elem(),
	)
}
