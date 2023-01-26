package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CaptureCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	PaymentID  string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type CaptureCheckoutPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CaptureCheckoutPayment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CaptureCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CaptureCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CaptureCheckoutPaymentRequest struct {
	PathParams CaptureCheckoutPaymentPathParams
	Security   CaptureCheckoutPaymentSecurity
}

type CaptureCheckoutPaymentResponse struct {
	ContentType                                    string
	StatusCode                                     int64
	CaptureCheckoutPayment401ApplicationJSONObject *CaptureCheckoutPayment401ApplicationJSON
	CaptureCheckoutPayment404ApplicationJSONObject *CaptureCheckoutPayment404ApplicationJSON
	CaptureCheckoutPayment422ApplicationJSONObject *CaptureCheckoutPayment422ApplicationJSON
	Payment                                        *shared.Payment
}
