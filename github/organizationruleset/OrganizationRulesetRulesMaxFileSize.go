// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetRulesMaxFileSize struct {
	// The maximum allowed size of a file in megabytes (MB). Valid range is 1-100 MB.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/organization_ruleset#max_file_size OrganizationRuleset#max_file_size}
	MaxFileSize *float64 `field:"required" json:"maxFileSize" yaml:"maxFileSize"`
}

