package repositorycollaborators


type RepositoryCollaboratorsTeam struct {
	// Team ID or slug to add to the repository as a collaborator.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_collaborators#team_id RepositoryCollaborators#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_collaborators#permission RepositoryCollaborators#permission}.
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

