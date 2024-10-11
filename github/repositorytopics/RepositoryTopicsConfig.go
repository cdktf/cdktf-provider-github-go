// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositorytopics

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RepositoryTopicsConfig struct {
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
	// The name of the repository. The name is not case sensitive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/repository_topics#repository RepositoryTopics#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// An array of topics to add to the repository.
	//
	// Pass one or more topics to replace the set of existing topics. Send an empty array ([]) to clear all topics from the repository. Note: Topic names cannot contain uppercase letters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/repository_topics#topics RepositoryTopics#topics}
	Topics *[]*string `field:"required" json:"topics" yaml:"topics"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.1/docs/resources/repository_topics#id RepositoryTopics#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

