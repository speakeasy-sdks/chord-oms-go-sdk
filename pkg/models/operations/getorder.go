package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetOrderPathParams struct {
	Number string `pathParam:"style=simple,explode=false,name=number"`
}

type GetOrderSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type GetOrder401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrder404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderRequest struct {
	PathParams GetOrderPathParams
	Security   GetOrderSecurity
}

type GetOrderResponse struct {
	ContentType                      string
	StatusCode                       int64
	GetOrder401ApplicationJSONObject *GetOrder401ApplicationJSON
	GetOrder404ApplicationJSONObject *GetOrder404ApplicationJSON
	OrderBig                         *shared.OrderBig
}
