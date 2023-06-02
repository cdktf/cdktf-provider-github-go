package repository


type RepositoryPages struct {
	// The type the page should be sourced.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/repository#build_type Repository#build_type}
	BuildType *string `field:"optional" json:"buildType" yaml:"buildType"`
	// The custom domain for the repository. This can only be set after the repository has been created.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/repository#cname Repository#cname}
	Cname *string `field:"optional" json:"cname" yaml:"cname"`
	// source block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.26.0/docs/resources/repository#source Repository#source}
	Source *RepositoryPagesSource `field:"optional" json:"source" yaml:"source"`
}

