package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CancelOrderPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type CancelOrderSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type CancelOrder401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CancelOrder404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CancelOrder422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CancelOrderRequest struct {
	PathParams CancelOrderPathParams
	Security   CancelOrderSecurity
}

type CancelOrderResponse struct {
	ContentType                         string
	StatusCode                          int64
	CancelOrder401ApplicationJSONObject *CancelOrder401ApplicationJSON
	CancelOrder404ApplicationJSONObject *CancelOrder404ApplicationJSON
	CancelOrder422ApplicationJSONObject *CancelOrder422ApplicationJSON
	OrderBig                            *shared.OrderBig
}
