package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateOptionTypeValuePathParams struct {
	OptionTypeID string `pathParam:"style=simple,explode=false,name=option_type_id"`
}

type CreateOptionTypeValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateOptionTypeValueRequest struct {
	PathParams CreateOptionTypeValuePathParams
	Request    shared.OptionValueInput `request:"mediaType=application/json"`
	Security   CreateOptionTypeValueSecurity
}

type CreateOptionTypeValue422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOptionTypeValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOptionTypeValue401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOptionTypeValueResponse struct {
	ContentType                                   string
	StatusCode                                    int
	CreateOptionTypeValue401ApplicationJSONObject *CreateOptionTypeValue401ApplicationJSON
	CreateOptionTypeValue404ApplicationJSONObject *CreateOptionTypeValue404ApplicationJSON
	CreateOptionTypeValue422ApplicationJSONObject *CreateOptionTypeValue422ApplicationJSON
	OptionValue                                   *shared.OptionValue
}
