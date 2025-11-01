// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesMaxFileSize struct {
	// The maximum allowed size of a file in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#max_file_size RepositoryRuleset#max_file_size}
	MaxFileSize *float64 `field:"required" json:"maxFileSize" yaml:"maxFileSize"`
}

