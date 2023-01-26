package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateOrderPaymentPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type CreateOrderPaymentSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOrderPayment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOrderPaymentRequest struct {
	PathParams CreateOrderPaymentPathParams
	Request    shared.PaymentInput `request:"mediaType=application/json"`
	Security   CreateOrderPaymentSecurity
}

type CreateOrderPaymentResponse struct {
	ContentType                                string
	StatusCode                                 int64
	CreateOrderPayment401ApplicationJSONObject *CreateOrderPayment401ApplicationJSON
	CreateOrderPayment404ApplicationJSONObject *CreateOrderPayment404ApplicationJSON
	CreateOrderPayment422ApplicationJSONObject *CreateOrderPayment422ApplicationJSON
	Payment                                    *shared.Payment
}
