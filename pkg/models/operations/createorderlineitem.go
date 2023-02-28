package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateOrderLineItemPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type CreateOrderLineItemSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	OrderToken      *shared.SchemeOrderToken      `security:"scheme,type=apiKey,subtype=header"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOrderLineItemRequest struct {
	PathParams CreateOrderLineItemPathParams
	Request    shared.AddLineItemInput `request:"mediaType=application/json"`
	Security   CreateOrderLineItemSecurity
}

type CreateOrderLineItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOrderLineItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderLineItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderLineItemResponse struct {
	ContentType                                 string
	StatusCode                                  int
	CreateOrderLineItem401ApplicationJSONObject *CreateOrderLineItem401ApplicationJSON
	CreateOrderLineItem404ApplicationJSONObject *CreateOrderLineItem404ApplicationJSON
	CreateOrderLineItem422ApplicationJSONObject *CreateOrderLineItem422ApplicationJSON
	LineItem                                    *shared.LineItem
}
