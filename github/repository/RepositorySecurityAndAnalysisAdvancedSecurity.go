// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package repository


type RepositorySecurityAndAnalysisAdvancedSecurity struct {
	// Set to 'enabled' to enable advanced security features on the repository. Can be 'enabled' or 'disabled'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.8.2/docs/resources/repository#status Repository#status}
	Status *string `field:"required" json:"status" yaml:"status"`
}

