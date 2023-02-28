package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListOrdersQueryParams struct {
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Q       *string `queryParam:"style=form,explode=true,name=q"`
}

type ListOrdersSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type ListOrdersRequest struct {
	QueryParams ListOrdersQueryParams
	Security    ListOrdersSecurity
}

type ListOrders401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListOrdersPaginationData struct {
	Count       *int64              `json:"count,omitempty"`
	CurrentPage *int64              `json:"current_page,omitempty"`
	Orders      []shared.OrderSmall `json:"orders,omitempty"`
	Pages       *int64              `json:"pages,omitempty"`
	PerPage     *int64              `json:"per_page,omitempty"`
	TotalCount  *int64              `json:"total_count,omitempty"`
}

type ListOrdersResponse struct {
	ContentType                        string
	PaginationData                     *ListOrdersPaginationData
	StatusCode                         int
	ListOrders401ApplicationJSONObject *ListOrders401ApplicationJSON
}
