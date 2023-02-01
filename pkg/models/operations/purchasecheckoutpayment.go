package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type PurchaseCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	PaymentID  string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type PurchaseCheckoutPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PurchaseCheckoutPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PurchaseCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PurchaseCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type PurchaseCheckoutPaymentRequest struct {
	PathParams PurchaseCheckoutPaymentPathParams
	Security   PurchaseCheckoutPaymentSecurity
}

type PurchaseCheckoutPaymentResponse struct {
	ContentType                                     string
	StatusCode                                      int64
	Payment                                         *shared.Payment
	PurchaseCheckoutPayment401ApplicationJSONObject *PurchaseCheckoutPayment401ApplicationJSON
	PurchaseCheckoutPayment404ApplicationJSONObject *PurchaseCheckoutPayment404ApplicationJSON
	PurchaseCheckoutPayment422ApplicationJSONObject *PurchaseCheckoutPayment422ApplicationJSON
}
