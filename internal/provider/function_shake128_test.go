// Copyright (c) HashiCorp, Inc.

package provider

import (
	"testing"

	"regexp"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/tfversion"
)

func TestAccShake128Function_basic(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::utilities::shake128("be121283-1084-4e20-990a-e86d5995433c", 12)
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("test", "52c2059843efc17a0a7e5704"),
				),
			},
		},
	})
}

func TestAccShake128Function_differentLength(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::utilities::shake128("be121283-1084-4e20-990a-e86d5995433c", 32)
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("test", "52c2059843efc17a0a7e5704a03975ec69d919a6f19b4ba8c95057325ab766ea"),
				),
			},
		},
	})
}

func TestAccShake128Function_emptyInput(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::utilities::shake128("", 16)
					}
				`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckOutput("test", "7f9c2ba4e88f827d616045507605853e"),
				),
			},
		},
	})
}

func TestAccShake128Function_zeroLength(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::utilities::shake128("test", 0)
					}
				`,
				ExpectError: regexp.MustCompile(`length must be a positive integer`),
			},
		},
	})
}

func TestAccShake128Function_negativeLength(t *testing.T) {
	resource.UnitTest(t, resource.TestCase{
		TerraformVersionChecks: []tfversion.TerraformVersionCheck{
			tfversion.SkipBelow(tfversion.Version1_8_0),
		},
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: `
					output "test" {
						value = provider::utilities::shake128("test", -1)
					}
				`,
				ExpectError: regexp.MustCompile(`length must be a positive integer`),
			},
		},
	})
}
