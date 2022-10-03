package repository


type RepositoryPagesSource struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#branch Repository#branch}.
	Branch *string `field:"required" json:"branch" yaml:"branch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#path Repository#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
}

