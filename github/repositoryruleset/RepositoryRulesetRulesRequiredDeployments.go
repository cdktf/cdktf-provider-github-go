// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetRulesRequiredDeployments struct {
	// The environments that must be successfully deployed to before branches can be merged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.4.0/docs/resources/repository_ruleset#required_deployment_environments RepositoryRuleset#required_deployment_environments}
	RequiredDeploymentEnvironments *[]*string `field:"required" json:"requiredDeploymentEnvironments" yaml:"requiredDeploymentEnvironments"`
}

