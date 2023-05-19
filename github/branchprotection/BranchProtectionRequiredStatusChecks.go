package branchprotection


type BranchProtectionRequiredStatusChecks struct {
	// The list of status checks to require in order to merge into this branch.
	//
	// No status checks are required by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.1/docs/resources/branch_protection#contexts BranchProtection#contexts}
	Contexts *[]*string `field:"optional" json:"contexts" yaml:"contexts"`
	// Require branches to be up to date before merging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.25.1/docs/resources/branch_protection#strict BranchProtection#strict}
	Strict interface{} `field:"optional" json:"strict" yaml:"strict"`
}

