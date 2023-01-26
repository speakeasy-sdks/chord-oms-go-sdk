package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreditCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	PaymentID  string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type CreditCheckoutPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreditCheckoutPayment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreditCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreditCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreditCheckoutPaymentRequest struct {
	PathParams CreditCheckoutPaymentPathParams
	Security   CreditCheckoutPaymentSecurity
}

type CreditCheckoutPaymentResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	CreditCheckoutPayment401ApplicationJSONObject *CreditCheckoutPayment401ApplicationJSON
	CreditCheckoutPayment404ApplicationJSONObject *CreditCheckoutPayment404ApplicationJSON
	CreditCheckoutPayment422ApplicationJSONObject *CreditCheckoutPayment422ApplicationJSON
	Payment                                       *shared.Payment
}
