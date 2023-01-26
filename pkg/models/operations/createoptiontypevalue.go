package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateOptionTypeValuePathParams struct {
	OptionTypeID string `pathParam:"style=simple,explode=false,name=option_type_id"`
}

type CreateOptionTypeValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOptionTypeValue401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateOptionTypeValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOptionTypeValue422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOptionTypeValueRequest struct {
	PathParams CreateOptionTypeValuePathParams
	Request    shared.OptionValueInput `request:"mediaType=application/json"`
	Security   CreateOptionTypeValueSecurity
}

type CreateOptionTypeValueResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	CreateOptionTypeValue401ApplicationJSONObject *CreateOptionTypeValue401ApplicationJSON
	CreateOptionTypeValue404ApplicationJSONObject *CreateOptionTypeValue404ApplicationJSON
	CreateOptionTypeValue422ApplicationJSONObject *CreateOptionTypeValue422ApplicationJSON
	OptionValue                                   *shared.OptionValue
}
