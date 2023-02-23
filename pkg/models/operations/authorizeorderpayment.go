package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type AuthorizeOrderPaymentPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
	PaymentID   string `pathParam:"style=simple,explode=false,name=payment_id"`
}

type AuthorizeOrderPaymentSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type AuthorizeOrderPaymentRequest struct {
	PathParams AuthorizeOrderPaymentPathParams
	Security   AuthorizeOrderPaymentSecurity
}

type AuthorizeOrderPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type AuthorizeOrderPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type AuthorizeOrderPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type AuthorizeOrderPaymentResponse struct {
	ContentType                                   string
	StatusCode                                    int
	AuthorizeOrderPayment401ApplicationJSONObject *AuthorizeOrderPayment401ApplicationJSON
	AuthorizeOrderPayment404ApplicationJSONObject *AuthorizeOrderPayment404ApplicationJSON
	AuthorizeOrderPayment422ApplicationJSONObject *AuthorizeOrderPayment422ApplicationJSON
	Payment                                       *shared.Payment
}
