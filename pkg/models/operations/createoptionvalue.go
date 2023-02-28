package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateOptionValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateOptionValueRequest struct {
	Request  shared.OptionValueInput `request:"mediaType=application/json"`
	Security CreateOptionValueSecurity
}

type CreateOptionValue422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOptionValue401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOptionValueResponse struct {
	ContentType                               string
	StatusCode                                int
	CreateOptionValue401ApplicationJSONObject *CreateOptionValue401ApplicationJSON
	CreateOptionValue422ApplicationJSONObject *CreateOptionValue422ApplicationJSON
	OptionValue                               *shared.OptionValue
}
