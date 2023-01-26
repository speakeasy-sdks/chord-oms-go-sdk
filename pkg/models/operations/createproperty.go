package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreatePropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateProperty401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateProperty422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreatePropertyRequest struct {
	Request  shared.PropertyInput `request:"mediaType=application/json"`
	Security CreatePropertySecurity
}

type CreatePropertyResponse struct {
	ContentType                            string
	StatusCode                             int64
	CreateProperty401ApplicationJSONObject *CreateProperty401ApplicationJSON
	CreateProperty422ApplicationJSONObject *CreateProperty422ApplicationJSON
	Property                               *shared.Property
}
