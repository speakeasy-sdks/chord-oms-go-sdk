package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetStockLocationMovementPathParams struct {
	ID              string `pathParam:"style=simple,explode=false,name=id"`
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type GetStockLocationMovementSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetStockLocationMovementRequest struct {
	PathParams GetStockLocationMovementPathParams
	Security   GetStockLocationMovementSecurity
}

type GetStockLocationMovement404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStockLocationMovement401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStockLocationMovementResponse struct {
	ContentType                                      string
	StatusCode                                       int
	GetStockLocationMovement401ApplicationJSONObject *GetStockLocationMovement401ApplicationJSON
	GetStockLocationMovement404ApplicationJSONObject *GetStockLocationMovement404ApplicationJSON
	StockMovement                                    *shared.StockMovement
}
