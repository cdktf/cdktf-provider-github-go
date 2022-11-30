package repository


type RepositorySecurityAndAnalysisAdvancedSecurity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#status Repository#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

