package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateStockItemPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateStockItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateStockItemRequest struct {
	PathParams UpdateStockItemPathParams
	Request    shared.StockItemInput `request:"mediaType=application/json"`
	Security   UpdateStockItemSecurity
}

type UpdateStockItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateStockItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStockItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStockItemResponse struct {
	ContentType                             string
	StatusCode                              int
	StockItem                               *shared.StockItem
	UpdateStockItem401ApplicationJSONObject *UpdateStockItem401ApplicationJSON
	UpdateStockItem404ApplicationJSONObject *UpdateStockItem404ApplicationJSON
	UpdateStockItem422ApplicationJSONObject *UpdateStockItem422ApplicationJSON
}
