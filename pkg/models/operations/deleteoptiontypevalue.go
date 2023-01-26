package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteOptionTypeValuePathParams struct {
	ID           string `pathParam:"style=simple,explode=false,name=id"`
	OptionTypeID string `pathParam:"style=simple,explode=false,name=option_type_id"`
}

type DeleteOptionTypeValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteOptionTypeValue401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteOptionTypeValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionTypeValue422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionTypeValueRequest struct {
	PathParams DeleteOptionTypeValuePathParams
	Security   DeleteOptionTypeValueSecurity
}

type DeleteOptionTypeValueResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	DeleteOptionTypeValue401ApplicationJSONObject *DeleteOptionTypeValue401ApplicationJSON
	DeleteOptionTypeValue404ApplicationJSONObject *DeleteOptionTypeValue404ApplicationJSON
	DeleteOptionTypeValue422ApplicationJSONObject *DeleteOptionTypeValue422ApplicationJSON
}
