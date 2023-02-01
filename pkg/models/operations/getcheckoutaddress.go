package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetCheckoutAddressPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type GetCheckoutAddressSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type GetCheckoutAddress401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCheckoutAddress404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCheckoutAddressRequest struct {
	PathParams GetCheckoutAddressPathParams
	Security   GetCheckoutAddressSecurity
}

type GetCheckoutAddressResponse struct {
	ContentType                                string
	StatusCode                                 int64
	Address                                    *shared.Address
	GetCheckoutAddress401ApplicationJSONObject *GetCheckoutAddress401ApplicationJSON
	GetCheckoutAddress404ApplicationJSONObject *GetCheckoutAddress404ApplicationJSON
}
