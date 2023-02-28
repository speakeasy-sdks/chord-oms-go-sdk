package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeletePropertyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeletePropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeletePropertyRequest struct {
	PathParams DeletePropertyPathParams
	Security   DeletePropertySecurity
}

type DeleteProperty422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProperty401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeletePropertyResponse struct {
	ContentType                            string
	StatusCode                             int
	DeleteProperty401ApplicationJSONObject *DeleteProperty401ApplicationJSON
	DeleteProperty404ApplicationJSONObject *DeleteProperty404ApplicationJSON
	DeleteProperty422ApplicationJSONObject *DeleteProperty422ApplicationJSON
}
