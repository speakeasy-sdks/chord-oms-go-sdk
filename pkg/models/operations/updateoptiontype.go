package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateOptionTypePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateOptionTypeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateOptionTypeRequest struct {
	PathParams UpdateOptionTypePathParams
	Request    shared.OptionType `request:"mediaType=application/json"`
	Security   UpdateOptionTypeSecurity
}

type UpdateOptionType422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOptionType404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOptionType401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOptionTypeResponse struct {
	ContentType                              string
	StatusCode                               int
	OptionType                               *shared.OptionType
	UpdateOptionType401ApplicationJSONObject *UpdateOptionType401ApplicationJSON
	UpdateOptionType404ApplicationJSONObject *UpdateOptionType404ApplicationJSON
	UpdateOptionType422ApplicationJSONObject *UpdateOptionType422ApplicationJSON
}
