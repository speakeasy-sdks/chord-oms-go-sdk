package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateStockLocationItemPathParams struct {
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type CreateStockLocationItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateStockLocationItem401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateStockLocationItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateStockLocationItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateStockLocationItemRequest struct {
	PathParams CreateStockLocationItemPathParams
	Request    shared.StockItemInput `request:"mediaType=application/json"`
	Security   CreateStockLocationItemSecurity
}

type CreateStockLocationItemResponse struct {
	ContentType                                     string
	StatusCode                                      int64
	CreateStockLocationItem401ApplicationJSONObject *CreateStockLocationItem401ApplicationJSON
	CreateStockLocationItem404ApplicationJSONObject *CreateStockLocationItem404ApplicationJSON
	CreateStockLocationItem422ApplicationJSONObject *CreateStockLocationItem422ApplicationJSON
	StockItem                                       *shared.StockItem
}
