// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesCommitAuthorEmailPattern struct {
	// The operator to use for matching. Can be one of: `starts_with`, `ends_with`, `contains`, `regex`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.42.0/docs/resources/repository_ruleset#operator RepositoryRuleset#operator}
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// The pattern to match with.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.42.0/docs/resources/repository_ruleset#pattern RepositoryRuleset#pattern}
	Pattern *string `field:"required" json:"pattern" yaml:"pattern"`
	// How this rule will appear to users.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.42.0/docs/resources/repository_ruleset#name RepositoryRuleset#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// If true, the rule will fail if the pattern matches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.42.0/docs/resources/repository_ruleset#negate RepositoryRuleset#negate}
	Negate interface{} `field:"optional" json:"negate" yaml:"negate"`
}

