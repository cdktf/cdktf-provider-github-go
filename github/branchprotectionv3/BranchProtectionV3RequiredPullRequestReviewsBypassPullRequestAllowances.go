package branchprotectionv3


type BranchProtectionV3RequiredPullRequestReviewsBypassPullRequestAllowances struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/branch_protection_v3#apps BranchProtectionV3#apps}.
	Apps *[]*string `field:"optional" json:"apps" yaml:"apps"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/branch_protection_v3#teams BranchProtectionV3#teams}.
	Teams *[]*string `field:"optional" json:"teams" yaml:"teams"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.0/docs/resources/branch_protection_v3#users BranchProtectionV3#users}.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

