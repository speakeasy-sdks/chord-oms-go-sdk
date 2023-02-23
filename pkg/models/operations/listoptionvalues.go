package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListOptionValuesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListOptionValuesRequest struct {
	Security ListOptionValuesSecurity
}

type ListOptionValues401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListOptionValuesResponse struct {
	ContentType                              string
	StatusCode                               int
	ListOptionValues401ApplicationJSONObject *ListOptionValues401ApplicationJSON
	OptionValues                             []shared.OptionValue
}
