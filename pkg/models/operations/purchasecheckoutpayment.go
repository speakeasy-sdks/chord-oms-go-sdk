package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type PurchaseCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	PaymentID  string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type PurchaseCheckoutPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type PurchaseCheckoutPaymentRequest struct {
	PathParams PurchaseCheckoutPaymentPathParams
	Security   PurchaseCheckoutPaymentSecurity
}

type PurchaseCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type PurchaseCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PurchaseCheckoutPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PurchaseCheckoutPaymentResponse struct {
	ContentType                                     string
	StatusCode                                      int
	Payment                                         *shared.Payment
	PurchaseCheckoutPayment401ApplicationJSONObject *PurchaseCheckoutPayment401ApplicationJSON
	PurchaseCheckoutPayment404ApplicationJSONObject *PurchaseCheckoutPayment404ApplicationJSON
	PurchaseCheckoutPayment422ApplicationJSONObject *PurchaseCheckoutPayment422ApplicationJSON
}
