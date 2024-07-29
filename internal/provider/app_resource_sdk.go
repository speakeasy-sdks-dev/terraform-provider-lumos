// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/teamlumos/terraform-provider-lumos/internal/sdk/models/shared"
)

func (r *AppResourceModel) ToSharedAppInputCreate() *shared.AppInputCreate {
	var name string
	name = r.Name.ValueString()

	var category string
	category = r.Category.ValueString()

	var description string
	description = r.Description.ValueString()

	logoURL := new(string)
	if !r.LogoURL.IsUnknown() && !r.LogoURL.IsNull() {
		*logoURL = r.LogoURL.ValueString()
	} else {
		logoURL = nil
	}
	websiteURL := new(string)
	if !r.WebsiteURL.IsUnknown() && !r.WebsiteURL.IsNull() {
		*websiteURL = r.WebsiteURL.ValueString()
	} else {
		websiteURL = nil
	}
	requestInstructions := new(string)
	if !r.RequestInstructions.IsUnknown() && !r.RequestInstructions.IsNull() {
		*requestInstructions = r.RequestInstructions.ValueString()
	} else {
		requestInstructions = nil
	}
	out := shared.AppInputCreate{
		Name:                name,
		Category:            category,
		Description:         description,
		LogoURL:             logoURL,
		WebsiteURL:          websiteURL,
		RequestInstructions: requestInstructions,
	}
	return &out
}

func (r *AppResourceModel) RefreshFromSharedApp(resp *shared.App) {
	if resp != nil {
		r.AllowMultiplePermissionSelection = types.BoolValue(resp.AllowMultiplePermissionSelection)
		r.AppClassID = types.StringValue(resp.AppClassID)
		r.ID = types.StringValue(resp.ID)
		r.InstanceID = types.StringValue(resp.InstanceID)
		r.LogoURL = types.StringPointerValue(resp.LogoURL)
		r.RequestInstructions = types.StringPointerValue(resp.RequestInstructions)
		r.Sources = []types.String{}
		for _, v := range resp.Sources {
			r.Sources = append(r.Sources, types.StringValue(string(v)))
		}
		r.Status = types.StringValue(string(resp.Status))
		r.UserFriendlyLabel = types.StringValue(resp.UserFriendlyLabel)
		r.WebsiteURL = types.StringPointerValue(resp.WebsiteURL)
	}
}
