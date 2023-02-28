package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateOrderPathParams struct {
	OrderNumber string `pathParam:"style=simple,explode=false,name=order_number"`
}

type UpdateOrderSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	OrderToken      *shared.SchemeOrderToken      `security:"scheme,type=apiKey,subtype=header"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateOrderRequest struct {
	PathParams UpdateOrderPathParams
	Request    map[string]interface{} `request:"mediaType=application/json"`
	Security   UpdateOrderSecurity
}

type UpdateOrder500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type UpdateOrder422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOrder404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrder401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOrderResponse struct {
	ContentType                         string
	OrderBig                            *shared.OrderBig
	StatusCode                          int
	UpdateOrder401ApplicationJSONObject *UpdateOrder401ApplicationJSON
	UpdateOrder404ApplicationJSONObject *UpdateOrder404ApplicationJSON
	UpdateOrder422ApplicationJSONObject *UpdateOrder422ApplicationJSON
	UpdateOrder500ApplicationJSONObject *UpdateOrder500ApplicationJSON
}
