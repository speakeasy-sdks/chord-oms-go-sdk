package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateStockLocationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateStockLocation401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateStockLocation422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateStockLocationRequest struct {
	Request  shared.StockLocationInput `request:"mediaType=application/json"`
	Security CreateStockLocationSecurity
}

type CreateStockLocationResponse struct {
	ContentType                                 string
	StatusCode                                  int64
	CreateStockLocation401ApplicationJSONObject *CreateStockLocation401ApplicationJSON
	CreateStockLocation422ApplicationJSONObject *CreateStockLocation422ApplicationJSON
	StockLocation                               *shared.StockLocation
}
