// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repository


type RepositorySecurityAndAnalysisSecretScanning struct {
	// Set to 'enabled' to enable secret scanning on the repository.
	//
	// Can be 'enabled' or 'disabled'. If set to 'enabled', the repository's visibility must be 'public', 'security_and_analysis[0].advanced_security[0].status' must also be set to 'enabled', or your Organization must have split licensing for Advanced security.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs/resources/repository#status Repository#status}
	Status *string `field:"required" json:"status" yaml:"status"`
}

