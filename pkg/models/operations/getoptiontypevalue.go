package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetOptionTypeValuePathParams struct {
	ID           string `pathParam:"style=simple,explode=false,name=id"`
	OptionTypeID string `pathParam:"style=simple,explode=false,name=option_type_id"`
}

type GetOptionTypeValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetOptionTypeValueRequest struct {
	PathParams GetOptionTypeValuePathParams
	Security   GetOptionTypeValueSecurity
}

type GetOptionTypeValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOptionTypeValue401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOptionTypeValueResponse struct {
	ContentType                                string
	StatusCode                                 int
	GetOptionTypeValue401ApplicationJSONObject *GetOptionTypeValue401ApplicationJSON
	GetOptionTypeValue404ApplicationJSONObject *GetOptionTypeValue404ApplicationJSON
	OptionValue                                *shared.OptionValue
}
