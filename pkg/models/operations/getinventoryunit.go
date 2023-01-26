package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetInventoryUnitPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetInventoryUnitSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetInventoryUnit401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetInventoryUnit404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetInventoryUnitRequest struct {
	PathParams GetInventoryUnitPathParams
	Security   GetInventoryUnitSecurity
}

type GetInventoryUnitResponse struct {
	ContentType                              string
	StatusCode                               int64
	GetInventoryUnit401ApplicationJSONObject *GetInventoryUnit401ApplicationJSON
	GetInventoryUnit404ApplicationJSONObject *GetInventoryUnit404ApplicationJSON
	InventoryUnit                            *shared.InventoryUnit
}
