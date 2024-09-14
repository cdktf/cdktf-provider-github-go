// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package issuelabels


type IssueLabelsLabel struct {
	// A 6 character hex code, without the leading '#', identifying the color of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/issue_labels#color IssueLabels#color}
	Color *string `field:"required" json:"color" yaml:"color"`
	// The name of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/issue_labels#name IssueLabels#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// A short description of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/integrations/github/6.3.0/docs/resources/issue_labels#description IssueLabels#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

