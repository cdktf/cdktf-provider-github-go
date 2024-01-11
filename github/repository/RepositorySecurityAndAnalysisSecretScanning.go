// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repository


type RepositorySecurityAndAnalysisSecretScanning struct {
	// Set to 'enabled' to enable secret scanning on the repository.
	//
	// Can be 'enabled' or 'disabled'. If set to 'enabled', the repository's visibility must be 'public' or 'security_and_analysis[0].advanced_security[0].status' must also be set to 'enabled'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.44.0/docs/resources/repository#status Repository#status}
	Status *string `field:"required" json:"status" yaml:"status"`
}

