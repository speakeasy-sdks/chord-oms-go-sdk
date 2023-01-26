package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteCheckoutLineItemPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteCheckoutLineItemSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteCheckoutLineItem401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteCheckoutLineItem404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteCheckoutLineItem422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteCheckoutLineItemRequest struct {
	PathParams DeleteCheckoutLineItemPathParams
	Security   DeleteCheckoutLineItemSecurity
}

type DeleteCheckoutLineItemResponse struct {
	ContentType                                    string
	StatusCode                                     int64
	DeleteCheckoutLineItem401ApplicationJSONObject *DeleteCheckoutLineItem401ApplicationJSON
	DeleteCheckoutLineItem404ApplicationJSONObject *DeleteCheckoutLineItem404ApplicationJSON
	DeleteCheckoutLineItem422ApplicationJSONObject *DeleteCheckoutLineItem422ApplicationJSON
}
