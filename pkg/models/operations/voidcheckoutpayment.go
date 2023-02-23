package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type VoidCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	PaymentID  string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type VoidCheckoutPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type VoidCheckoutPaymentRequest struct {
	PathParams VoidCheckoutPaymentPathParams
	Security   VoidCheckoutPaymentSecurity
}

type VoidCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type VoidCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type VoidCheckoutPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type VoidCheckoutPaymentResponse struct {
	ContentType                                 string
	StatusCode                                  int
	Payment                                     *shared.Payment
	VoidCheckoutPayment401ApplicationJSONObject *VoidCheckoutPayment401ApplicationJSON
	VoidCheckoutPayment404ApplicationJSONObject *VoidCheckoutPayment404ApplicationJSON
	VoidCheckoutPayment422ApplicationJSONObject *VoidCheckoutPayment422ApplicationJSON
}
