package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type VoidOrderPaymentPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
	PaymentID   string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type VoidOrderPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type VoidOrderPayment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type VoidOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type VoidOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type VoidOrderPaymentRequest struct {
	PathParams VoidOrderPaymentPathParams
	Security   VoidOrderPaymentSecurity
}

type VoidOrderPaymentResponse struct {
	ContentType                              string
	StatusCode                               int64
	Payment                                  *shared.Payment
	VoidOrderPayment401ApplicationJSONObject *VoidOrderPayment401ApplicationJSON
	VoidOrderPayment404ApplicationJSONObject *VoidOrderPayment404ApplicationJSON
	VoidOrderPayment422ApplicationJSONObject *VoidOrderPayment422ApplicationJSON
}
