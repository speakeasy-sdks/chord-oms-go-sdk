package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type QuickCheckoutSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type QuickCheckoutRequest struct {
	Request  map[string]interface{} `request:"mediaType=application/json"`
	Security QuickCheckoutSecurity
}

type QuickCheckout401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type QuickCheckout201ApplicationJSON struct {
	ID          *string `json:"id,omitempty"`
	OrderNumber *string `json:"order_number,omitempty"`
	URL         *string `json:"url,omitempty"`
}

type QuickCheckoutResponse struct {
	ContentType                           string
	Headers                               map[string][]string
	StatusCode                            int
	QuickCheckout201ApplicationJSONObject *QuickCheckout201ApplicationJSON
	QuickCheckout401ApplicationJSONObject *QuickCheckout401ApplicationJSON
}
