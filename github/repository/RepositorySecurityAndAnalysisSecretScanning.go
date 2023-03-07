package repository


type RepositorySecurityAndAnalysisSecretScanning struct {
	// Set to 'enabled' to enable secret scanning on the repository.
	//
	// Can be 'enabled' or 'disabled'. If set to 'enabled', the repository's visibility must be 'public' or 'security_and_analysis[0].advanced_security[0].status' must also be set to 'enabled'.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#status Repository#status}
	Status *string `field:"required" json:"status" yaml:"status"`
}
