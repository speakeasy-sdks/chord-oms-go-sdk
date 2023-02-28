package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetZonePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetZoneSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetZoneRequest struct {
	PathParams GetZonePathParams
	Security   GetZoneSecurity
}

type GetZone404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetZone401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetZoneResponse struct {
	ContentType                     string
	StatusCode                      int
	GetZone401ApplicationJSONObject *GetZone401ApplicationJSON
	GetZone404ApplicationJSONObject *GetZone404ApplicationJSON
	Zone                            *shared.Zone
}
