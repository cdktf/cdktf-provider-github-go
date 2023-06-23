package repository


type RepositorySecurityAndAnalysisAdvancedSecurity struct {
	// Set to 'enabled' to enable advanced security features on the repository. Can be 'enabled' or 'disabled'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.28.1/docs/resources/repository#status Repository#status}
	Status *string `field:"required" json:"status" yaml:"status"`
}

