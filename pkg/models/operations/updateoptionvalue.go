package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateOptionValuePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateOptionValueSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateOptionValue401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOptionValue404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateOptionValue422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateOptionValueRequest struct {
	PathParams UpdateOptionValuePathParams
	Request    shared.OptionValueInput `request:"mediaType=application/json"`
	Security   UpdateOptionValueSecurity
}

type UpdateOptionValueResponse struct {
	ContentType                               string
	StatusCode                                int64
	OptionValue                               *shared.OptionValue
	UpdateOptionValue401ApplicationJSONObject *UpdateOptionValue401ApplicationJSON
	UpdateOptionValue404ApplicationJSONObject *UpdateOptionValue404ApplicationJSON
	UpdateOptionValue422ApplicationJSONObject *UpdateOptionValue422ApplicationJSON
}
