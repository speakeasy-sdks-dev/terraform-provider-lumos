// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/teamlumos/terraform-provider-lumos/internal/sdk/internal/utils"
	"github.com/teamlumos/terraform-provider-lumos/internal/sdk/models/shared"
	"net/http"
)

type GetGroupsRequest struct {
	IntegrationSpecificID *string `queryParam:"style=form,explode=true,name=integration_specific_id"`
	Name                  *string `queryParam:"style=form,explode=true,name=name"`
	ExactMatch            *bool   `default:"false" queryParam:"style=form,explode=true,name=exact_match"`
	AppID                 *string `queryParam:"style=form,explode=true,name=app_id"`
	Page                  *int64  `default:"1" queryParam:"style=form,explode=true,name=page"`
	Size                  *int64  `default:"50" queryParam:"style=form,explode=true,name=size"`
}

func (g GetGroupsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GetGroupsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GetGroupsRequest) GetIntegrationSpecificID() *string {
	if o == nil {
		return nil
	}
	return o.IntegrationSpecificID
}

func (o *GetGroupsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *GetGroupsRequest) GetExactMatch() *bool {
	if o == nil {
		return nil
	}
	return o.ExactMatch
}

func (o *GetGroupsRequest) GetAppID() *string {
	if o == nil {
		return nil
	}
	return o.AppID
}

func (o *GetGroupsRequest) GetPage() *int64 {
	if o == nil {
		return nil
	}
	return o.Page
}

func (o *GetGroupsRequest) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

type GetGroupsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successful Response
	PageGroup *shared.PageGroup
	// Validation Error
	HTTPValidationError *shared.HTTPValidationError
}

func (o *GetGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetGroupsResponse) GetPageGroup() *shared.PageGroup {
	if o == nil {
		return nil
	}
	return o.PageGroup
}

func (o *GetGroupsResponse) GetHTTPValidationError() *shared.HTTPValidationError {
	if o == nil {
		return nil
	}
	return o.HTTPValidationError
}
