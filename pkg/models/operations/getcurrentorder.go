package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetCurrentOrderSecurity struct {
	APIKey     *shared.SchemeAPIKey     `security:"scheme,type=apiKey,subtype=header"`
	OrderToken *shared.SchemeOrderToken `security:"scheme,type=apiKey,subtype=header"`
}

type GetCurrentOrder401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetCurrentOrderRequest struct {
	Security GetCurrentOrderSecurity
}

type GetCurrentOrderResponse struct {
	ContentType                             string
	StatusCode                              int64
	GetCurrentOrder401ApplicationJSONObject *GetCurrentOrder401ApplicationJSON
	OrderBig                                *shared.OrderBig
}
