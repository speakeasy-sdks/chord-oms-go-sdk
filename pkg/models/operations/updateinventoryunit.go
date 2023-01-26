package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateInventoryUnitPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateInventoryUnitSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateInventoryUnit401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateInventoryUnit404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateInventoryUnit422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateInventoryUnitRequest struct {
	PathParams UpdateInventoryUnitPathParams
	Request    *shared.InventoryUnitInput `request:"mediaType=application/json"`
	Security   UpdateInventoryUnitSecurity
}

type UpdateInventoryUnitResponse struct {
	ContentType                                 string
	StatusCode                                  int64
	InventoryUnit                               *shared.InventoryUnit
	UpdateInventoryUnit401ApplicationJSONObject *UpdateInventoryUnit401ApplicationJSON
	UpdateInventoryUnit404ApplicationJSONObject *UpdateInventoryUnit404ApplicationJSON
	UpdateInventoryUnit422ApplicationJSONObject *UpdateInventoryUnit422ApplicationJSON
}
