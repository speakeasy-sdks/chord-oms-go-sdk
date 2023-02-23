package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreditCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	PaymentID  string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type CreditCheckoutPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreditCheckoutPaymentRequest struct {
	PathParams CreditCheckoutPaymentPathParams
	Security   CreditCheckoutPaymentSecurity
}

type CreditCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreditCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreditCheckoutPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreditCheckoutPaymentResponse struct {
	ContentType                                   string
	StatusCode                                    int
	CreditCheckoutPayment401ApplicationJSONObject *CreditCheckoutPayment401ApplicationJSON
	CreditCheckoutPayment404ApplicationJSONObject *CreditCheckoutPayment404ApplicationJSON
	CreditCheckoutPayment422ApplicationJSONObject *CreditCheckoutPayment422ApplicationJSON
	Payment                                       *shared.Payment
}
