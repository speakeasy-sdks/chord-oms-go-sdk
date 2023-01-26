package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteOptionTypePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteOptionTypeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteOptionType401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteOptionType404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteOptionType422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type DeleteOptionTypeRequest struct {
	PathParams DeleteOptionTypePathParams
	Security   DeleteOptionTypeSecurity
}

type DeleteOptionTypeResponse struct {
	ContentType                              string
	StatusCode                               int64
	DeleteOptionType401ApplicationJSONObject *DeleteOptionType401ApplicationJSON
	DeleteOptionType404ApplicationJSONObject *DeleteOptionType404ApplicationJSON
	DeleteOptionType422ApplicationJSONObject *DeleteOptionType422ApplicationJSON
	OptionType                               *shared.OptionType
}
