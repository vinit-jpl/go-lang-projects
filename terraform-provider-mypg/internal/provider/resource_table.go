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

type tableResourceModel struct {
	ID       types.String `tfsdk:"id"`
	Name     types.String `tfsdk:"name"`
	Columns  types.List   `tfsdk:"columns"`
	Database types.String `tfsdk:"database"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
	Host     types.String `tfsdk:"host"`
	Port     types.Int64  `tfsdk:"port"`
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
			"id":       schema.StringAttribute{Computed: true},
			"name":     schema.StringAttribute{Required: true},
			"columns":  schema.ListAttribute{Required: true, ElementType: types.StringType},
			"database": schema.StringAttribute{Required: true},
			"username": schema.StringAttribute{Required: true},
			"password": schema.StringAttribute{Required: true, Sensitive: true},
			"host":     schema.StringAttribute{Required: true},
			"port":     schema.Int64Attribute{Optional: true},
		},
	}
}

func (r *tableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data tableResourceModel
	diags := req.Plan.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		data.Host.ValueString(),
		data.Port.ValueInt64(),
		data.Username.ValueString(),
		data.Password.ValueString(),
		data.Database.ValueString(),
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		resp.Diagnostics.AddError("Database Connection Failed", err.Error())
		return
	}
	defer db.Close()

	columnDefs := ""
	var columnList []string
	diags = data.Columns.ElementsAs(ctx, &columnList, false)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	for i, col := range columnList {
		columnDefs += col
		if i < len(columnList)-1 {
			columnDefs += ", "
		}
	}

	createStmt := fmt.Sprintf("CREATE TABLE %s (%s);", data.Name.ValueString(), columnDefs)
	_, err = db.Exec(createStmt)
	if err != nil {
		resp.Diagnostics.AddError("Table Creation Failed", err.Error())
		return
	}

	data.ID = types.StringValue(data.Name.ValueString())
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *tableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Optional: Implement if needed
}

func (r *tableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Optional: Implement if needed
}

func (r *tableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Optional: Implement table deletion logic if needed
}
