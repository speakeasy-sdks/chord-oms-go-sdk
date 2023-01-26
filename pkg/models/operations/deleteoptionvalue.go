package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteOptionValuePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteOptionValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteOptionValue401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteOptionValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionValue422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionValueRequest struct {
	PathParams DeleteOptionValuePathParams
	Security   DeleteOptionValueSecurity
}

type DeleteOptionValueResponse struct {
	ContentType                               string
	StatusCode                                int64
	DeleteOptionValue401ApplicationJSONObject *DeleteOptionValue401ApplicationJSON
	DeleteOptionValue404ApplicationJSONObject *DeleteOptionValue404ApplicationJSON
	DeleteOptionValue422ApplicationJSONObject *DeleteOptionValue422ApplicationJSON
}
