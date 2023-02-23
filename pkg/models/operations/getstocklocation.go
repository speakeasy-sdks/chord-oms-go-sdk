package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetStockLocationPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetStockLocationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetStockLocationRequest struct {
	PathParams GetStockLocationPathParams
	Security   GetStockLocationSecurity
}

type GetStockLocation404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStockLocation401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStockLocationResponse struct {
	ContentType                              string
	StatusCode                               int
	GetStockLocation401ApplicationJSONObject *GetStockLocation401ApplicationJSON
	GetStockLocation404ApplicationJSONObject *GetStockLocation404ApplicationJSON
	StockLocation                            *shared.StockLocation
}
