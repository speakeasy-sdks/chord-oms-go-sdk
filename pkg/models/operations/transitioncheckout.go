package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type TransitionCheckoutPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type TransitionCheckoutRequestBody struct {
	ExpectedTotal *string `json:"expected_total,omitempty"`
}

type TransitionCheckoutSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type TransitionCheckout401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type TransitionCheckout404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type TransitionCheckout422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type TransitionCheckoutRequest struct {
	PathParams TransitionCheckoutPathParams
	Request    TransitionCheckoutRequestBody `request:"mediaType=application/json"`
	Security   TransitionCheckoutSecurity
}

type TransitionCheckoutResponse struct {
	ContentType                                string
	StatusCode                                 int64
	OrderBig                                   *shared.OrderBig
	TransitionCheckout401ApplicationJSONObject *TransitionCheckout401ApplicationJSON
	TransitionCheckout404ApplicationJSONObject *TransitionCheckout404ApplicationJSON
	TransitionCheckout422ApplicationJSONObject *TransitionCheckout422ApplicationJSON
}
