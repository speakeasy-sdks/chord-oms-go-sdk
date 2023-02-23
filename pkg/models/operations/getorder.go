package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetOrderPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type GetOrderSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	OrderToken      *shared.SchemeOrderToken      `security:"scheme,type=apiKey,subtype=header"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type GetOrderRequest struct {
	PathParams GetOrderPathParams
	Security   GetOrderSecurity
}

type GetOrder500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type GetOrder404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrder401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOrderResponse struct {
	ContentType                      string
	OrderBig                         *shared.OrderBig
	StatusCode                       int
	GetOrder401ApplicationJSONObject *GetOrder401ApplicationJSON
	GetOrder404ApplicationJSONObject *GetOrder404ApplicationJSON
	GetOrder500ApplicationJSONObject *GetOrder500ApplicationJSON
}
