package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateOrderLineItemPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type UpdateOrderLineItemSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateOrderLineItemRequest struct {
	PathParams UpdateOrderLineItemPathParams
	Request    shared.LineItemInput `request:"mediaType=application/json"`
	Security   UpdateOrderLineItemSecurity
}

type UpdateOrderLineItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOrderLineItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderLineItem401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderLineItemResponse struct {
	ContentType                                 string
	StatusCode                                  int
	LineItem                                    *shared.LineItem
	UpdateOrderLineItem401ApplicationJSONObject *UpdateOrderLineItem401ApplicationJSON
	UpdateOrderLineItem404ApplicationJSONObject *UpdateOrderLineItem404ApplicationJSON
	UpdateOrderLineItem422ApplicationJSONObject *UpdateOrderLineItem422ApplicationJSON
}
