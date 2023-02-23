package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListCurrentUserOrdersQueryParams struct {
	Page    *int64  `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64  `queryParam:"style=form,explode=true,name=per_page"`
	Q       *string `queryParam:"style=form,explode=true,name=q"`
}

type ListCurrentUserOrdersSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type ListCurrentUserOrdersRequest struct {
	QueryParams ListCurrentUserOrdersQueryParams
	Security    ListCurrentUserOrdersSecurity
}

type ListCurrentUserOrders500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type ListCurrentUserOrders401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListCurrentUserOrdersPaginationData struct {
	Count       *int64              `json:"count,omitempty"`
	CurrentPage *int64              `json:"current_page,omitempty"`
	Orders      []shared.OrderSmall `json:"orders,omitempty"`
	Pages       *int64              `json:"pages,omitempty"`
	PerPage     *int64              `json:"per_page,omitempty"`
	TotalCount  *int64              `json:"total_count,omitempty"`
}

type ListCurrentUserOrdersResponse struct {
	ContentType                                   string
	PaginationData                                *ListCurrentUserOrdersPaginationData
	StatusCode                                    int
	ListCurrentUserOrders401ApplicationJSONObject *ListCurrentUserOrders401ApplicationJSON
	ListCurrentUserOrders500ApplicationJSONObject *ListCurrentUserOrders500ApplicationJSON
}
