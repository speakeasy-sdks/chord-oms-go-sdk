package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateStockLocationItemPathParams struct {
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type CreateStockLocationItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateStockLocationItemRequest struct {
	PathParams CreateStockLocationItemPathParams
	Request    shared.StockItemInput `request:"mediaType=application/json"`
	Security   CreateStockLocationItemSecurity
}

type CreateStockLocationItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateStockLocationItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateStockLocationItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateStockLocationItemResponse struct {
	ContentType                                     string
	StatusCode                                      int
	CreateStockLocationItem401ApplicationJSONObject *CreateStockLocationItem401ApplicationJSON
	CreateStockLocationItem404ApplicationJSONObject *CreateStockLocationItem404ApplicationJSON
	CreateStockLocationItem422ApplicationJSONObject *CreateStockLocationItem422ApplicationJSON
	StockItem                                       *shared.StockItem
}
