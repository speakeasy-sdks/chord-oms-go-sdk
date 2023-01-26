package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetStockLocationMovementPathParams struct {
	ID              string `pathParam:"style=simple,explode=false,name=id"`
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type GetStockLocationMovementSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetStockLocationMovement401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetStockLocationMovement404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStockLocationMovementRequest struct {
	PathParams GetStockLocationMovementPathParams
	Security   GetStockLocationMovementSecurity
}

type GetStockLocationMovementResponse struct {
	ContentType                                      string
	StatusCode                                       int64
	GetStockLocationMovement401ApplicationJSONObject *GetStockLocationMovement401ApplicationJSON
	GetStockLocationMovement404ApplicationJSONObject *GetStockLocationMovement404ApplicationJSON
	StockMovement                                    *shared.StockMovement
}
