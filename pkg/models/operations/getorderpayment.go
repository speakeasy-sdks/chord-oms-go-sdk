package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetOrderPaymentPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type GetOrderPaymentSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type GetOrderPaymentRequest struct {
	PathParams GetOrderPaymentPathParams
	Security   GetOrderPaymentSecurity
}

type GetOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderPaymentResponse struct {
	ContentType                             string
	StatusCode                              int
	GetOrderPayment401ApplicationJSONObject *GetOrderPayment401ApplicationJSON
	GetOrderPayment404ApplicationJSONObject *GetOrderPayment404ApplicationJSON
	Payment                                 *shared.Payment
}
