package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CaptureOrderPaymentPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
	PaymentID   string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type CaptureOrderPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CaptureOrderPayment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CaptureOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CaptureOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CaptureOrderPaymentRequest struct {
	PathParams CaptureOrderPaymentPathParams
	Security   CaptureOrderPaymentSecurity
}

type CaptureOrderPaymentResponse struct {
	ContentType                                 string
	StatusCode                                  int64
	CaptureOrderPayment401ApplicationJSONObject *CaptureOrderPayment401ApplicationJSON
	CaptureOrderPayment404ApplicationJSONObject *CaptureOrderPayment404ApplicationJSON
	CaptureOrderPayment422ApplicationJSONObject *CaptureOrderPayment422ApplicationJSON
	Payment                                     *shared.Payment
}
