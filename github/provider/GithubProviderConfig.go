// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type GithubProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#alias GithubProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// app_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#app_auth GithubProvider#app_auth}
	AppAuth *GithubProviderAppAuth `field:"optional" json:"appAuth" yaml:"appAuth"`
	// The GitHub Base API URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#base_url GithubProvider#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Enable `insecure` mode for testing purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#insecure GithubProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#organization GithubProvider#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#owner GithubProvider#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Allow the provider to make parallel API calls to GitHub.
	//
	// You may want to set it to true when you have a private Github Enterprise without strict rate limits. Although, it is not possible to enable this setting on github.com because we enforce the respect of github.com's best practices to avoid hitting abuse rate limitsDefaults to false if not set
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#parallel_requests GithubProvider#parallel_requests}
	ParallelRequests interface{} `field:"optional" json:"parallelRequests" yaml:"parallelRequests"`
	// Amount of time in milliseconds to sleep in between non-write requests to GitHub API.
	//
	// Defaults to 0ms if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#read_delay_ms GithubProvider#read_delay_ms}
	ReadDelayMs *float64 `field:"optional" json:"readDelayMs" yaml:"readDelayMs"`
	// The OAuth token used to connect to GitHub.
	//
	// Anonymous mode is enabled if both `token` and `app_auth` are not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#token GithubProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Amount of time in milliseconds to sleep in between writes to GitHub API.
	//
	// Defaults to 1000ms or 1s if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/5.40.0/docs#write_delay_ms GithubProvider#write_delay_ms}
	WriteDelayMs *float64 `field:"optional" json:"writeDelayMs" yaml:"writeDelayMs"`
}

