package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetOrderPaymentPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type GetOrderPaymentSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type GetOrderPayment401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderPaymentRequest struct {
	PathParams GetOrderPaymentPathParams
	Security   GetOrderPaymentSecurity
}

type GetOrderPaymentResponse struct {
	ContentType                             string
	StatusCode                              int64
	GetOrderPayment401ApplicationJSONObject *GetOrderPayment401ApplicationJSON
	GetOrderPayment404ApplicationJSONObject *GetOrderPayment404ApplicationJSON
	Payment                                 *shared.Payment
}
