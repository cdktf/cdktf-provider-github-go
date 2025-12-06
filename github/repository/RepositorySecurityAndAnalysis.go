// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repository


type RepositorySecurityAndAnalysis struct {
	// advanced_security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/repository#advanced_security Repository#advanced_security}
	AdvancedSecurity *RepositorySecurityAndAnalysisAdvancedSecurity `field:"optional" json:"advancedSecurity" yaml:"advancedSecurity"`
	// code_security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/repository#code_security Repository#code_security}
	CodeSecurity *RepositorySecurityAndAnalysisCodeSecurity `field:"optional" json:"codeSecurity" yaml:"codeSecurity"`
	// secret_scanning block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/repository#secret_scanning Repository#secret_scanning}
	SecretScanning *RepositorySecurityAndAnalysisSecretScanning `field:"optional" json:"secretScanning" yaml:"secretScanning"`
	// secret_scanning_ai_detection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/repository#secret_scanning_ai_detection Repository#secret_scanning_ai_detection}
	SecretScanningAiDetection *RepositorySecurityAndAnalysisSecretScanningAiDetection `field:"optional" json:"secretScanningAiDetection" yaml:"secretScanningAiDetection"`
	// secret_scanning_non_provider_patterns block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/repository#secret_scanning_non_provider_patterns Repository#secret_scanning_non_provider_patterns}
	SecretScanningNonProviderPatterns *RepositorySecurityAndAnalysisSecretScanningNonProviderPatterns `field:"optional" json:"secretScanningNonProviderPatterns" yaml:"secretScanningNonProviderPatterns"`
	// secret_scanning_push_protection block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/repository#secret_scanning_push_protection Repository#secret_scanning_push_protection}
	SecretScanningPushProtection *RepositorySecurityAndAnalysisSecretScanningPushProtection `field:"optional" json:"secretScanningPushProtection" yaml:"secretScanningPushProtection"`
}

