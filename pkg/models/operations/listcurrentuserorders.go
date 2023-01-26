package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListCurrentUserOrdersQueryParams struct {
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Q       *string `queryParam:"style=form,explode=true,name=q"`
}

type ListCurrentUserOrdersSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListCurrentUserOrdersPaginationData struct {
	Count       *int64              `json:"count,omitempty"`
	CurrentPage *int64              `json:"current_page,omitempty"`
	Orders      []shared.OrderSmall `json:"orders,omitempty"`
	Pages       *int64              `json:"pages,omitempty"`
	PerPage     *int64              `json:"per_page,omitempty"`
	TotalCount  *int64              `json:"total_count,omitempty"`
}

type ListCurrentUserOrders401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListCurrentUserOrdersRequest struct {
	QueryParams ListCurrentUserOrdersQueryParams
	Security    ListCurrentUserOrdersSecurity
}

type ListCurrentUserOrdersResponse struct {
	ContentType                                   string
	PaginationData                                *ListCurrentUserOrdersPaginationData
	StatusCode                                    int64
	ListCurrentUserOrders401ApplicationJSONObject *ListCurrentUserOrders401ApplicationJSON
}
