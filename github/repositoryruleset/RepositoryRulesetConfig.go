// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryruleset

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryRulesetConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Possible values for Enforcement are `disabled`, `active`, `evaluate`. Note: `evaluate` is currently only supported for owners of type `organization`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository_ruleset#enforcement RepositoryRuleset#enforcement}
	Enforcement *string `field:"required" json:"enforcement" yaml:"enforcement"`
	// The name of the ruleset.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository_ruleset#name RepositoryRuleset#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository_ruleset#rules RepositoryRuleset#rules}
	Rules *RepositoryRulesetRules `field:"required" json:"rules" yaml:"rules"`
	// Possible values are `branch` and `tag`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository_ruleset#target RepositoryRuleset#target}
	Target *string `field:"required" json:"target" yaml:"target"`
	// bypass_actors block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository_ruleset#bypass_actors RepositoryRuleset#bypass_actors}
	BypassActors interface{} `field:"optional" json:"bypassActors" yaml:"bypassActors"`
	// conditions block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository_ruleset#conditions RepositoryRuleset#conditions}
	Conditions *RepositoryRulesetConditions `field:"optional" json:"conditions" yaml:"conditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository_ruleset#id RepositoryRuleset#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Name of the repository to apply rulset to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository_ruleset#repository RepositoryRuleset#repository}
	Repository *string `field:"optional" json:"repository" yaml:"repository"`
}

