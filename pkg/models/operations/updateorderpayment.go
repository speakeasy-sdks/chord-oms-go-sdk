package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateOrderPaymentPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type UpdateOrderPaymentSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateOrderPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOrderPaymentRequest struct {
	PathParams UpdateOrderPaymentPathParams
	Request    shared.PaymentInput `request:"mediaType=application/json"`
	Security   UpdateOrderPaymentSecurity
}

type UpdateOrderPaymentResponse struct {
	ContentType                                string
	StatusCode                                 int64
	Payment                                    *shared.Payment
	UpdateOrderPayment401ApplicationJSONObject *UpdateOrderPayment401ApplicationJSON
	UpdateOrderPayment404ApplicationJSONObject *UpdateOrderPayment404ApplicationJSON
	UpdateOrderPayment422ApplicationJSONObject *UpdateOrderPayment422ApplicationJSON
}
