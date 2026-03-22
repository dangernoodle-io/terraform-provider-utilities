# Copyright (c) HashiCorp, Inc.

terraform {
  required_providers {
    utilities = {
      source = "registry.terraform.io/dangernoodle-io/utilities"
    }
  }
}

provider "utilities" {}

output "shake128" {
  value = provider::utilities::shake128("be121283-1084-4e20-990a-e86d5995433c", 12)
}
