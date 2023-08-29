// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repositoryenvironment


type RepositoryEnvironmentReviewers struct {
	// Up to 6 IDs for teams who may review jobs that reference the environment.
	//
	// Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/repository_environment#teams RepositoryEnvironment#teams}
	Teams *[]*float64 `field:"optional" json:"teams" yaml:"teams"`
	// Up to 6 IDs for users who may review jobs that reference the environment.
	//
	// Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.33.0/docs/resources/repository_environment#users RepositoryEnvironment#users}
	Users *[]*float64 `field:"optional" json:"users" yaml:"users"`
}

