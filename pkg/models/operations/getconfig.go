package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetConfigSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetConfigRequest struct {
	Security GetConfigSecurity
}

type GetConfig401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetConfig200ApplicationJSON struct {
	DefaultCountryID  *int64  `json:"default_country_id,omitempty"`
	DefaultCountryIso *string `json:"default_country_iso,omitempty"`
}

type GetConfigResponse struct {
	ContentType                       string
	StatusCode                        int
	GetConfig200ApplicationJSONObject *GetConfig200ApplicationJSON
	GetConfig401ApplicationJSONObject *GetConfig401ApplicationJSON
}
