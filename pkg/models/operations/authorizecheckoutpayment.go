package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type AuthorizeCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	PaymentID  string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type AuthorizeCheckoutPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type AuthorizeCheckoutPaymentRequest struct {
	PathParams AuthorizeCheckoutPaymentPathParams
	Security   AuthorizeCheckoutPaymentSecurity
}

type AuthorizeCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type AuthorizeCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type AuthorizeCheckoutPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type AuthorizeCheckoutPaymentResponse struct {
	ContentType                                      string
	StatusCode                                       int
	AuthorizeCheckoutPayment401ApplicationJSONObject *AuthorizeCheckoutPayment401ApplicationJSON
	AuthorizeCheckoutPayment404ApplicationJSONObject *AuthorizeCheckoutPayment404ApplicationJSON
	AuthorizeCheckoutPayment422ApplicationJSONObject *AuthorizeCheckoutPayment422ApplicationJSON
	Payment                                          *shared.Payment
}
