// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesFilePathRestriction struct {
	// The file paths that are restricted from being pushed to the commit graph.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/repository_ruleset#restricted_file_paths RepositoryRuleset#restricted_file_paths}
	RestrictedFilePaths *[]*string `field:"required" json:"restrictedFilePaths" yaml:"restrictedFilePaths"`
}

