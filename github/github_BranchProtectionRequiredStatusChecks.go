// Prebuilt github Provider for Terraform CDK (cdktf)
package github


type BranchProtectionRequiredStatusChecks struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection#contexts BranchProtection#contexts}.
	Contexts *[]*string `field:"optional" json:"contexts" yaml:"contexts"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/branch_protection#strict BranchProtection#strict}.
	Strict interface{} `field:"optional" json:"strict" yaml:"strict"`
}

