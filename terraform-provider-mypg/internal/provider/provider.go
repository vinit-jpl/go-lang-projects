package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type postgresProvider struct{}

func (p *postgresProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "mypg"
}

func (p *postgresProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Required:    true,
				Description: "PostgreSQL host.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Description: "PostgreSQL port.",
			},
			"username": schema.StringAttribute{
				Required:    true,
				Description: "PostgreSQL username.",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "PostgreSQL password.",
			},
			"database": schema.StringAttribute{
				Required:    true,
				Description: "Database to connect to.",
			},
		},
	}

}

func (p *postgresProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewTableResource,
	}
}

func (p *postgresProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

func (p *postgresProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {

}

func New() provider.Provider {
	return &postgresProvider{}
}
