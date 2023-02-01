package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateOrderPathParams struct {
	Number string `pathParam:"style=simple,explode=false,name=number"`
}

type UpdateOrderSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateOrder401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrder404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrder422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOrderRequest struct {
	PathParams UpdateOrderPathParams
	Request    shared.OrderInput `request:"mediaType=application/json"`
	Security   UpdateOrderSecurity
}

type UpdateOrderResponse struct {
	ContentType                         string
	StatusCode                          int64
	OrderBig                            *shared.OrderBig
	UpdateOrder401ApplicationJSONObject *UpdateOrder401ApplicationJSON
	UpdateOrder404ApplicationJSONObject *UpdateOrder404ApplicationJSON
	UpdateOrder422ApplicationJSONObject *UpdateOrder422ApplicationJSON
}
