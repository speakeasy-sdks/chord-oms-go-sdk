package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteOptionTypeValuePathParams struct {
	ID           string `pathParam:"style=simple,explode=false,name=id"`
	OptionTypeID string `pathParam:"style=simple,explode=false,name=option_type_id"`
}

type DeleteOptionTypeValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteOptionTypeValueRequest struct {
	PathParams DeleteOptionTypeValuePathParams
	Security   DeleteOptionTypeValueSecurity
}

type DeleteOptionTypeValue422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionTypeValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionTypeValue401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionTypeValueResponse struct {
	ContentType                                   string
	StatusCode                                    int
	DeleteOptionTypeValue401ApplicationJSONObject *DeleteOptionTypeValue401ApplicationJSON
	DeleteOptionTypeValue404ApplicationJSONObject *DeleteOptionTypeValue404ApplicationJSON
	DeleteOptionTypeValue422ApplicationJSONObject *DeleteOptionTypeValue422ApplicationJSON
}
