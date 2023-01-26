package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteZonePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteZoneSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteZone401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteZone404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteZone422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteZoneRequest struct {
	PathParams DeleteZonePathParams
	Security   DeleteZoneSecurity
}

type DeleteZoneResponse struct {
	ContentType                        string
	StatusCode                         int64
	DeleteZone401ApplicationJSONObject *DeleteZone401ApplicationJSON
	DeleteZone404ApplicationJSONObject *DeleteZone404ApplicationJSON
	DeleteZone422ApplicationJSONObject *DeleteZone422ApplicationJSON
}
