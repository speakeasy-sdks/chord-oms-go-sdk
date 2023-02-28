package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateStockLocationPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateStockLocationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateStockLocationRequest struct {
	PathParams UpdateStockLocationPathParams
	Request    shared.StockLocationInput `request:"mediaType=application/json"`
	Security   UpdateStockLocationSecurity
}

type UpdateStockLocation422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateStockLocation404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStockLocation401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStockLocationResponse struct {
	ContentType                                 string
	StatusCode                                  int
	StockLocation                               *shared.StockLocation
	UpdateStockLocation401ApplicationJSONObject *UpdateStockLocation401ApplicationJSON
	UpdateStockLocation404ApplicationJSONObject *UpdateStockLocation404ApplicationJSON
	UpdateStockLocation422ApplicationJSONObject *UpdateStockLocation422ApplicationJSON
}
