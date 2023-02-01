package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type VoidCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	PaymentID  string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type VoidCheckoutPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type VoidCheckoutPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type VoidCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type VoidCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type VoidCheckoutPaymentRequest struct {
	PathParams VoidCheckoutPaymentPathParams
	Security   VoidCheckoutPaymentSecurity
}

type VoidCheckoutPaymentResponse struct {
	ContentType                                 string
	StatusCode                                  int64
	Payment                                     *shared.Payment
	VoidCheckoutPayment401ApplicationJSONObject *VoidCheckoutPayment401ApplicationJSON
	VoidCheckoutPayment404ApplicationJSONObject *VoidCheckoutPayment404ApplicationJSON
	VoidCheckoutPayment422ApplicationJSONObject *VoidCheckoutPayment422ApplicationJSON
}
