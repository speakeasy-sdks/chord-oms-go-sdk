package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreatePropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreatePropertyRequest struct {
	Request  shared.PropertyInput `request:"mediaType=application/json"`
	Security CreatePropertySecurity
}

type CreateProperty422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateProperty401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreatePropertyResponse struct {
	ContentType                            string
	StatusCode                             int
	CreateProperty401ApplicationJSONObject *CreateProperty401ApplicationJSON
	CreateProperty422ApplicationJSONObject *CreateProperty422ApplicationJSON
	Property                               *shared.Property
}
