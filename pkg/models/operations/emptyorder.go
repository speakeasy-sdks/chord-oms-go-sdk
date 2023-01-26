package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type EmptyOrderPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type EmptyOrderSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type EmptyOrder401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type EmptyOrder404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type EmptyOrder422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type EmptyOrderRequest struct {
	PathParams EmptyOrderPathParams
	Security   EmptyOrderSecurity
}

type EmptyOrderResponse struct {
	ContentType                        string
	StatusCode                         int64
	EmptyOrder401ApplicationJSONObject *EmptyOrder401ApplicationJSON
	EmptyOrder404ApplicationJSONObject *EmptyOrder404ApplicationJSON
	EmptyOrder422ApplicationJSONObject *EmptyOrder422ApplicationJSON
}
