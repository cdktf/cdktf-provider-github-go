// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesRequiredStatusChecks struct {
	// required_check block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_ruleset#required_check RepositoryRuleset#required_check}
	RequiredCheck interface{} `field:"required" json:"requiredCheck" yaml:"requiredCheck"`
	// Allow repositories and branches to be created if a check would otherwise prohibit it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_ruleset#do_not_enforce_on_create RepositoryRuleset#do_not_enforce_on_create}
	DoNotEnforceOnCreate interface{} `field:"optional" json:"doNotEnforceOnCreate" yaml:"doNotEnforceOnCreate"`
	// Whether pull requests targeting a matching branch must be tested with the latest code.
	//
	// This setting will not take effect unless at least one status check is enabled. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.5.0/docs/resources/repository_ruleset#strict_required_status_checks_policy RepositoryRuleset#strict_required_status_checks_policy}
	StrictRequiredStatusChecksPolicy interface{} `field:"optional" json:"strictRequiredStatusChecksPolicy" yaml:"strictRequiredStatusChecksPolicy"`
}

