package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetOrderAddressPathParams struct {
	ID          string `pathParam:"style=simple,explode=false,name=id"`
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type GetOrderAddressSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type GetOrderAddress401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderAddress404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderAddressRequest struct {
	PathParams GetOrderAddressPathParams
	Security   GetOrderAddressSecurity
}

type GetOrderAddressResponse struct {
	ContentType                             string
	StatusCode                              int64
	Address                                 *shared.Address
	GetOrderAddress401ApplicationJSONObject *GetOrderAddress401ApplicationJSON
	GetOrderAddress404ApplicationJSONObject *GetOrderAddress404ApplicationJSON
}
