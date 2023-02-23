package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListReturnAuthorizationQueryParams struct {
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Q       *string `queryParam:"style=form,explode=true,name=q"`
}

type ListReturnAuthorizationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListReturnAuthorizationRequest struct {
	QueryParams ListReturnAuthorizationQueryParams
	Security    ListReturnAuthorizationSecurity
}

type ListReturnAuthorization404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListReturnAuthorization401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListReturnAuthorizationPaginationData struct {
	Count                *int64                   `json:"count,omitempty"`
	CurrentPage          *int64                   `json:"current_page,omitempty"`
	Pages                *int64                   `json:"pages,omitempty"`
	PerPage              *int64                   `json:"per_page,omitempty"`
	ReturnAuthorizations []map[string]interface{} `json:"return_authorizations,omitempty"`
	TotalCount           *int64                   `json:"total_count,omitempty"`
}

type ListReturnAuthorizationResponse struct {
	ContentType                                     string
	PaginationData                                  *ListReturnAuthorizationPaginationData
	StatusCode                                      int
	ListReturnAuthorization401ApplicationJSONObject *ListReturnAuthorization401ApplicationJSON
	ListReturnAuthorization404ApplicationJSONObject *ListReturnAuthorization404ApplicationJSON
}
