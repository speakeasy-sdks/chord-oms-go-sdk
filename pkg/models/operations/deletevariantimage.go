package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteVariantImagePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	VariantID string `pathParam:"style=simple,explode=false,name=variant_id"`
}

type DeleteVariantImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteVariantImageRequest struct {
	PathParams DeleteVariantImagePathParams
	Security   DeleteVariantImageSecurity
}

type DeleteVariantImage422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteVariantImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteVariantImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteVariantImageResponse struct {
	ContentType                                string
	StatusCode                                 int
	DeleteVariantImage401ApplicationJSONObject *DeleteVariantImage401ApplicationJSON
	DeleteVariantImage404ApplicationJSONObject *DeleteVariantImage404ApplicationJSON
	DeleteVariantImage422ApplicationJSONObject *DeleteVariantImage422ApplicationJSON
}
