package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreditOrderPaymentPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
	PaymentID   string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type CreditOrderPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreditOrderPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreditOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreditOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreditOrderPaymentRequest struct {
	PathParams CreditOrderPaymentPathParams
	Security   CreditOrderPaymentSecurity
}

type CreditOrderPaymentResponse struct {
	ContentType                                string
	StatusCode                                 int64
	CreditOrderPayment401ApplicationJSONObject *CreditOrderPayment401ApplicationJSON
	CreditOrderPayment404ApplicationJSONObject *CreditOrderPayment404ApplicationJSON
	CreditOrderPayment422ApplicationJSONObject *CreditOrderPayment422ApplicationJSON
	Payment                                    *shared.Payment
}
