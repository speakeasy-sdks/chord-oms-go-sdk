package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteZonePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteZoneSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteZoneRequest struct {
	PathParams DeleteZonePathParams
	Security   DeleteZoneSecurity
}

type DeleteZone422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteZone404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteZone401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteZoneResponse struct {
	ContentType                        string
	StatusCode                         int
	DeleteZone401ApplicationJSONObject *DeleteZone401ApplicationJSON
	DeleteZone404ApplicationJSONObject *DeleteZone404ApplicationJSON
	DeleteZone422ApplicationJSONObject *DeleteZone422ApplicationJSON
}
