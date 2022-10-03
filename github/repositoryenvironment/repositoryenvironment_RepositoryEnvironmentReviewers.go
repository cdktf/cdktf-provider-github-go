package repositoryenvironment


type RepositoryEnvironmentReviewers struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#teams RepositoryEnvironment#teams}.
	Teams *[]*float64 `field:"optional" json:"teams" yaml:"teams"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#users RepositoryEnvironment#users}.
	Users *[]*float64 `field:"optional" json:"users" yaml:"users"`
}

