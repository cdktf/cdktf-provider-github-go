// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesMaxFilePathLength struct {
	// The maximum allowed length of a file path.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.3/docs/resources/organization_ruleset#max_file_path_length OrganizationRuleset#max_file_path_length}
	MaxFilePathLength *float64 `field:"required" json:"maxFilePathLength" yaml:"maxFilePathLength"`
}

