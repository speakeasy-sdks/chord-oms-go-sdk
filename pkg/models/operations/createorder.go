package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateOrderSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOrder401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateOrder422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOrderRequest struct {
	Request  shared.OrderInput `request:"mediaType=application/json"`
	Security CreateOrderSecurity
}

type CreateOrderResponse struct {
	ContentType                         string
	StatusCode                          int64
	CreateOrder401ApplicationJSONObject *CreateOrder401ApplicationJSON
	CreateOrder422ApplicationJSONObject *CreateOrder422ApplicationJSON
	OrderBig                            *shared.OrderBig
}
