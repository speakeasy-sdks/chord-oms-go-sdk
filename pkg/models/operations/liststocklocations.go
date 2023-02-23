package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListStockLocationsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListStockLocationsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListStockLocationsRequest struct {
	QueryParams ListStockLocationsQueryParams
	Security    ListStockLocationsSecurity
}

type ListStockLocations401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListStockLocationsPaginationData struct {
	Count          *int64                 `json:"count,omitempty"`
	CurrentPage    *int64                 `json:"current_page,omitempty"`
	Pages          *int64                 `json:"pages,omitempty"`
	PerPage        *int64                 `json:"per_page,omitempty"`
	StockLocations []shared.StockLocation `json:"stock_locations,omitempty"`
	TotalCount     *int64                 `json:"total_count,omitempty"`
}

type ListStockLocationsResponse struct {
	ContentType                                string
	PaginationData                             *ListStockLocationsPaginationData
	StatusCode                                 int
	ListStockLocations401ApplicationJSONObject *ListStockLocations401ApplicationJSON
}
