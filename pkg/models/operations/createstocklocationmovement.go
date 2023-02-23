package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateStockLocationMovementPathParams struct {
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type CreateStockLocationMovementSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateStockLocationMovementRequest struct {
	PathParams CreateStockLocationMovementPathParams
	Request    shared.StockMovementInput `request:"mediaType=application/json"`
	Security   CreateStockLocationMovementSecurity
}

type CreateStockLocationMovement422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateStockLocationMovement404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateStockLocationMovement401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateStockLocationMovementResponse struct {
	ContentType                                         string
	StatusCode                                          int
	CreateStockLocationMovement401ApplicationJSONObject *CreateStockLocationMovement401ApplicationJSON
	CreateStockLocationMovement404ApplicationJSONObject *CreateStockLocationMovement404ApplicationJSON
	CreateStockLocationMovement422ApplicationJSONObject *CreateStockLocationMovement422ApplicationJSON
	StockMovement                                       *shared.StockMovement
}
