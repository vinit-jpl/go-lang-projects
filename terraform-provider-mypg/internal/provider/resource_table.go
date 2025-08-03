package provider

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	_ "github.com/lib/pq"
)

type tableResource struct{}

type tableModel struct {
	Name   types.String `tfsdk:"name"`
	Schema types.String `tfsdk:"schema"`
}

func NewTableResource() resource.Resource {
	return &tableResource{}
}

func (r *tableResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = "mypg_table"
}

func (r *tableResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name":   schema.StringAttribute{Required: true},
			"schema": schema.StringAttribute{Required: true},
		},
	}
}

func (r *tableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan tableModel
	_ = req.Plan.Get(ctx, &plan)

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		resp.Diagnostics.AddError("PostgreSQL Connection Failed", err.Error())
		return
	}
	defer db.Close()

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", plan.Name.ValueString(), plan.Schema.ValueString())
	_, err = db.Exec(query)
	if err != nil {
		resp.Diagnostics.AddError("Table Creation Failed", err.Error())
		return
	}

	resp.State.Set(ctx, plan)
}

func (r *tableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// No-op for now
	resp.State.Set(ctx, req.State)
}

func (r *tableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Not implemented
}

func (r *tableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state tableModel
	_ = req.State.Get(ctx, &state)

	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable")
	if err != nil {
		resp.Diagnostics.AddError("PostgreSQL Connection Failed", err.Error())
		return
	}
	defer db.Close()

	query := fmt.Sprintf("DROP TABLE IF EXISTS %s;", state.Name.ValueString())
	_, err = db.Exec(query)
	if err != nil {
		resp.Diagnostics.AddError("Table Deletion Failed", err.Error())
		return
	}
}
