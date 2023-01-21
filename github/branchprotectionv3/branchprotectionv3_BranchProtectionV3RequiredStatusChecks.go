package branchprotectionv3


type BranchProtectionV3RequiredStatusChecks struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#checks BranchProtectionV3#checks}.
	Checks *[]*string `field:"optional" json:"checks" yaml:"checks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#contexts BranchProtectionV3#contexts}.
	Contexts *[]*string `field:"optional" json:"contexts" yaml:"contexts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#include_admins BranchProtectionV3#include_admins}.
	IncludeAdmins interface{} `field:"optional" json:"includeAdmins" yaml:"includeAdmins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection_v3#strict BranchProtectionV3#strict}.
	Strict interface{} `field:"optional" json:"strict" yaml:"strict"`
}

