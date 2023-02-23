package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreditOrderPaymentPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
	PaymentID   string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type CreditOrderPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreditOrderPaymentRequest struct {
	PathParams CreditOrderPaymentPathParams
	Security   CreditOrderPaymentSecurity
}

type CreditOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreditOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreditOrderPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreditOrderPaymentResponse struct {
	ContentType                                string
	StatusCode                                 int
	CreditOrderPayment401ApplicationJSONObject *CreditOrderPayment401ApplicationJSON
	CreditOrderPayment404ApplicationJSONObject *CreditOrderPayment404ApplicationJSON
	CreditOrderPayment422ApplicationJSONObject *CreditOrderPayment422ApplicationJSON
	Payment                                    *shared.Payment
}
