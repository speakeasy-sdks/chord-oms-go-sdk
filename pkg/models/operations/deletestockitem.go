package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteStockItemPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteStockItemSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteStockItemRequest struct {
	PathParams DeleteStockItemPathParams
	Security   DeleteStockItemSecurity
}

type DeleteStockItem422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStockItemResponse struct {
	ContentType                             string
	StatusCode                              int
	DeleteStockItem401ApplicationJSONObject *DeleteStockItem401ApplicationJSON
	DeleteStockItem404ApplicationJSONObject *DeleteStockItem404ApplicationJSON
	DeleteStockItem422ApplicationJSONObject *DeleteStockItem422ApplicationJSON
}
