package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetMoneyConfigSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetMoneyConfigRequest struct {
	Security GetMoneyConfigSecurity
}

type GetMoneyConfig401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetMoneyConfig200ApplicationJSON struct {
	Symbol *string `json:"symbol,omitempty"`
}

type GetMoneyConfigResponse struct {
	ContentType                            string
	StatusCode                             int
	GetMoneyConfig200ApplicationJSONObject *GetMoneyConfig200ApplicationJSON
	GetMoneyConfig401ApplicationJSONObject *GetMoneyConfig401ApplicationJSON
}
