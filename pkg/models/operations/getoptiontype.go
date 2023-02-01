package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetOptionTypePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetOptionTypeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetOptionType401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOptionType404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetOptionTypeRequest struct {
	PathParams GetOptionTypePathParams
	Security   GetOptionTypeSecurity
}

type GetOptionTypeResponse struct {
	ContentType                           string
	StatusCode                            int64
	GetOptionType401ApplicationJSONObject *GetOptionType401ApplicationJSON
	GetOptionType404ApplicationJSONObject *GetOptionType404ApplicationJSON
	OptionType                            *shared.OptionType
}
