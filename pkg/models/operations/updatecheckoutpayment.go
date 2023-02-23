package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateCheckoutPaymentRequestBody struct {
	Amount          *float64              `json:"amount,omitempty"`
	PaymentMethod   *shared.PaymentMethod `json:"payment_method,omitempty"`
	PaymentMethodID *int64                `json:"payment_method_id,omitempty"`
}

type UpdateCheckoutPaymentSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateCheckoutPaymentRequest struct {
	PathParams UpdateCheckoutPaymentPathParams
	Request    UpdateCheckoutPaymentRequestBody `request:"mediaType=application/json"`
	Security   UpdateCheckoutPaymentSecurity
}

type UpdateCheckoutPayment422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCheckoutPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCheckoutPaymentResponse struct {
	ContentType                                   string
	StatusCode                                    int
	Payment                                       *shared.Payment
	UpdateCheckoutPayment401ApplicationJSONObject *UpdateCheckoutPayment401ApplicationJSON
	UpdateCheckoutPayment404ApplicationJSONObject *UpdateCheckoutPayment404ApplicationJSON
	UpdateCheckoutPayment422ApplicationJSONObject *UpdateCheckoutPayment422ApplicationJSON
}
