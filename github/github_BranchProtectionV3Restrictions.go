// Prebuilt github Provider for Terraform CDK (cdktf)
package github


type BranchProtectionV3Restrictions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#apps BranchProtectionV3#apps}.
	Apps *[]*string `field:"optional" json:"apps" yaml:"apps"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#teams BranchProtectionV3#teams}.
	Teams *[]*string `field:"optional" json:"teams" yaml:"teams"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#users BranchProtectionV3#users}.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

