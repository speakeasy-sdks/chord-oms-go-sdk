package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateOptionValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOptionValue401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOptionValue422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOptionValueRequest struct {
	Request  shared.OptionValueInput `request:"mediaType=application/json"`
	Security CreateOptionValueSecurity
}

type CreateOptionValueResponse struct {
	ContentType                               string
	StatusCode                                int64
	CreateOptionValue401ApplicationJSONObject *CreateOptionValue401ApplicationJSON
	CreateOptionValue422ApplicationJSONObject *CreateOptionValue422ApplicationJSON
	OptionValue                               *shared.OptionValue
}
