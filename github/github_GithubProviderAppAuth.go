// Prebuilt github Provider for Terraform CDK (cdktf)
package github


type GithubProviderAppAuth struct {
	// The GitHub App ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github#id GithubProvider#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The GitHub App installation instance ID.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github#installation_id GithubProvider#installation_id}
	InstallationId *string `field:"required" json:"installationId" yaml:"installationId"`
	// The GitHub App PEM file contents.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github#pem_file GithubProvider#pem_file}
	PemFile *string `field:"required" json:"pemFile" yaml:"pemFile"`
}

