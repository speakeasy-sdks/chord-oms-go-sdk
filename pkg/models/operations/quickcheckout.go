package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type QuickCheckoutSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type QuickCheckout201ApplicationJSON struct {
	ID          *string `json:"id,omitempty"`
	OrderNumber *string `json:"order_number,omitempty"`
	URL         *string `json:"url,omitempty"`
}

type QuickCheckout401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type QuickCheckoutRequest struct {
	Request  shared.OrderInput `request:"mediaType=application/json"`
	Security QuickCheckoutSecurity
}

type QuickCheckoutResponse struct {
	ContentType                           string
	StatusCode                            int64
	QuickCheckout201ApplicationJSONObject *QuickCheckout201ApplicationJSON
	QuickCheckout401ApplicationJSONObject *QuickCheckout401ApplicationJSON
}
