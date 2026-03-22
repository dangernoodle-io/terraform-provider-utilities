// Copyright (c) HashiCorp, Inc.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

var (
	_ provider.Provider              = &UtilitiesProvider{}
	_ provider.ProviderWithFunctions = &UtilitiesProvider{}
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &UtilitiesProvider{
			version: version,
		}
	}
}

type UtilitiesProvider struct {
	version string
}

func (p *UtilitiesProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "utilities"
	resp.Version = p.version
}

func (p *UtilitiesProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *UtilitiesProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *UtilitiesProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *UtilitiesProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

func (p *UtilitiesProvider) Functions(_ context.Context) []func() function.Function {
	return []func() function.Function{
		NewShake128Function,
	}
}
