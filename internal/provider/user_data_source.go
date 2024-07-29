// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/teamlumos/terraform-provider-lumos/internal/sdk"
	"github.com/teamlumos/terraform-provider-lumos/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &UserDataSource{}
var _ datasource.DataSourceWithConfigure = &UserDataSource{}

func NewUserDataSource() datasource.DataSource {
	return &UserDataSource{}
}

// UserDataSource is the data source implementation.
type UserDataSource struct {
	client *sdk.Lumos
}

// UserDataSourceModel describes the data model.
type UserDataSourceModel struct {
	Email      types.String `tfsdk:"email"`
	FamilyName types.String `tfsdk:"family_name"`
	GivenName  types.String `tfsdk:"given_name"`
	ID         types.String `tfsdk:"id"`
	Status     types.String `tfsdk:"status"`
	UserID     types.String `tfsdk:"user_id"`
}

// Metadata returns the data source type name.
func (r *UserDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

// Schema defines the schema for the data source.
func (r *UserDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "User DataSource",

		Attributes: map[string]schema.Attribute{
			"email": schema.StringAttribute{
				Computed:    true,
				Description: `The email of this user.`,
			},
			"family_name": schema.StringAttribute{
				Computed:    true,
				Description: `The family name of this user.`,
			},
			"given_name": schema.StringAttribute{
				Computed:    true,
				Description: `The given name of this user.`,
			},
			"id": schema.StringAttribute{
				Computed:    true,
				Description: `The ID of this user.`,
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: `An enumeration. must be one of ["STAGED", "ACTIVE", "SUSPENDED", "INACTIVE"]`,
			},
			"user_id": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r *UserDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Lumos)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Lumos, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *UserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *UserDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var userID string
	userID = data.UserID.ValueString()

	request := operations.GetUserRequest{
		UserID: userID,
	}
	res, err := r.client.Core.GetUser(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.User != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedUser(res.User)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
