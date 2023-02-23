package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type EmptyOrderPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type EmptyOrderSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type EmptyOrderRequest struct {
	PathParams EmptyOrderPathParams
	Security   EmptyOrderSecurity
}

type EmptyOrder422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type EmptyOrder404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type EmptyOrder401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type EmptyOrderResponse struct {
	ContentType                        string
	StatusCode                         int
	EmptyOrder401ApplicationJSONObject *EmptyOrder401ApplicationJSON
	EmptyOrder404ApplicationJSONObject *EmptyOrder404ApplicationJSON
	EmptyOrder422ApplicationJSONObject *EmptyOrder422ApplicationJSON
}
