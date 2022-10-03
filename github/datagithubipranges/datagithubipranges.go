package datagithubipranges

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-github.dataGithubIpRanges.DataGithubIpRanges",
		reflect.TypeOf((*DataGithubIpRanges)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "actions", GoGetter: "Actions"},
			_jsii_.MemberProperty{JsiiProperty: "actionsIpv4", GoGetter: "ActionsIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "actionsIpv6", GoGetter: "ActionsIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "api", GoGetter: "Api"},
			_jsii_.MemberProperty{JsiiProperty: "apiIpv4", GoGetter: "ApiIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "apiIpv6", GoGetter: "ApiIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependabot", GoGetter: "Dependabot"},
			_jsii_.MemberProperty{JsiiProperty: "dependabotIpv4", GoGetter: "DependabotIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "dependabotIpv6", GoGetter: "DependabotIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "git", GoGetter: "Git"},
			_jsii_.MemberProperty{JsiiProperty: "gitIpv4", GoGetter: "GitIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "gitIpv6", GoGetter: "GitIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "hooks", GoGetter: "Hooks"},
			_jsii_.MemberProperty{JsiiProperty: "hooksIpv4", GoGetter: "HooksIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "hooksIpv6", GoGetter: "HooksIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "importer", GoGetter: "Importer"},
			_jsii_.MemberProperty{JsiiProperty: "importerIpv4", GoGetter: "ImporterIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "importerIpv6", GoGetter: "ImporterIpv6"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "pages", GoGetter: "Pages"},
			_jsii_.MemberProperty{JsiiProperty: "pagesIpv4", GoGetter: "PagesIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "pagesIpv6", GoGetter: "PagesIpv6"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "web", GoGetter: "Web"},
			_jsii_.MemberProperty{JsiiProperty: "webIpv4", GoGetter: "WebIpv4"},
			_jsii_.MemberProperty{JsiiProperty: "webIpv6", GoGetter: "WebIpv6"},
		},
		func() interface{} {
			j := jsiiProxy_DataGithubIpRanges{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-github.dataGithubIpRanges.DataGithubIpRangesConfig",
		reflect.TypeOf((*DataGithubIpRangesConfig)(nil)).Elem(),
	)
}
