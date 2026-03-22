// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"
	"encoding/hex"

	"golang.org/x/crypto/sha3"

	"github.com/hashicorp/terraform-plugin-framework/function"
)

var _ function.Function = &Shake128Function{}

type Shake128Function struct{}

func NewShake128Function() function.Function {
	return &Shake128Function{}
}

func (f *Shake128Function) Metadata(ctx context.Context, req function.MetadataRequest, resp *function.MetadataResponse) {
	resp.Name = "shake128"
}

func (f *Shake128Function) Definition(ctx context.Context, req function.DefinitionRequest, resp *function.DefinitionResponse) {
	resp.Definition = function.Definition{
		Summary:     "SHAKE128 hash",
		Description: "Generate a SHAKE128 hash of an input string.",

		Parameters: []function.Parameter{
			function.StringParameter{
				Name:        "input",
				Description: "Input string to hash.",
			},
			function.Int32Parameter{
				Name:        "length",
				Description: "Length of the hash in bytes.",
			},
		},
		Return: function.StringReturn{},
	}
}

func (f *Shake128Function) Run(ctx context.Context, req function.RunRequest, resp *function.RunResponse) {
	var input string
	var length int

	resp.Error = req.Arguments.Get(ctx, &input, &length)
	if resp.Error != nil {
		return
	}

	if length <= 0 {
		resp.Error = function.NewArgumentFuncError(1, "length must be a positive integer")
		return
	}

	output := make([]byte, length)
	sha3.ShakeSum128(output, []byte(input))

	resp.Error = resp.Result.Set(ctx, hex.EncodeToString(output))
}
