package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateCheckoutAddressPathParams struct {
	CheckoutID string `pathParam:"style=simple,explode=false,name=checkout_id"`
	ID         string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateCheckoutAddressSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateCheckoutAddress401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateCheckoutAddress404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateCheckoutAddress422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateCheckoutAddressRequest struct {
	PathParams UpdateCheckoutAddressPathParams
	Request    *shared.AddressInput `request:"mediaType=application/json"`
	Security   UpdateCheckoutAddressSecurity
}

type UpdateCheckoutAddressResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	Address                                       *shared.Address
	UpdateCheckoutAddress401ApplicationJSONObject *UpdateCheckoutAddress401ApplicationJSON
	UpdateCheckoutAddress404ApplicationJSONObject *UpdateCheckoutAddress404ApplicationJSON
	UpdateCheckoutAddress422ApplicationJSONObject *UpdateCheckoutAddress422ApplicationJSON
}
