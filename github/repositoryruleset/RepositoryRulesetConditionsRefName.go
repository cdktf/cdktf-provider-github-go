// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetConditionsRefName struct {
	// Array of ref names or patterns to exclude. The condition will not pass if any of these patterns match.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.45.0/docs/resources/repository_ruleset#exclude RepositoryRuleset#exclude}
	Exclude *[]*string `field:"required" json:"exclude" yaml:"exclude"`
	// Array of ref names or patterns to include.
	//
	// One of these patterns must match for the condition to pass. Also accepts `~DEFAULT_BRANCH` to include the default branch or `~ALL` to include all branches.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.45.0/docs/resources/repository_ruleset#include RepositoryRuleset#include}
	Include *[]*string `field:"required" json:"include" yaml:"include"`
}

