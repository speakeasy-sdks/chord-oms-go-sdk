package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteStockItemPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteStockItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteStockItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockItem422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockItemRequest struct {
	PathParams DeleteStockItemPathParams
	Security   DeleteStockItemSecurity
}

type DeleteStockItemResponse struct {
	ContentType                             string
	StatusCode                              int64
	DeleteStockItem401ApplicationJSONObject *DeleteStockItem401ApplicationJSON
	DeleteStockItem404ApplicationJSONObject *DeleteStockItem404ApplicationJSON
	DeleteStockItem422ApplicationJSONObject *DeleteStockItem422ApplicationJSON
}
