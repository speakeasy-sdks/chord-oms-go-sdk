package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
}

type CreateCheckoutPaymentSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type CreateCheckoutPaymentRequest struct {
	PathParams CreateCheckoutPaymentPathParams
	Request    *shared.PaymentInput `request:"mediaType=application/json"`
	Security   CreateCheckoutPaymentSecurity
}

type CreateCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateCheckoutPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateCheckoutPaymentResponse struct {
	ContentType                                   string
	StatusCode                                    int
	CreateCheckoutPayment401ApplicationJSONObject *CreateCheckoutPayment401ApplicationJSON
	CreateCheckoutPayment404ApplicationJSONObject *CreateCheckoutPayment404ApplicationJSON
	CreateCheckoutPayment422ApplicationJSONObject *CreateCheckoutPayment422ApplicationJSON
	Payment                                       *shared.Payment
}
