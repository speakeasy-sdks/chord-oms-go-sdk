package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateOptionTypeValuePathParams struct {
	ID           string `pathParam:"style=simple,explode=false,name=id"`
	OptionTypeID string `pathParam:"style=simple,explode=false,name=option_type_id"`
}

type UpdateOptionTypeValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateOptionTypeValueRequest struct {
	PathParams UpdateOptionTypeValuePathParams
	Request    shared.OptionValueInput `request:"mediaType=application/json"`
	Security   UpdateOptionTypeValueSecurity
}

type UpdateOptionTypeValue422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOptionTypeValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOptionTypeValue401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOptionTypeValueResponse struct {
	ContentType                                   string
	StatusCode                                    int
	OptionValue                                   *shared.OptionValue
	UpdateOptionTypeValue401ApplicationJSONObject *UpdateOptionTypeValue401ApplicationJSON
	UpdateOptionTypeValue404ApplicationJSONObject *UpdateOptionTypeValue404ApplicationJSON
	UpdateOptionTypeValue422ApplicationJSONObject *UpdateOptionTypeValue422ApplicationJSON
}
