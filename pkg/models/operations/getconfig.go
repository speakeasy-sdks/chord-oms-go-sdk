package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetConfigSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetConfig200ApplicationJSON struct {
	DefaultCountryID  *int64  `json:"default_country_id,omitempty"`
	DefaultCountryIso *string `json:"default_country_iso,omitempty"`
}

type GetConfig401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetConfigRequest struct {
	Security GetConfigSecurity
}

type GetConfigResponse struct {
	ContentType                       string
	StatusCode                        int64
	GetConfig200ApplicationJSONObject *GetConfig200ApplicationJSON
	GetConfig401ApplicationJSONObject *GetConfig401ApplicationJSON
}
