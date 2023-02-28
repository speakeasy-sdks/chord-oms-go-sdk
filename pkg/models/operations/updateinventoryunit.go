package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateInventoryUnitPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateInventoryUnitSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateInventoryUnitRequest struct {
	PathParams UpdateInventoryUnitPathParams
	Request    *shared.InventoryUnitInput `request:"mediaType=application/json"`
	Security   UpdateInventoryUnitSecurity
}

type UpdateInventoryUnit422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateInventoryUnit404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateInventoryUnit401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateInventoryUnitResponse struct {
	ContentType                                 string
	StatusCode                                  int
	InventoryUnit                               *shared.InventoryUnit
	UpdateInventoryUnit401ApplicationJSONObject *UpdateInventoryUnit401ApplicationJSON
	UpdateInventoryUnit404ApplicationJSONObject *UpdateInventoryUnit404ApplicationJSON
	UpdateInventoryUnit422ApplicationJSONObject *UpdateInventoryUnit422ApplicationJSON
}
