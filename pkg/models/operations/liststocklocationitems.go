package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListStockLocationItemsPathParams struct {
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type ListStockLocationItemsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListStockLocationItemsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListStockLocationItemsPaginationData struct {
	Count       *int64             `json:"count,omitempty"`
	CurrentPage *int64             `json:"current_page,omitempty"`
	Pages       *int64             `json:"pages,omitempty"`
	PerPage     *int64             `json:"per_page,omitempty"`
	StockItems  []shared.StockItem `json:"stock_items,omitempty"`
	TotalCount  *int64             `json:"total_count,omitempty"`
}

type ListStockLocationItems401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListStockLocationItems404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListStockLocationItemsRequest struct {
	PathParams  ListStockLocationItemsPathParams
	QueryParams ListStockLocationItemsQueryParams
	Security    ListStockLocationItemsSecurity
}

type ListStockLocationItemsResponse struct {
	ContentType                                    string
	PaginationData                                 *ListStockLocationItemsPaginationData
	StatusCode                                     int64
	ListStockLocationItems401ApplicationJSONObject *ListStockLocationItems401ApplicationJSON
	ListStockLocationItems404ApplicationJSONObject *ListStockLocationItems404ApplicationJSON
}
