// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type GithubProviderAppAuth struct {
	// The GitHub App ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.1/docs#id GithubProvider#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The GitHub App installation instance ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.1/docs#installation_id GithubProvider#installation_id}
	InstallationId *string `field:"required" json:"installationId" yaml:"installationId"`
	// The GitHub App PEM file contents.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.7.1/docs#pem_file GithubProvider#pem_file}
	PemFile *string `field:"required" json:"pemFile" yaml:"pemFile"`
}

