package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListOptionTypesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListOptionTypesRequest struct {
	Security ListOptionTypesSecurity
}

type ListOptionTypes401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListOptionTypesResponse struct {
	ContentType                             string
	StatusCode                              int
	ListOptionTypes401ApplicationJSONObject *ListOptionTypes401ApplicationJSON
	OptionTypes                             []shared.OptionType
}
