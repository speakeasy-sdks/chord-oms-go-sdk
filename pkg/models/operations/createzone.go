package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateZoneSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateZone401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateZone422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateZoneRequest struct {
	Request  shared.ZoneInput `request:"mediaType=application/json"`
	Security CreateZoneSecurity
}

type CreateZoneResponse struct {
	ContentType                        string
	StatusCode                         int64
	CreateZone401ApplicationJSONObject *CreateZone401ApplicationJSON
	CreateZone422ApplicationJSONObject *CreateZone422ApplicationJSON
	Zone                               *shared.Zone
}
