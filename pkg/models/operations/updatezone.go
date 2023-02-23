package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateZonePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateZoneSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateZoneRequest struct {
	PathParams UpdateZonePathParams
	Request    shared.ZoneInput `request:"mediaType=application/json"`
	Security   UpdateZoneSecurity
}

type UpdateZone422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateZone404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateZone401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateZoneResponse struct {
	ContentType                        string
	StatusCode                         int
	UpdateZone401ApplicationJSONObject *UpdateZone401ApplicationJSON
	UpdateZone404ApplicationJSONObject *UpdateZone404ApplicationJSON
	UpdateZone422ApplicationJSONObject *UpdateZone422ApplicationJSON
	Zone                               *shared.Zone
}
