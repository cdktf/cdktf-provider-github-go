// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesMaxFilePathLength struct {
	// The maximum allowed length of a file path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/repository_ruleset#max_file_path_length RepositoryRuleset#max_file_path_length}
	MaxFilePathLength *float64 `field:"required" json:"maxFilePathLength" yaml:"maxFilePathLength"`
}

