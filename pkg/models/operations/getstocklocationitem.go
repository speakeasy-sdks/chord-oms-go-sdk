package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetStockLocationItemPathParams struct {
	ID              string `pathParam:"style=simple,explode=false,name=id"`
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type GetStockLocationItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetStockLocationItemRequest struct {
	PathParams GetStockLocationItemPathParams
	Security   GetStockLocationItemSecurity
}

type GetStockLocationItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStockLocationItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStockLocationItemResponse struct {
	ContentType                                  string
	StatusCode                                   int
	GetStockLocationItem401ApplicationJSONObject *GetStockLocationItem401ApplicationJSON
	GetStockLocationItem404ApplicationJSONObject *GetStockLocationItem404ApplicationJSON
	StockItem                                    *shared.StockItem
}
