package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateOrderPaymentPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type UpdateOrderPaymentSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateOrderPaymentRequest struct {
	PathParams UpdateOrderPaymentPathParams
	Request    shared.PaymentInput `request:"mediaType=application/json"`
	Security   UpdateOrderPaymentSecurity
}

type UpdateOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderPaymentResponse struct {
	ContentType                                string
	StatusCode                                 int
	Payment                                    *shared.Payment
	UpdateOrderPayment401ApplicationJSONObject *UpdateOrderPayment401ApplicationJSON
	UpdateOrderPayment404ApplicationJSONObject *UpdateOrderPayment404ApplicationJSON
	UpdateOrderPayment422ApplicationJSONObject *UpdateOrderPayment422ApplicationJSON
}
