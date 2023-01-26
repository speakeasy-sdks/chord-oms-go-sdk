package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateZonePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateZoneSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateZone401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateZone404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateZone422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateZoneRequest struct {
	PathParams UpdateZonePathParams
	Request    shared.ZoneInput `request:"mediaType=application/json"`
	Security   UpdateZoneSecurity
}

type UpdateZoneResponse struct {
	ContentType                        string
	StatusCode                         int64
	UpdateZone401ApplicationJSONObject *UpdateZone401ApplicationJSON
	UpdateZone404ApplicationJSONObject *UpdateZone404ApplicationJSON
	UpdateZone422ApplicationJSONObject *UpdateZone422ApplicationJSON
	Zone                               *shared.Zone
}
