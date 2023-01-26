package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetPropertyPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetPropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetProperty401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetPropertyRequest struct {
	PathParams GetPropertyPathParams
	Security   GetPropertySecurity
}

type GetPropertyResponse struct {
	ContentType                         string
	StatusCode                          int64
	GetProperty401ApplicationJSONObject *GetProperty401ApplicationJSON
	GetProperty404ApplicationJSONObject *GetProperty404ApplicationJSON
	Property                            *shared.Property
}
