package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type helloProvider struct{}

func New() provider.Provider {
	return &helloProvider{}
}

func (p *helloProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "hello"
}

func (p *helloProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

func (p *helloProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *helloProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewHelloResource,
	}
}

func (p *helloProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}
