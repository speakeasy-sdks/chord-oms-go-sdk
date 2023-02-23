package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteStockLocationPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteStockLocationSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteStockLocationRequest struct {
	PathParams DeleteStockLocationPathParams
	Security   DeleteStockLocationSecurity
}

type DeleteStockLocation422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockLocation404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockLocation401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockLocationResponse struct {
	ContentType                                 string
	StatusCode                                  int
	DeleteStockLocation401ApplicationJSONObject *DeleteStockLocation401ApplicationJSON
	DeleteStockLocation404ApplicationJSONObject *DeleteStockLocation404ApplicationJSON
	DeleteStockLocation422ApplicationJSONObject *DeleteStockLocation422ApplicationJSON
}
