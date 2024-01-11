// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesRequiredStatusChecks struct {
	// required_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_ruleset#required_check RepositoryRuleset#required_check}
	RequiredCheck interface{} `field:"required" json:"requiredCheck" yaml:"requiredCheck"`
	// Whether pull requests targeting a matching branch must be tested with the latest code.
	//
	// This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository_ruleset#strict_required_status_checks_policy RepositoryRuleset#strict_required_status_checks_policy}
	StrictRequiredStatusChecksPolicy interface{} `field:"optional" json:"strictRequiredStatusChecksPolicy" yaml:"strictRequiredStatusChecksPolicy"`
}

