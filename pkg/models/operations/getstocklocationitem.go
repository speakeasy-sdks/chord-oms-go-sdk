package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetStockLocationItemPathParams struct {
	ID              string `pathParam:"style=simple,explode=false,name=id"`
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type GetStockLocationItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetStockLocationItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStockLocationItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStockLocationItemRequest struct {
	PathParams GetStockLocationItemPathParams
	Security   GetStockLocationItemSecurity
}

type GetStockLocationItemResponse struct {
	ContentType                                  string
	StatusCode                                   int64
	GetStockLocationItem401ApplicationJSONObject *GetStockLocationItem401ApplicationJSON
	GetStockLocationItem404ApplicationJSONObject *GetStockLocationItem404ApplicationJSON
	StockItem                                    *shared.StockItem
}
