package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetVariantImagePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	VariantID string `pathParam:"style=simple,explode=false,name=variant_id"`
}

type GetVariantImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetVariantImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetVariantImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetVariantImageRequest struct {
	PathParams GetVariantImagePathParams
	Security   GetVariantImageSecurity
}

type GetVariantImageResponse struct {
	ContentType                             string
	StatusCode                              int64
	GetVariantImage401ApplicationJSONObject *GetVariantImage401ApplicationJSON
	GetVariantImage404ApplicationJSONObject *GetVariantImage404ApplicationJSON
	Image                                   *shared.Image
}
