package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListOptionValuesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListOptionValues401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListOptionValuesRequest struct {
	Security ListOptionValuesSecurity
}

type ListOptionValuesResponse struct {
	ContentType                              string
	StatusCode                               int64
	ListOptionValues401ApplicationJSONObject *ListOptionValues401ApplicationJSON
	OptionValues                             []shared.OptionValue
}
