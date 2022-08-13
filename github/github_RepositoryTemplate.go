// Prebuilt github Provider for Terraform CDK (cdktf)
package github


type RepositoryTemplate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#owner Repository#owner}.
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository#repository Repository#repository}.
	Repository *string `field:"required" json:"repository" yaml:"repository"`
}

