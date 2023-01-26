package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteStockLocationItemPathParams struct {
	ID              string `pathParam:"style=simple,explode=false,name=id"`
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type DeleteStockLocationItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteStockLocationItem401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteStockLocationItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockLocationItem422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockLocationItemRequest struct {
	PathParams DeleteStockLocationItemPathParams
	Security   DeleteStockLocationItemSecurity
}

type DeleteStockLocationItemResponse struct {
	ContentType                                     string
	StatusCode                                      int64
	DeleteStockLocationItem401ApplicationJSONObject *DeleteStockLocationItem401ApplicationJSON
	DeleteStockLocationItem404ApplicationJSONObject *DeleteStockLocationItem404ApplicationJSON
	DeleteStockLocationItem422ApplicationJSONObject *DeleteStockLocationItem422ApplicationJSON
}
