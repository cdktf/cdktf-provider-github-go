// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesMaxFileSize struct {
	// The maximum allowed size of a file in bytes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.3/docs/resources/organization_ruleset#max_file_size OrganizationRuleset#max_file_size}
	MaxFileSize *float64 `field:"required" json:"maxFileSize" yaml:"maxFileSize"`
}

