package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListOptionTypesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListOptionTypes401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListOptionTypesRequest struct {
	Security ListOptionTypesSecurity
}

type ListOptionTypesResponse struct {
	ContentType                             string
	StatusCode                              int64
	ListOptionTypes401ApplicationJSONObject *ListOptionTypes401ApplicationJSON
	OptionTypes                             []shared.OptionType
}
