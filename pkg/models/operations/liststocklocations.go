package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListStockLocationsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListStockLocationsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListStockLocationsPaginationData struct {
	Count          *int64                 `json:"count,omitempty"`
	CurrentPage    *int64                 `json:"current_page,omitempty"`
	Pages          *int64                 `json:"pages,omitempty"`
	PerPage        *int64                 `json:"per_page,omitempty"`
	StockLocations []shared.StockLocation `json:"stock_locations,omitempty"`
	TotalCount     *int64                 `json:"total_count,omitempty"`
}

type ListStockLocations401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListStockLocationsRequest struct {
	QueryParams ListStockLocationsQueryParams
	Security    ListStockLocationsSecurity
}

type ListStockLocationsResponse struct {
	ContentType                                string
	PaginationData                             *ListStockLocationsPaginationData
	StatusCode                                 int64
	ListStockLocations401ApplicationJSONObject *ListStockLocations401ApplicationJSON
}
