package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CaptureOrderPaymentPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
	PaymentID   string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type CaptureOrderPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CaptureOrderPaymentRequest struct {
	PathParams CaptureOrderPaymentPathParams
	Security   CaptureOrderPaymentSecurity
}

type CaptureOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CaptureOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CaptureOrderPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CaptureOrderPaymentResponse struct {
	ContentType                                 string
	StatusCode                                  int
	CaptureOrderPayment401ApplicationJSONObject *CaptureOrderPayment401ApplicationJSON
	CaptureOrderPayment404ApplicationJSONObject *CaptureOrderPayment404ApplicationJSON
	CaptureOrderPayment422ApplicationJSONObject *CaptureOrderPayment422ApplicationJSON
	Payment                                     *shared.Payment
}
