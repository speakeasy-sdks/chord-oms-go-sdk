package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateOptionTypeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateOptionTypeRequest struct {
	Request  shared.OptionTypeInput `request:"mediaType=application/json"`
	Security CreateOptionTypeSecurity
}

type CreateOptionType422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateOptionType401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateOptionTypeResponse struct {
	ContentType                              string
	StatusCode                               int
	CreateOptionType401ApplicationJSONObject *CreateOptionType401ApplicationJSON
	CreateOptionType422ApplicationJSONObject *CreateOptionType422ApplicationJSON
	OptionType                               *shared.OptionType
}
