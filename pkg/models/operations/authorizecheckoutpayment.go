package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type AuthorizeCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	PaymentID  string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type AuthorizeCheckoutPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type AuthorizeCheckoutPayment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type AuthorizeCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type AuthorizeCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type AuthorizeCheckoutPaymentRequest struct {
	PathParams AuthorizeCheckoutPaymentPathParams
	Security   AuthorizeCheckoutPaymentSecurity
}

type AuthorizeCheckoutPaymentResponse struct {
	ContentType                                      string
	StatusCode                                       int64
	AuthorizeCheckoutPayment401ApplicationJSONObject *AuthorizeCheckoutPayment401ApplicationJSON
	AuthorizeCheckoutPayment404ApplicationJSONObject *AuthorizeCheckoutPayment404ApplicationJSON
	AuthorizeCheckoutPayment422ApplicationJSONObject *AuthorizeCheckoutPayment422ApplicationJSON
	Payment                                          *shared.Payment
}
