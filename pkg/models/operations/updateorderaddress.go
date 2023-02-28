package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateOrderAddressPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type UpdateOrderAddressSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateOrderAddressRequest struct {
	PathParams UpdateOrderAddressPathParams
	Request    shared.AddressInput `request:"mediaType=application/json"`
	Security   UpdateOrderAddressSecurity
}

type UpdateOrderAddress422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOrderAddress404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderAddress401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderAddressResponse struct {
	ContentType                                string
	StatusCode                                 int
	Address                                    *shared.Address
	UpdateOrderAddress401ApplicationJSONObject *UpdateOrderAddress401ApplicationJSON
	UpdateOrderAddress404ApplicationJSONObject *UpdateOrderAddress404ApplicationJSON
	UpdateOrderAddress422ApplicationJSONObject *UpdateOrderAddress422ApplicationJSON
}
