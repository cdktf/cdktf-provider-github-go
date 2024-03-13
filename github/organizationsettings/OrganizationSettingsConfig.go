// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package organizationsettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrganizationSettingsConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
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
	// The billing email address for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#billing_email OrganizationSettings#billing_email}
	BillingEmail *string `field:"required" json:"billingEmail" yaml:"billingEmail"`
	// Whether or not advanced security is enabled for new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#advanced_security_enabled_for_new_repositories OrganizationSettings#advanced_security_enabled_for_new_repositories}
	AdvancedSecurityEnabledForNewRepositories interface{} `field:"optional" json:"advancedSecurityEnabledForNewRepositories" yaml:"advancedSecurityEnabledForNewRepositories"`
	// The blog URL for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#blog OrganizationSettings#blog}
	Blog *string `field:"optional" json:"blog" yaml:"blog"`
	// The company name for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#company OrganizationSettings#company}
	Company *string `field:"optional" json:"company" yaml:"company"`
	// The default permission for organization members to create new repositories. Can be one of 'read', 'write', 'admin' or 'none'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#default_repository_permission OrganizationSettings#default_repository_permission}
	DefaultRepositoryPermission *string `field:"optional" json:"defaultRepositoryPermission" yaml:"defaultRepositoryPermission"`
	// Whether or not dependabot alerts are enabled for new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#dependabot_alerts_enabled_for_new_repositories OrganizationSettings#dependabot_alerts_enabled_for_new_repositories}
	DependabotAlertsEnabledForNewRepositories interface{} `field:"optional" json:"dependabotAlertsEnabledForNewRepositories" yaml:"dependabotAlertsEnabledForNewRepositories"`
	// Whether or not dependabot security updates are enabled for new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#dependabot_security_updates_enabled_for_new_repositories OrganizationSettings#dependabot_security_updates_enabled_for_new_repositories}
	DependabotSecurityUpdatesEnabledForNewRepositories interface{} `field:"optional" json:"dependabotSecurityUpdatesEnabledForNewRepositories" yaml:"dependabotSecurityUpdatesEnabledForNewRepositories"`
	// Whether or not dependency graph is enabled for new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#dependency_graph_enabled_for_new_repositories OrganizationSettings#dependency_graph_enabled_for_new_repositories}
	DependencyGraphEnabledForNewRepositories interface{} `field:"optional" json:"dependencyGraphEnabledForNewRepositories" yaml:"dependencyGraphEnabledForNewRepositories"`
	// The description for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#description OrganizationSettings#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The email address for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#email OrganizationSettings#email}
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Whether or not organization projects are enabled for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#has_organization_projects OrganizationSettings#has_organization_projects}
	HasOrganizationProjects interface{} `field:"optional" json:"hasOrganizationProjects" yaml:"hasOrganizationProjects"`
	// Whether or not repository projects are enabled for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#has_repository_projects OrganizationSettings#has_repository_projects}
	HasRepositoryProjects interface{} `field:"optional" json:"hasRepositoryProjects" yaml:"hasRepositoryProjects"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#id OrganizationSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The location for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#location OrganizationSettings#location}
	Location *string `field:"optional" json:"location" yaml:"location"`
	// Whether or not organization members can create new internal repositories. For Enterprise Organizations only.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#members_can_create_internal_repositories OrganizationSettings#members_can_create_internal_repositories}
	MembersCanCreateInternalRepositories interface{} `field:"optional" json:"membersCanCreateInternalRepositories" yaml:"membersCanCreateInternalRepositories"`
	// Whether or not organization members can create new pages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#members_can_create_pages OrganizationSettings#members_can_create_pages}
	MembersCanCreatePages interface{} `field:"optional" json:"membersCanCreatePages" yaml:"membersCanCreatePages"`
	// Whether or not organization members can create new private pages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#members_can_create_private_pages OrganizationSettings#members_can_create_private_pages}
	MembersCanCreatePrivatePages interface{} `field:"optional" json:"membersCanCreatePrivatePages" yaml:"membersCanCreatePrivatePages"`
	// Whether or not organization members can create new private repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#members_can_create_private_repositories OrganizationSettings#members_can_create_private_repositories}
	MembersCanCreatePrivateRepositories interface{} `field:"optional" json:"membersCanCreatePrivateRepositories" yaml:"membersCanCreatePrivateRepositories"`
	// Whether or not organization members can create new public pages.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#members_can_create_public_pages OrganizationSettings#members_can_create_public_pages}
	MembersCanCreatePublicPages interface{} `field:"optional" json:"membersCanCreatePublicPages" yaml:"membersCanCreatePublicPages"`
	// Whether or not organization members can create new public repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#members_can_create_public_repositories OrganizationSettings#members_can_create_public_repositories}
	MembersCanCreatePublicRepositories interface{} `field:"optional" json:"membersCanCreatePublicRepositories" yaml:"membersCanCreatePublicRepositories"`
	// Whether or not organization members can create new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#members_can_create_repositories OrganizationSettings#members_can_create_repositories}
	MembersCanCreateRepositories interface{} `field:"optional" json:"membersCanCreateRepositories" yaml:"membersCanCreateRepositories"`
	// Whether or not organization members can fork private repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#members_can_fork_private_repositories OrganizationSettings#members_can_fork_private_repositories}
	MembersCanForkPrivateRepositories interface{} `field:"optional" json:"membersCanForkPrivateRepositories" yaml:"membersCanForkPrivateRepositories"`
	// The name for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#name OrganizationSettings#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether or not secret scanning is enabled for new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#secret_scanning_enabled_for_new_repositories OrganizationSettings#secret_scanning_enabled_for_new_repositories}
	SecretScanningEnabledForNewRepositories interface{} `field:"optional" json:"secretScanningEnabledForNewRepositories" yaml:"secretScanningEnabledForNewRepositories"`
	// Whether or not secret scanning push protection is enabled for new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#secret_scanning_push_protection_enabled_for_new_repositories OrganizationSettings#secret_scanning_push_protection_enabled_for_new_repositories}
	SecretScanningPushProtectionEnabledForNewRepositories interface{} `field:"optional" json:"secretScanningPushProtectionEnabledForNewRepositories" yaml:"secretScanningPushProtectionEnabledForNewRepositories"`
	// The Twitter username for the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#twitter_username OrganizationSettings#twitter_username}
	TwitterUsername *string `field:"optional" json:"twitterUsername" yaml:"twitterUsername"`
	// Whether or not commit signatures are required for commits to the organization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.1.0/docs/resources/organization_settings#web_commit_signoff_required OrganizationSettings#web_commit_signoff_required}
	WebCommitSignoffRequired interface{} `field:"optional" json:"webCommitSignoffRequired" yaml:"webCommitSignoffRequired"`
}

