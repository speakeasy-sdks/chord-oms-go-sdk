package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CancelOrderPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type CancelOrderSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=http,subtype=bearer"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type CancelOrderRequest struct {
	PathParams CancelOrderPathParams
	Security   CancelOrderSecurity
}

type CancelOrder422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CancelOrder404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CancelOrder401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CancelOrderResponse struct {
	ContentType                         string
	OrderBig                            *shared.OrderBig
	StatusCode                          int
	CancelOrder401ApplicationJSONObject *CancelOrder401ApplicationJSON
	CancelOrder404ApplicationJSONObject *CancelOrder404ApplicationJSON
	CancelOrder422ApplicationJSONObject *CancelOrder422ApplicationJSON
}
