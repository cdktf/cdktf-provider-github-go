package repository


type RepositoryPages struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.23.0/docs/resources/repository#source Repository#source}
	Source *RepositoryPagesSource `field:"required" json:"source" yaml:"source"`
	// The custom domain for the repository. This can only be set after the repository has been created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.23.0/docs/resources/repository#cname Repository#cname}
	Cname *string `field:"optional" json:"cname" yaml:"cname"`
}

