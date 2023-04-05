package repositorycollaborators


type RepositoryCollaboratorsUser struct {
	// (Required) The user to add to the repository as a collaborator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_collaborators#username RepositoryCollaborators#username}
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_collaborators#permission RepositoryCollaborators#permission}.
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

