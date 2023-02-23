package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type TransitionCheckoutPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type TransitionCheckoutRequestBody struct {
	ExpectedTotal *string `json:"expected_total,omitempty"`
}

type TransitionCheckoutSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type TransitionCheckoutRequest struct {
	PathParams TransitionCheckoutPathParams
	Request    TransitionCheckoutRequestBody `request:"mediaType=application/json"`
	Security   TransitionCheckoutSecurity
}

type TransitionCheckout422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type TransitionCheckout404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type TransitionCheckout401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type TransitionCheckoutResponse struct {
	ContentType                                string
	OrderBig                                   *shared.OrderBig
	StatusCode                                 int
	TransitionCheckout401ApplicationJSONObject *TransitionCheckout401ApplicationJSON
	TransitionCheckout404ApplicationJSONObject *TransitionCheckout404ApplicationJSON
	TransitionCheckout422ApplicationJSONObject *TransitionCheckout422ApplicationJSON
}
