package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetOptionValuePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetOptionValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetOptionValueRequest struct {
	PathParams GetOptionValuePathParams
	Security   GetOptionValueSecurity
}

type GetOptionValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOptionValue401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOptionValueResponse struct {
	ContentType                            string
	StatusCode                             int
	GetOptionValue401ApplicationJSONObject *GetOptionValue401ApplicationJSON
	GetOptionValue404ApplicationJSONObject *GetOptionValue404ApplicationJSON
	OptionValue                            *shared.OptionValue
}
