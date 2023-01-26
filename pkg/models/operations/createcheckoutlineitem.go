package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateCheckoutLineItemPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type CreateCheckoutLineItemSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type CreateCheckoutLineItem401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateCheckoutLineItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateCheckoutLineItem422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateCheckoutLineItemRequest struct {
	PathParams CreateCheckoutLineItemPathParams
	Request    shared.LineItemInput `request:"mediaType=application/json"`
	Security   CreateCheckoutLineItemSecurity
}

type CreateCheckoutLineItemResponse struct {
	ContentType                                    string
	StatusCode                                     int64
	CreateCheckoutLineItem401ApplicationJSONObject *CreateCheckoutLineItem401ApplicationJSON
	CreateCheckoutLineItem404ApplicationJSONObject *CreateCheckoutLineItem404ApplicationJSON
	CreateCheckoutLineItem422ApplicationJSONObject *CreateCheckoutLineItem422ApplicationJSON
	LineItem                                       *shared.LineItem
}
