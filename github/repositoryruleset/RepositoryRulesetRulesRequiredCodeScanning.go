// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesRequiredCodeScanning struct {
	// required_code_scanning_tool block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.4.0/docs/resources/repository_ruleset#required_code_scanning_tool RepositoryRuleset#required_code_scanning_tool}
	RequiredCodeScanningTool interface{} `field:"required" json:"requiredCodeScanningTool" yaml:"requiredCodeScanningTool"`
}

