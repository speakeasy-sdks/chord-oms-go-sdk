package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetZonePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetZoneSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetZone401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetZone404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetZoneRequest struct {
	PathParams GetZonePathParams
	Security   GetZoneSecurity
}

type GetZoneResponse struct {
	ContentType                     string
	StatusCode                      int64
	GetZone401ApplicationJSONObject *GetZone401ApplicationJSON
	GetZone404ApplicationJSONObject *GetZone404ApplicationJSON
	Zone                            *shared.Zone
}
