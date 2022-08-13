// Prebuilt github Provider for Terraform CDK (cdktf)
package github


type RepositoryEnvironmentDeploymentBranchPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#custom_branch_policies RepositoryEnvironment#custom_branch_policies}.
	CustomBranchPolicies interface{} `field:"required" json:"customBranchPolicies" yaml:"customBranchPolicies"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#protected_branches RepositoryEnvironment#protected_branches}.
	ProtectedBranches interface{} `field:"required" json:"protectedBranches" yaml:"protectedBranches"`
}

