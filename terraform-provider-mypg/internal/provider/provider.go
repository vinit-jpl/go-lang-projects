package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Define the provider struct
type mypgProvider struct{}

// Metadata sets the provider type name
func (p *mypgProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "mypg"
	resp.Version = "1.0.0"
}

// GetSchema returns the provider-level schema (empty for now)
func (p *mypgProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}

// Configure configures the provider (no-op for now)
func (p *mypgProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// No configuration for now
}

// Resources returns the list of supported resources
func (p *mypgProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTableResource,
	}
}

// DataSources returns the list of supported data sources
func (p *mypgProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// New returns a new instance of the provider
func New() provider.Provider {
	return &mypgProvider{}
}
