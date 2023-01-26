package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateStockLocationMovementPathParams struct {
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type CreateStockLocationMovementSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateStockLocationMovement401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateStockLocationMovement404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateStockLocationMovement422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateStockLocationMovementRequest struct {
	PathParams CreateStockLocationMovementPathParams
	Request    shared.StockMovementInput `request:"mediaType=application/json"`
	Security   CreateStockLocationMovementSecurity
}

type CreateStockLocationMovementResponse struct {
	ContentType                                         string
	StatusCode                                          int64
	CreateStockLocationMovement401ApplicationJSONObject *CreateStockLocationMovement401ApplicationJSON
	CreateStockLocationMovement404ApplicationJSONObject *CreateStockLocationMovement404ApplicationJSON
	CreateStockLocationMovement422ApplicationJSONObject *CreateStockLocationMovement422ApplicationJSON
	StockMovement                                       *shared.StockMovement
}
