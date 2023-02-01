package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListReturnAuthorizationReturnItemsPathParams struct {
	ReturnAuthorizationNumber string `pathParam:"style=simple,explode=false,name=return_authorization_number"`
}

type ListReturnAuthorizationReturnItemsQueryParams struct {
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Q       *string `queryParam:"style=form,explode=true,name=q"`
}

type ListReturnAuthorizationReturnItemsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListReturnAuthorizationReturnItemsPaginationData struct {
	Count       *int64              `json:"count,omitempty"`
	CurrentPage *int64              `json:"current_page,omitempty"`
	Pages       *int64              `json:"pages,omitempty"`
	PerPage     *int64              `json:"per_page,omitempty"`
	ReturnItems []shared.ReturnItem `json:"return_items,omitempty"`
	TotalCount  *int64              `json:"total_count,omitempty"`
}

type ListReturnAuthorizationReturnItems401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListReturnAuthorizationReturnItems404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListReturnAuthorizationReturnItemsRequest struct {
	PathParams  ListReturnAuthorizationReturnItemsPathParams
	QueryParams ListReturnAuthorizationReturnItemsQueryParams
	Security    ListReturnAuthorizationReturnItemsSecurity
}

type ListReturnAuthorizationReturnItemsResponse struct {
	ContentType                                                string
	PaginationData                                             *ListReturnAuthorizationReturnItemsPaginationData
	StatusCode                                                 int64
	ListReturnAuthorizationReturnItems401ApplicationJSONObject *ListReturnAuthorizationReturnItems401ApplicationJSON
	ListReturnAuthorizationReturnItems404ApplicationJSONObject *ListReturnAuthorizationReturnItems404ApplicationJSON
}
