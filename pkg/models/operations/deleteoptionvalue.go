package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteOptionValuePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteOptionValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteOptionValueRequest struct {
	PathParams DeleteOptionValuePathParams
	Security   DeleteOptionValueSecurity
}

type DeleteOptionValue422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionValue401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionValueResponse struct {
	ContentType                               string
	StatusCode                                int
	DeleteOptionValue401ApplicationJSONObject *DeleteOptionValue401ApplicationJSON
	DeleteOptionValue404ApplicationJSONObject *DeleteOptionValue404ApplicationJSON
	DeleteOptionValue422ApplicationJSONObject *DeleteOptionValue422ApplicationJSON
}
