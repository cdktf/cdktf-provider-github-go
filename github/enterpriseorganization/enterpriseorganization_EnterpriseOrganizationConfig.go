package enterpriseorganization

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EnterpriseOrganizationConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/enterprise_organization#admin_logins EnterpriseOrganization#admin_logins}.
	AdminLogins *[]*string `field:"required" json:"adminLogins" yaml:"adminLogins"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/enterprise_organization#billing_email EnterpriseOrganization#billing_email}.
	BillingEmail *string `field:"required" json:"billingEmail" yaml:"billingEmail"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/enterprise_organization#enterprise_id EnterpriseOrganization#enterprise_id}.
	EnterpriseId *string `field:"required" json:"enterpriseId" yaml:"enterpriseId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/enterprise_organization#name EnterpriseOrganization#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/enterprise_organization#description EnterpriseOrganization#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/github/r/enterprise_organization#id EnterpriseOrganization#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

