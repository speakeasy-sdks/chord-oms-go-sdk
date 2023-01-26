package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetMoneyConfigSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetMoneyConfig200ApplicationJSON struct {
	Symbol *string `json:"symbol,omitempty"`
}

type GetMoneyConfig401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetMoneyConfigRequest struct {
	Security GetMoneyConfigSecurity
}

type GetMoneyConfigResponse struct {
	ContentType                            string
	StatusCode                             int64
	GetMoneyConfig200ApplicationJSONObject *GetMoneyConfig200ApplicationJSON
	GetMoneyConfig401ApplicationJSONObject *GetMoneyConfig401ApplicationJSON
}
