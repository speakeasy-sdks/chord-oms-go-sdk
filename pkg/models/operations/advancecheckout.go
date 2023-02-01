package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type AdvanceCheckoutPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type AdvanceCheckoutSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type AdvanceCheckout401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type AdvanceCheckout404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type AdvanceCheckout422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type AdvanceCheckoutRequest struct {
	PathParams AdvanceCheckoutPathParams
	Security   AdvanceCheckoutSecurity
}

type AdvanceCheckoutResponse struct {
	ContentType                             string
	StatusCode                              int64
	AdvanceCheckout401ApplicationJSONObject *AdvanceCheckout401ApplicationJSON
	AdvanceCheckout404ApplicationJSONObject *AdvanceCheckout404ApplicationJSON
	AdvanceCheckout422ApplicationJSONObject *AdvanceCheckout422ApplicationJSON
	OrderBig                                *shared.OrderBig
}
