package repository


type RepositoryPages struct {
	// source block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#source Repository#source}
	Source *RepositoryPagesSource `field:"required" json:"source" yaml:"source"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#cname Repository#cname}.
	Cname *string `field:"optional" json:"cname" yaml:"cname"`
}

