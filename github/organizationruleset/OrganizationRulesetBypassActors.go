// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationruleset


type OrganizationRulesetBypassActors struct {
	// The ID of the actor that can bypass a ruleset.
	//
	// When `actor_type` is `OrganizationAdmin`, this should be set to `1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/organization_ruleset#actor_id OrganizationRuleset#actor_id}
	ActorId *float64 `field:"required" json:"actorId" yaml:"actorId"`
	// The type of actor that can bypass a ruleset. Can be one of: `RepositoryRole`, `Team`, `Integration`, `OrganizationAdmin`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/organization_ruleset#actor_type OrganizationRuleset#actor_type}
	ActorType *string `field:"required" json:"actorType" yaml:"actorType"`
	// When the specified actor can bypass the ruleset.
	//
	// pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.2.2/docs/resources/organization_ruleset#bypass_mode OrganizationRuleset#bypass_mode}
	BypassMode *string `field:"required" json:"bypassMode" yaml:"bypassMode"`
}

