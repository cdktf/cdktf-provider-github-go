// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repository


type RepositoryPagesSource struct {
	// The repository branch used to publish the site's source files. (i.e. 'main' or 'gh-pages').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository#branch Repository#branch}
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// The repository directory from which the site publishes (Default: '/').
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.38.0/docs/resources/repository#path Repository#path}
	Path *string `field:"optional" json:"path" yaml:"path"`
}

