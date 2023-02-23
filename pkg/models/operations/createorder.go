package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateOrderSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOrderRequest struct {
	Request  map[string]interface{} `request:"mediaType=application/json"`
	Security CreateOrderSecurity
}

type CreateOrder500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type CreateOrder422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOrder401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOrderResponse struct {
	ContentType                         string
	OrderBig                            *shared.OrderBig
	StatusCode                          int
	CreateOrder401ApplicationJSONObject *CreateOrder401ApplicationJSON
	CreateOrder422ApplicationJSONObject *CreateOrder422ApplicationJSON
	CreateOrder500ApplicationJSONObject *CreateOrder500ApplicationJSON
}
