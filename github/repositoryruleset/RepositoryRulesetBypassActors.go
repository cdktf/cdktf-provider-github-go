// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset


type RepositoryRulesetBypassActors struct {
	// The type of actor that can bypass a ruleset. See https://docs.github.com/en/rest/repos/rules for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#actor_type RepositoryRuleset#actor_type}
	ActorType *string `field:"required" json:"actorType" yaml:"actorType"`
	// When the specified actor can bypass the ruleset.
	//
	// pull_request means that an actor can only bypass rules on pull requests. Can be one of: `always`, `pull_request`, `exempt`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#bypass_mode RepositoryRuleset#bypass_mode}
	BypassMode *string `field:"required" json:"bypassMode" yaml:"bypassMode"`
	// The ID of the actor that can bypass a ruleset.
	//
	// When `actor_type` is `OrganizationAdmin`, this should be set to `1`. Some resources such as DeployKey do not have an ID and this should be omitted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.5/docs/resources/repository_ruleset#actor_id RepositoryRuleset#actor_id}
	ActorId *float64 `field:"optional" json:"actorId" yaml:"actorId"`
}

