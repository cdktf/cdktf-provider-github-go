package repository


type RepositoryTemplate struct {
	// The GitHub organization or user the template repository is owned by.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#owner Repository#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// The name of the template repository.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#repository Repository#repository}
	Repository *string `field:"required" json:"repository" yaml:"repository"`
	// Whether the new repository should include all the branches from the template repository (defaults to 'false', which includes only the default branch from the template).
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#include_all_branches Repository#include_all_branches}
	IncludeAllBranches interface{} `field:"optional" json:"includeAllBranches" yaml:"includeAllBranches"`
}

