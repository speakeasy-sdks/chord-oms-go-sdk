package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetVariantPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetVariant401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetVariantRequest struct {
	PathParams GetVariantPathParams
	Security   GetVariantSecurity
}

type GetVariantResponse struct {
	ContentType                        string
	StatusCode                         int64
	GetVariant401ApplicationJSONObject *GetVariant401ApplicationJSON
	GetVariant404ApplicationJSONObject *GetVariant404ApplicationJSON
	Variant                            *shared.Variant
}
