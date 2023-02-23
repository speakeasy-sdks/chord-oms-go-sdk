package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetVariantPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetVariantRequest struct {
	PathParams GetVariantPathParams
	Security   GetVariantSecurity
}

type GetVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetVariant401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetVariantResponse struct {
	ContentType                        string
	StatusCode                         int
	GetVariant401ApplicationJSONObject *GetVariant401ApplicationJSON
	GetVariant404ApplicationJSONObject *GetVariant404ApplicationJSON
	Variant                            *shared.Variant
}
