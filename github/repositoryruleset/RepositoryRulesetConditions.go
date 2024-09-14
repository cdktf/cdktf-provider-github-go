// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetConditions struct {
	// ref_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/repository_ruleset#ref_name RepositoryRuleset#ref_name}
	RefName *RepositoryRulesetConditionsRefName `field:"required" json:"refName" yaml:"refName"`
}

