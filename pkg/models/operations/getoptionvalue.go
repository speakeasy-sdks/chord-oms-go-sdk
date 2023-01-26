package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetOptionValuePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetOptionValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetOptionValue401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetOptionValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOptionValueRequest struct {
	PathParams GetOptionValuePathParams
	Security   GetOptionValueSecurity
}

type GetOptionValueResponse struct {
	ContentType                            string
	StatusCode                             int64
	GetOptionValue401ApplicationJSONObject *GetOptionValue401ApplicationJSON
	GetOptionValue404ApplicationJSONObject *GetOptionValue404ApplicationJSON
	OptionValue                            *shared.OptionValue
}
