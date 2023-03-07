package repositoryenvironment


type RepositoryEnvironmentDeploymentBranchPolicy struct {
	// Whether only branches that match the specified name patterns can deploy to this environment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#custom_branch_policies RepositoryEnvironment#custom_branch_policies}
	CustomBranchPolicies interface{} `field:"required" json:"customBranchPolicies" yaml:"customBranchPolicies"`
	// Whether only branches with branch protection rules can deploy to this environment.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/repository_environment#protected_branches RepositoryEnvironment#protected_branches}
	ProtectedBranches interface{} `field:"required" json:"protectedBranches" yaml:"protectedBranches"`
}
