package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetOrderAddressPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type GetOrderAddressSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type GetOrderAddressRequest struct {
	PathParams GetOrderAddressPathParams
	Security   GetOrderAddressSecurity
}

type GetOrderAddress404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderAddress401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderAddressResponse struct {
	ContentType                             string
	StatusCode                              int
	Address                                 *shared.Address
	GetOrderAddress401ApplicationJSONObject *GetOrderAddress401ApplicationJSON
	GetOrderAddress404ApplicationJSONObject *GetOrderAddress404ApplicationJSON
}
