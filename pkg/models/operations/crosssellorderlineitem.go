package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CrossSellOrderLineItemPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type CrossSellOrderLineItemSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type CrossSellOrderLineItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CrossSellOrderLineItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CrossSellOrderLineItem405ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CrossSellOrderLineItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CrossSellOrderLineItemRequest struct {
	PathParams CrossSellOrderLineItemPathParams
	Request    shared.LineItemInput `request:"mediaType=application/json"`
	Security   CrossSellOrderLineItemSecurity
}

type CrossSellOrderLineItemResponse struct {
	ContentType                                    string
	StatusCode                                     int64
	CrossSellOrderLineItem401ApplicationJSONObject *CrossSellOrderLineItem401ApplicationJSON
	CrossSellOrderLineItem404ApplicationJSONObject *CrossSellOrderLineItem404ApplicationJSON
	CrossSellOrderLineItem405ApplicationJSONObject *CrossSellOrderLineItem405ApplicationJSON
	CrossSellOrderLineItem422ApplicationJSONObject *CrossSellOrderLineItem422ApplicationJSON
	LineItem                                       *shared.LineItem
}
