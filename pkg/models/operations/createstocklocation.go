package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateStockLocationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateStockLocationRequest struct {
	Request  shared.StockLocationInput `request:"mediaType=application/json"`
	Security CreateStockLocationSecurity
}

type CreateStockLocation422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateStockLocation401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateStockLocationResponse struct {
	ContentType                                 string
	StatusCode                                  int
	CreateStockLocation401ApplicationJSONObject *CreateStockLocation401ApplicationJSON
	CreateStockLocation422ApplicationJSONObject *CreateStockLocation422ApplicationJSON
	StockLocation                               *shared.StockLocation
}
