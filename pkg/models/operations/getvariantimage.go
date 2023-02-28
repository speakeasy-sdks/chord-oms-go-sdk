package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetVariantImagePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	VariantID string `pathParam:"style=simple,explode=false,name=variant_id"`
}

type GetVariantImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetVariantImageRequest struct {
	PathParams GetVariantImagePathParams
	Security   GetVariantImageSecurity
}

type GetVariantImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetVariantImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetVariantImageResponse struct {
	ContentType                             string
	StatusCode                              int
	GetVariantImage401ApplicationJSONObject *GetVariantImage401ApplicationJSON
	GetVariantImage404ApplicationJSONObject *GetVariantImage404ApplicationJSON
	Image                                   *shared.Image
}
