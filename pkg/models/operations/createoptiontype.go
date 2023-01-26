package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateOptionTypeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateOptionType401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateOptionType422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOptionTypeRequest struct {
	Request  shared.OptionTypeInput `request:"mediaType=application/json"`
	Security CreateOptionTypeSecurity
}

type CreateOptionTypeResponse struct {
	ContentType                              string
	StatusCode                               int64
	CreateOptionType401ApplicationJSONObject *CreateOptionType401ApplicationJSON
	CreateOptionType422ApplicationJSONObject *CreateOptionType422ApplicationJSON
	OptionType                               *shared.OptionType
}
