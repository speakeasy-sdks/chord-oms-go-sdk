package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetCheckoutPaymentPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCheckoutPaymentSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type GetCheckoutPayment401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCheckoutPayment404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCheckoutPaymentRequest struct {
	PathParams GetCheckoutPaymentPathParams
	Security   GetCheckoutPaymentSecurity
}

type GetCheckoutPaymentResponse struct {
	ContentType                                string
	StatusCode                                 int64
	GetCheckoutPayment401ApplicationJSONObject *GetCheckoutPayment401ApplicationJSON
	GetCheckoutPayment404ApplicationJSONObject *GetCheckoutPayment404ApplicationJSON
	Payment                                    *shared.Payment
}
