// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider


type GithubProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#alias GithubProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// app_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#app_auth GithubProvider#app_auth}
	AppAuth *GithubProviderAppAuth `field:"optional" json:"appAuth" yaml:"appAuth"`
	// The GitHub Base API URL.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#base_url GithubProvider#base_url}
	BaseUrl *string `field:"optional" json:"baseUrl" yaml:"baseUrl"`
	// Enable `insecure` mode for testing purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#insecure GithubProvider#insecure}
	Insecure interface{} `field:"optional" json:"insecure" yaml:"insecure"`
	// Number of items per page for paginationDefaults to 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#max_per_page GithubProvider#max_per_page}
	MaxPerPage *float64 `field:"optional" json:"maxPerPage" yaml:"maxPerPage"`
	// Number of times to retry a request after receiving an error status codeDefaults to 3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#max_retries GithubProvider#max_retries}
	MaxRetries *float64 `field:"optional" json:"maxRetries" yaml:"maxRetries"`
	// The GitHub organization name to manage. Use this field instead of `owner` when managing organization accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#organization GithubProvider#organization}
	Organization *string `field:"optional" json:"organization" yaml:"organization"`
	// The GitHub owner name to manage. Use this field instead of `organization` when managing individual accounts.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#owner GithubProvider#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Allow the provider to make parallel API calls to GitHub.
	//
	// You may want to set it to true when you have a private Github Enterprise without strict rate limits. While it is possible to enable this setting on github.com, github.com's best practices recommend using serialization to avoid hitting abuse rate limitsDefaults to false if not set
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#parallel_requests GithubProvider#parallel_requests}
	ParallelRequests interface{} `field:"optional" json:"parallelRequests" yaml:"parallelRequests"`
	// Amount of time in milliseconds to sleep in between non-write requests to GitHub API.
	//
	// Defaults to 0ms if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#read_delay_ms GithubProvider#read_delay_ms}
	ReadDelayMs *float64 `field:"optional" json:"readDelayMs" yaml:"readDelayMs"`
	// Allow the provider to retry after receiving an error status code, the max_retries should be set for this to workDefaults to [500, 502, 503, 504].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#retryable_errors GithubProvider#retryable_errors}
	RetryableErrors *[]*float64 `field:"optional" json:"retryableErrors" yaml:"retryableErrors"`
	// Amount of time in milliseconds to sleep in between requests to GitHub API after an error response.
	//
	// Defaults to 1000ms or 1s if not set, the max_retries must be set to greater than zero.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#retry_delay_ms GithubProvider#retry_delay_ms}
	RetryDelayMs *float64 `field:"optional" json:"retryDelayMs" yaml:"retryDelayMs"`
	// The OAuth token used to connect to GitHub.
	//
	// Anonymous mode is enabled if both `token` and `app_auth` are not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#token GithubProvider#token}
	Token *string `field:"optional" json:"token" yaml:"token"`
	// Amount of time in milliseconds to sleep in between writes to GitHub API.
	//
	// Defaults to 1000ms or 1s if not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.9.0/docs#write_delay_ms GithubProvider#write_delay_ms}
	WriteDelayMs *float64 `field:"optional" json:"writeDelayMs" yaml:"writeDelayMs"`
}

