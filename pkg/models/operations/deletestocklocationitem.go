package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteStockLocationItemPathParams struct {
	ID              string `pathParam:"style=simple,explode=false,name=id"`
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type DeleteStockLocationItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteStockLocationItemRequest struct {
	PathParams DeleteStockLocationItemPathParams
	Security   DeleteStockLocationItemSecurity
}

type DeleteStockLocationItem422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockLocationItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockLocationItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockLocationItemResponse struct {
	ContentType                                     string
	StatusCode                                      int
	DeleteStockLocationItem401ApplicationJSONObject *DeleteStockLocationItem401ApplicationJSON
	DeleteStockLocationItem404ApplicationJSONObject *DeleteStockLocationItem404ApplicationJSON
	DeleteStockLocationItem422ApplicationJSONObject *DeleteStockLocationItem422ApplicationJSON
}
