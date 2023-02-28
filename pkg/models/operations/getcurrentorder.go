package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetCurrentOrderSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	OrderToken      *shared.SchemeOrderToken      `security:"scheme,type=apiKey,subtype=header"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type GetCurrentOrderRequest struct {
	Security GetCurrentOrderSecurity
}

type GetCurrentOrder500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type GetCurrentOrder401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCurrentOrderResponse struct {
	ContentType                             string
	OrderBig                                *shared.OrderBig
	StatusCode                              int
	GetCurrentOrder401ApplicationJSONObject *GetCurrentOrder401ApplicationJSON
	GetCurrentOrder500ApplicationJSONObject *GetCurrentOrder500ApplicationJSON
}
