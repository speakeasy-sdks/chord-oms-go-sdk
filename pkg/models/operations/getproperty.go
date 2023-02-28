package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetPropertyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetPropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetPropertyRequest struct {
	PathParams GetPropertyPathParams
	Security   GetPropertySecurity
}

type GetProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProperty401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetPropertyResponse struct {
	ContentType                         string
	StatusCode                          int
	GetProperty401ApplicationJSONObject *GetProperty401ApplicationJSON
	GetProperty404ApplicationJSONObject *GetProperty404ApplicationJSON
	Property                            *shared.Property
}
