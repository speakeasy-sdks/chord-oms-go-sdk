package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetInventoryUnitPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetInventoryUnitSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetInventoryUnitRequest struct {
	PathParams GetInventoryUnitPathParams
	Security   GetInventoryUnitSecurity
}

type GetInventoryUnit404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetInventoryUnit401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetInventoryUnitResponse struct {
	ContentType                              string
	StatusCode                               int
	GetInventoryUnit401ApplicationJSONObject *GetInventoryUnit401ApplicationJSON
	GetInventoryUnit404ApplicationJSONObject *GetInventoryUnit404ApplicationJSON
	InventoryUnit                            *shared.InventoryUnit
}
