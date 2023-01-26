package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateStockItemPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateStockItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateStockItem401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateStockItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStockItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateStockItemRequest struct {
	PathParams UpdateStockItemPathParams
	Request    shared.StockItemInput `request:"mediaType=application/json"`
	Security   UpdateStockItemSecurity
}

type UpdateStockItemResponse struct {
	ContentType                             string
	StatusCode                              int64
	StockItem                               *shared.StockItem
	UpdateStockItem401ApplicationJSONObject *UpdateStockItem401ApplicationJSON
	UpdateStockItem404ApplicationJSONObject *UpdateStockItem404ApplicationJSON
	UpdateStockItem422ApplicationJSONObject *UpdateStockItem422ApplicationJSON
}
