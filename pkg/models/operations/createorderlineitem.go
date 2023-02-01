package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateOrderLineItemPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type CreateOrderLineItemSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOrderLineItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderLineItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderLineItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOrderLineItemRequest struct {
	PathParams CreateOrderLineItemPathParams
	Request    shared.LineItemInput `request:"mediaType=application/json"`
	Security   CreateOrderLineItemSecurity
}

type CreateOrderLineItemResponse struct {
	ContentType                                 string
	StatusCode                                  int64
	CreateOrderLineItem401ApplicationJSONObject *CreateOrderLineItem401ApplicationJSON
	CreateOrderLineItem404ApplicationJSONObject *CreateOrderLineItem404ApplicationJSON
	CreateOrderLineItem422ApplicationJSONObject *CreateOrderLineItem422ApplicationJSON
	LineItem                                    *shared.LineItem
}
