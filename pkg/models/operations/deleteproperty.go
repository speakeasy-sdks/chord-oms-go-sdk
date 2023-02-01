package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeletePropertyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeletePropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteProperty401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProperty422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeletePropertyRequest struct {
	PathParams DeletePropertyPathParams
	Security   DeletePropertySecurity
}

type DeletePropertyResponse struct {
	ContentType                            string
	StatusCode                             int64
	DeleteProperty401ApplicationJSONObject *DeleteProperty401ApplicationJSON
	DeleteProperty404ApplicationJSONObject *DeleteProperty404ApplicationJSON
	DeleteProperty422ApplicationJSONObject *DeleteProperty422ApplicationJSON
}
