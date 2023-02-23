package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type PurchaseOrderPaymentPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
	PaymentID   string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type PurchaseOrderPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type PurchaseOrderPaymentRequest struct {
	PathParams PurchaseOrderPaymentPathParams
	Security   PurchaseOrderPaymentSecurity
}

type PurchaseOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type PurchaseOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PurchaseOrderPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PurchaseOrderPaymentResponse struct {
	ContentType                                  string
	StatusCode                                   int
	Payment                                      *shared.Payment
	PurchaseOrderPayment401ApplicationJSONObject *PurchaseOrderPayment401ApplicationJSON
	PurchaseOrderPayment404ApplicationJSONObject *PurchaseOrderPayment404ApplicationJSON
	PurchaseOrderPayment422ApplicationJSONObject *PurchaseOrderPayment422ApplicationJSON
}
