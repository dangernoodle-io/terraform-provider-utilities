# terraform-provider-utilities
[![terraform-provider-test](https://github.com/dangernoodle-io/terraform-provider-utilities/actions/workflows/test.yml/badge.svg?branch=main)](https://github.com/dangernoodle-io/terraform-provider-utilities/actions/workflows/test.yml)
[![terraform-provider-release](https://github.com/dangernoodle-io/terraform-provider-utilities/actions/workflows/release.yml/badge.svg)](https://github.com/dangernoodle-io/terraform-provider-utilities/actions/workflows/release.yml)
[![GitHub Release](https://img.shields.io/github/v/release/dangernoodle-io/terraform-provider-utilities)](https://github.com/dangernoodle-io/terraform-provider-utilities/releases/latest)
[![Coverage Status](https://coveralls.io/repos/github/dangernoodle-io/terraform-provider-utilities/badge.svg?branch=main)](https://coveralls.io/github/dangernoodle-io/terraform-provider-utilities?branch=main)
[![Terraform Registry](https://img.shields.io/badge/Terraform%20Registry-dangernoodle--io%2Futilities-7B42BC?logo=terraform)](https://registry.terraform.io/providers/dangernoodle-io/utilities/latest/docs)
[![Go version](https://img.shields.io/github/go-mod/go-version/dangernoodle-io/terraform-provider-utilities)](go.mod)

A Terraform provider that exposes utility functions for use in Terraform configurations.

## Functions

### `shake128`

Generates a [SHAKE-128](https://en.wikipedia.org/wiki/SHA-3#Instances) hash of an input string with a configurable output length.

```terraform
terraform {
  required_providers {
    utilities = {
      source = "dangernoodle-io/utilities"
    }
  }
}

provider "utilities" {}

output "hash" {
  value = provider::utilities::shake128("my-input-string", 16)
}
```

#### Arguments

- `input` (String) - Input string to hash.
- `length` (Number) - Length of the hash output in bytes.

#### Return

A hex-encoded string of the SHAKE-128 hash.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.23 (for building)

## Developing

```shell
make          # fmt, lint, install, generate
make test     # unit tests
make testacc  # acceptance tests
```
