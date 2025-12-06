// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package enterprisesecurityanalysissettings

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseSecurityAnalysisSettingsConfig struct {
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
	// The slug of the enterprise.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings#enterprise_slug EnterpriseSecurityAnalysisSettings#enterprise_slug}
	EnterpriseSlug *string `field:"required" json:"enterpriseSlug" yaml:"enterpriseSlug"`
	// Whether GitHub Advanced Security is automatically enabled for new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings#advanced_security_enabled_for_new_repositories EnterpriseSecurityAnalysisSettings#advanced_security_enabled_for_new_repositories}
	AdvancedSecurityEnabledForNewRepositories interface{} `field:"optional" json:"advancedSecurityEnabledForNewRepositories" yaml:"advancedSecurityEnabledForNewRepositories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings#id EnterpriseSecurityAnalysisSettings#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Whether secret scanning is automatically enabled for new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings#secret_scanning_enabled_for_new_repositories EnterpriseSecurityAnalysisSettings#secret_scanning_enabled_for_new_repositories}
	SecretScanningEnabledForNewRepositories interface{} `field:"optional" json:"secretScanningEnabledForNewRepositories" yaml:"secretScanningEnabledForNewRepositories"`
	// Custom URL for secret scanning push protection bypass instructions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings#secret_scanning_push_protection_custom_link EnterpriseSecurityAnalysisSettings#secret_scanning_push_protection_custom_link}
	SecretScanningPushProtectionCustomLink *string `field:"optional" json:"secretScanningPushProtectionCustomLink" yaml:"secretScanningPushProtectionCustomLink"`
	// Whether secret scanning push protection is automatically enabled for new repositories.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings#secret_scanning_push_protection_enabled_for_new_repositories EnterpriseSecurityAnalysisSettings#secret_scanning_push_protection_enabled_for_new_repositories}
	SecretScanningPushProtectionEnabledForNewRepositories interface{} `field:"optional" json:"secretScanningPushProtectionEnabledForNewRepositories" yaml:"secretScanningPushProtectionEnabledForNewRepositories"`
	// Whether secret scanning validity checks are enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/enterprise_security_analysis_settings#secret_scanning_validity_checks_enabled EnterpriseSecurityAnalysisSettings#secret_scanning_validity_checks_enabled}
	SecretScanningValidityChecksEnabled interface{} `field:"optional" json:"secretScanningValidityChecksEnabled" yaml:"secretScanningValidityChecksEnabled"`
}

