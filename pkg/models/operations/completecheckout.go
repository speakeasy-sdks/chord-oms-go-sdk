package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CompleteCheckoutPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type CompleteCheckoutRequestBody struct {
	ExpectedTotal *string `json:"expected_total,omitempty"`
}

type CompleteCheckoutSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type CompleteCheckoutRequest struct {
	PathParams CompleteCheckoutPathParams
	Request    CompleteCheckoutRequestBody `request:"mediaType=application/json"`
	Security   CompleteCheckoutSecurity
}

type CompleteCheckout422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CompleteCheckout404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CompleteCheckout401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CompleteCheckoutResponse struct {
	ContentType                              string
	OrderBig                                 *shared.OrderBig
	StatusCode                               int
	CompleteCheckout401ApplicationJSONObject *CompleteCheckout401ApplicationJSON
	CompleteCheckout404ApplicationJSONObject *CompleteCheckout404ApplicationJSON
	CompleteCheckout422ApplicationJSONObject *CompleteCheckout422ApplicationJSON
}
