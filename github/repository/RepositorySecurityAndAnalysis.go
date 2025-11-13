// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repository


type RepositorySecurityAndAnalysis struct {
	// advanced_security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/repository#advanced_security Repository#advanced_security}
	AdvancedSecurity *RepositorySecurityAndAnalysisAdvancedSecurity `field:"optional" json:"advancedSecurity" yaml:"advancedSecurity"`
	// secret_scanning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/repository#secret_scanning Repository#secret_scanning}
	SecretScanning *RepositorySecurityAndAnalysisSecretScanning `field:"optional" json:"secretScanning" yaml:"secretScanning"`
	// secret_scanning_push_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.0/docs/resources/repository#secret_scanning_push_protection Repository#secret_scanning_push_protection}
	SecretScanningPushProtection *RepositorySecurityAndAnalysisSecretScanningPushProtection `field:"optional" json:"secretScanningPushProtection" yaml:"secretScanningPushProtection"`
}

