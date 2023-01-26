package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateOptionTypePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateOptionTypeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateOptionType401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateOptionType404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOptionType422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOptionTypeRequest struct {
	PathParams UpdateOptionTypePathParams
	Request    shared.OptionType `request:"mediaType=application/json"`
	Security   UpdateOptionTypeSecurity
}

type UpdateOptionTypeResponse struct {
	ContentType                              string
	StatusCode                               int64
	OptionType                               *shared.OptionType
	UpdateOptionType401ApplicationJSONObject *UpdateOptionType401ApplicationJSON
	UpdateOptionType404ApplicationJSONObject *UpdateOptionType404ApplicationJSON
	UpdateOptionType422ApplicationJSONObject *UpdateOptionType422ApplicationJSON
}
