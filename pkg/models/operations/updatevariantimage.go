package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateVariantImagePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	VariantID string `pathParam:"style=simple,explode=false,name=variant_id"`
}

type UpdateVariantImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateVariantImageRequest struct {
	PathParams UpdateVariantImagePathParams
	Request    shared.ImageInput `request:"mediaType=application/json"`
	Security   UpdateVariantImageSecurity
}

type UpdateVariantImage422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateVariantImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateVariantImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateVariantImageResponse struct {
	ContentType                                string
	StatusCode                                 int
	Image                                      *shared.Image
	UpdateVariantImage401ApplicationJSONObject *UpdateVariantImage401ApplicationJSON
	UpdateVariantImage404ApplicationJSONObject *UpdateVariantImage404ApplicationJSON
	UpdateVariantImage422ApplicationJSONObject *UpdateVariantImage422ApplicationJSON
}
