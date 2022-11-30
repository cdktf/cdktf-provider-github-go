package organizationsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationSettingsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#billing_email OrganizationSettings#billing_email}.
	BillingEmail *string `field:"required" json:"billingEmail" yaml:"billingEmail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#advanced_security_enabled_for_new_repositories OrganizationSettings#advanced_security_enabled_for_new_repositories}.
	AdvancedSecurityEnabledForNewRepositories interface{} `field:"optional" json:"advancedSecurityEnabledForNewRepositories" yaml:"advancedSecurityEnabledForNewRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#blog OrganizationSettings#blog}.
	Blog *string `field:"optional" json:"blog" yaml:"blog"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#company OrganizationSettings#company}.
	Company *string `field:"optional" json:"company" yaml:"company"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#default_repository_permission OrganizationSettings#default_repository_permission}.
	DefaultRepositoryPermission *string `field:"optional" json:"defaultRepositoryPermission" yaml:"defaultRepositoryPermission"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#dependabot_alerts_enabled_for_new_repositories OrganizationSettings#dependabot_alerts_enabled_for_new_repositories}.
	DependabotAlertsEnabledForNewRepositories interface{} `field:"optional" json:"dependabotAlertsEnabledForNewRepositories" yaml:"dependabotAlertsEnabledForNewRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#dependabot_security_updates_enabled_for_new_repositories OrganizationSettings#dependabot_security_updates_enabled_for_new_repositories}.
	DependabotSecurityUpdatesEnabledForNewRepositories interface{} `field:"optional" json:"dependabotSecurityUpdatesEnabledForNewRepositories" yaml:"dependabotSecurityUpdatesEnabledForNewRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#dependency_graph_enabled_for_new_repositories OrganizationSettings#dependency_graph_enabled_for_new_repositories}.
	DependencyGraphEnabledForNewRepositories interface{} `field:"optional" json:"dependencyGraphEnabledForNewRepositories" yaml:"dependencyGraphEnabledForNewRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#description OrganizationSettings#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#email OrganizationSettings#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#has_organization_projects OrganizationSettings#has_organization_projects}.
	HasOrganizationProjects interface{} `field:"optional" json:"hasOrganizationProjects" yaml:"hasOrganizationProjects"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#has_repository_projects OrganizationSettings#has_repository_projects}.
	HasRepositoryProjects interface{} `field:"optional" json:"hasRepositoryProjects" yaml:"hasRepositoryProjects"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#id OrganizationSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#location OrganizationSettings#location}.
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Setting to true allows organization members to create internal repositories. Only available to Enterprise Organizations.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#members_can_create_internal_repositories OrganizationSettings#members_can_create_internal_repositories}
	MembersCanCreateInternalRepositories interface{} `field:"optional" json:"membersCanCreateInternalRepositories" yaml:"membersCanCreateInternalRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#members_can_create_pages OrganizationSettings#members_can_create_pages}.
	MembersCanCreatePages interface{} `field:"optional" json:"membersCanCreatePages" yaml:"membersCanCreatePages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#members_can_create_private_pages OrganizationSettings#members_can_create_private_pages}.
	MembersCanCreatePrivatePages interface{} `field:"optional" json:"membersCanCreatePrivatePages" yaml:"membersCanCreatePrivatePages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#members_can_create_private_repositories OrganizationSettings#members_can_create_private_repositories}.
	MembersCanCreatePrivateRepositories interface{} `field:"optional" json:"membersCanCreatePrivateRepositories" yaml:"membersCanCreatePrivateRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#members_can_create_public_pages OrganizationSettings#members_can_create_public_pages}.
	MembersCanCreatePublicPages interface{} `field:"optional" json:"membersCanCreatePublicPages" yaml:"membersCanCreatePublicPages"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#members_can_create_public_repositories OrganizationSettings#members_can_create_public_repositories}.
	MembersCanCreatePublicRepositories interface{} `field:"optional" json:"membersCanCreatePublicRepositories" yaml:"membersCanCreatePublicRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#members_can_create_repositories OrganizationSettings#members_can_create_repositories}.
	MembersCanCreateRepositories interface{} `field:"optional" json:"membersCanCreateRepositories" yaml:"membersCanCreateRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#members_can_fork_private_repositories OrganizationSettings#members_can_fork_private_repositories}.
	MembersCanForkPrivateRepositories interface{} `field:"optional" json:"membersCanForkPrivateRepositories" yaml:"membersCanForkPrivateRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#name OrganizationSettings#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#secret_scanning_enabled_for_new_repositories OrganizationSettings#secret_scanning_enabled_for_new_repositories}.
	SecretScanningEnabledForNewRepositories interface{} `field:"optional" json:"secretScanningEnabledForNewRepositories" yaml:"secretScanningEnabledForNewRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#secret_scanning_push_protection_enabled_for_new_repositories OrganizationSettings#secret_scanning_push_protection_enabled_for_new_repositories}.
	SecretScanningPushProtectionEnabledForNewRepositories interface{} `field:"optional" json:"secretScanningPushProtectionEnabledForNewRepositories" yaml:"secretScanningPushProtectionEnabledForNewRepositories"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#twitter_username OrganizationSettings#twitter_username}.
	TwitterUsername *string `field:"optional" json:"twitterUsername" yaml:"twitterUsername"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/organization_settings#web_commit_signoff_required OrganizationSettings#web_commit_signoff_required}.
	WebCommitSignoffRequired interface{} `field:"optional" json:"webCommitSignoffRequired" yaml:"webCommitSignoffRequired"`
}

