package repositoryenvironment


type RepositoryEnvironmentReviewers struct {
	// Up to 6 IDs for teams who may review jobs that reference the environment.
	//
	// Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#teams RepositoryEnvironment#teams}
	Teams *[]*float64 `field:"optional" json:"teams" yaml:"teams"`
	// Up to 6 IDs for users who may review jobs that reference the environment.
	//
	// Reviewers must have at least read access to the repository. Only one of the required reviewers needs to approve the job for it to proceed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#users RepositoryEnvironment#users}
	Users *[]*float64 `field:"optional" json:"users" yaml:"users"`
}

