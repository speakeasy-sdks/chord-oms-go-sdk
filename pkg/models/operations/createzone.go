package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateZoneSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateZoneRequest struct {
	Request  shared.ZoneInput `request:"mediaType=application/json"`
	Security CreateZoneSecurity
}

type CreateZone422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateZone401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateZoneResponse struct {
	ContentType                        string
	StatusCode                         int
	CreateZone401ApplicationJSONObject *CreateZone401ApplicationJSON
	CreateZone422ApplicationJSONObject *CreateZone422ApplicationJSON
	Zone                               *shared.Zone
}
