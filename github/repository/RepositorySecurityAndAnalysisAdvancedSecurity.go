package repository


type RepositorySecurityAndAnalysisAdvancedSecurity struct {
	// Set to 'enabled' to enable advanced security features on the repository. Can be 'enabled' or 'disabled'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#status Repository#status}
	Status *string `field:"required" json:"status" yaml:"status"`
}

