package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateStockLocationItemPathParams struct {
	ID              string `pathParam:"style=simple,explode=false,name=id"`
	StockLocationID string `pathParam:"style=simple,explode=false,name=stock_location_id"`
}

type UpdateStockLocationItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateStockLocationItemRequest struct {
	PathParams UpdateStockLocationItemPathParams
	Request    shared.StockItemInput `request:"mediaType=application/json"`
	Security   UpdateStockLocationItemSecurity
}

type UpdateStockLocationItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateStockLocationItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStockLocationItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStockLocationItemResponse struct {
	ContentType                                     string
	StatusCode                                      int
	StockItem                                       *shared.StockItem
	UpdateStockLocationItem401ApplicationJSONObject *UpdateStockLocationItem401ApplicationJSON
	UpdateStockLocationItem404ApplicationJSONObject *UpdateStockLocationItem404ApplicationJSON
	UpdateStockLocationItem422ApplicationJSONObject *UpdateStockLocationItem422ApplicationJSON
}
