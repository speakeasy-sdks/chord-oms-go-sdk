package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListOrdersQueryParams struct {
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Q       *string `queryParam:"style=form,explode=true,name=q"`
}

type ListOrdersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListOrdersPaginationData struct {
	Count       *int64              `json:"count,omitempty"`
	CurrentPage *int64              `json:"current_page,omitempty"`
	Orders      []shared.OrderSmall `json:"orders,omitempty"`
	Pages       *int64              `json:"pages,omitempty"`
	PerPage     *int64              `json:"per_page,omitempty"`
	TotalCount  *int64              `json:"total_count,omitempty"`
}

type ListOrders401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListOrdersRequest struct {
	QueryParams ListOrdersQueryParams
	Security    ListOrdersSecurity
}

type ListOrdersResponse struct {
	ContentType                        string
	PaginationData                     *ListOrdersPaginationData
	StatusCode                         int64
	ListOrders401ApplicationJSONObject *ListOrders401ApplicationJSON
}
