package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type PurchaseOrderPaymentPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
	PaymentID   string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type PurchaseOrderPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PurchaseOrderPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PurchaseOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PurchaseOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type PurchaseOrderPaymentRequest struct {
	PathParams PurchaseOrderPaymentPathParams
	Security   PurchaseOrderPaymentSecurity
}

type PurchaseOrderPaymentResponse struct {
	ContentType                                  string
	StatusCode                                   int64
	Payment                                      *shared.Payment
	PurchaseOrderPayment401ApplicationJSONObject *PurchaseOrderPayment401ApplicationJSON
	PurchaseOrderPayment404ApplicationJSONObject *PurchaseOrderPayment404ApplicationJSON
	PurchaseOrderPayment422ApplicationJSONObject *PurchaseOrderPayment422ApplicationJSON
}
