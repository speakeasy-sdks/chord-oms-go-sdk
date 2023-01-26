package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateCheckoutLineItemPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateCheckoutLineItemRequestBodyLineItemsAttributes struct {
	ID *int64 `json:"id,omitempty"`
}

type UpdateCheckoutLineItemRequestBody struct {
	LineItemsAttributes *UpdateCheckoutLineItemRequestBodyLineItemsAttributes `json:"line_items_attributes,omitempty"`
}

type UpdateCheckoutLineItemSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateCheckoutLineItem401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateCheckoutLineItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCheckoutLineItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateCheckoutLineItemRequest struct {
	PathParams UpdateCheckoutLineItemPathParams
	Request    UpdateCheckoutLineItemRequestBody `request:"mediaType=application/json"`
	Security   UpdateCheckoutLineItemSecurity
}

type UpdateCheckoutLineItemResponse struct {
	ContentType                                    string
	StatusCode                                     int64
	LineItem                                       *shared.LineItem
	UpdateCheckoutLineItem401ApplicationJSONObject *UpdateCheckoutLineItem401ApplicationJSON
	UpdateCheckoutLineItem404ApplicationJSONObject *UpdateCheckoutLineItem404ApplicationJSON
	UpdateCheckoutLineItem422ApplicationJSONObject *UpdateCheckoutLineItem422ApplicationJSON
}
