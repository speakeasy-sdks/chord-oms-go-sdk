package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListStockLocationMovementsPathParams struct {
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type ListStockLocationMovementsQueryParams struct {
	Page    *int64 `queryParam:"style=form,explode=true,name=page"`
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
}

type ListStockLocationMovementsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListStockLocationMovementsRequest struct {
	PathParams  ListStockLocationMovementsPathParams
	QueryParams ListStockLocationMovementsQueryParams
	Security    ListStockLocationMovementsSecurity
}

type ListStockLocationMovements404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListStockLocationMovements401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListStockLocationMovementsPaginationData struct {
	Count          *int64                `json:"count,omitempty"`
	CurrentPage    *int64                `json:"current_page,omitempty"`
	Pages          *int64                `json:"pages,omitempty"`
	PerPage        *int64                `json:"per_page,omitempty"`
	StockMovements *shared.StockMovement `json:"stock_movements,omitempty"`
	TotalCount     *int64                `json:"total_count,omitempty"`
}

type ListStockLocationMovementsResponse struct {
	ContentType                                        string
	PaginationData                                     *ListStockLocationMovementsPaginationData
	StatusCode                                         int
	ListStockLocationMovements401ApplicationJSONObject *ListStockLocationMovements401ApplicationJSON
	ListStockLocationMovements404ApplicationJSONObject *ListStockLocationMovements404ApplicationJSON
}
