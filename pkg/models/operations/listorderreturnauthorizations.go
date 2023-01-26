package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListOrderReturnAuthorizationsPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type ListOrderReturnAuthorizationsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListOrderReturnAuthorizationsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListOrderReturnAuthorizationsPaginationData struct {
	Count                *int64                   `json:"count,omitempty"`
	CurrentPage          *int64                   `json:"current_page,omitempty"`
	Pages                *int64                   `json:"pages,omitempty"`
	PerPage              *int64                   `json:"per_page,omitempty"`
	ReturnAuthorizations []map[string]interface{} `json:"return_authorizations,omitempty"`
	TotalCount           *int64                   `json:"total_count,omitempty"`
}

type ListOrderReturnAuthorizations401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListOrderReturnAuthorizations404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListOrderReturnAuthorizationsRequest struct {
	PathParams  ListOrderReturnAuthorizationsPathParams
	QueryParams ListOrderReturnAuthorizationsQueryParams
	Security    ListOrderReturnAuthorizationsSecurity
}

type ListOrderReturnAuthorizationsResponse struct {
	ContentType                                           string
	PaginationData                                        *ListOrderReturnAuthorizationsPaginationData
	StatusCode                                            int64
	ListOrderReturnAuthorizations401ApplicationJSONObject *ListOrderReturnAuthorizations401ApplicationJSON
	ListOrderReturnAuthorizations404ApplicationJSONObject *ListOrderReturnAuthorizations404ApplicationJSON
}
