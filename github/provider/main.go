// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-github.provider.GithubProvider",
		reflect.TypeOf((*GithubProvider)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "alias", GoGetter: "Alias"},
			_jsii_.MemberProperty{JsiiProperty: "aliasInput", GoGetter: "AliasInput"},
			_jsii_.MemberProperty{JsiiProperty: "appAuth", GoGetter: "AppAuth"},
			_jsii_.MemberProperty{JsiiProperty: "appAuthInput", GoGetter: "AppAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrl", GoGetter: "BaseUrl"},
			_jsii_.MemberProperty{JsiiProperty: "baseUrlInput", GoGetter: "BaseUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberProperty{JsiiProperty: "insecure", GoGetter: "Insecure"},
			_jsii_.MemberProperty{JsiiProperty: "insecureInput", GoGetter: "InsecureInput"},
			_jsii_.MemberProperty{JsiiProperty: "metaAttributes", GoGetter: "MetaAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "organization", GoGetter: "Organization"},
			_jsii_.MemberProperty{JsiiProperty: "organizationInput", GoGetter: "OrganizationInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "owner", GoGetter: "Owner"},
			_jsii_.MemberProperty{JsiiProperty: "ownerInput", GoGetter: "OwnerInput"},
			_jsii_.MemberProperty{JsiiProperty: "parallelRequests", GoGetter: "ParallelRequests"},
			_jsii_.MemberProperty{JsiiProperty: "parallelRequestsInput", GoGetter: "ParallelRequestsInput"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "readDelayMs", GoGetter: "ReadDelayMs"},
			_jsii_.MemberProperty{JsiiProperty: "readDelayMsInput", GoGetter: "ReadDelayMsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAlias", GoMethod: "ResetAlias"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppAuth", GoMethod: "ResetAppAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseUrl", GoMethod: "ResetBaseUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetInsecure", GoMethod: "ResetInsecure"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrganization", GoMethod: "ResetOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOwner", GoMethod: "ResetOwner"},
			_jsii_.MemberMethod{JsiiMethod: "resetParallelRequests", GoMethod: "ResetParallelRequests"},
			_jsii_.MemberMethod{JsiiMethod: "resetReadDelayMs", GoMethod: "ResetReadDelayMs"},
			_jsii_.MemberMethod{JsiiMethod: "resetToken", GoMethod: "ResetToken"},
			_jsii_.MemberMethod{JsiiMethod: "resetWriteDelayMs", GoMethod: "ResetWriteDelayMs"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformProviderSource", GoGetter: "TerraformProviderSource"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "token", GoGetter: "Token"},
			_jsii_.MemberProperty{JsiiProperty: "tokenInput", GoGetter: "TokenInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "writeDelayMs", GoGetter: "WriteDelayMs"},
			_jsii_.MemberProperty{JsiiProperty: "writeDelayMsInput", GoGetter: "WriteDelayMsInput"},
		},
		func() interface{} {
			j := jsiiProxy_GithubProvider{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformProvider)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.provider.GithubProviderAppAuth",
		reflect.TypeOf((*GithubProviderAppAuth)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.provider.GithubProviderConfig",
		reflect.TypeOf((*GithubProviderConfig)(nil)).Elem(),
	)
}
