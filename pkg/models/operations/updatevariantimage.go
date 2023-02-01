package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateVariantImagePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	VariantID string `pathParam:"style=simple,explode=false,name=variant_id"`
}

type UpdateVariantImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateVariantImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateVariantImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateVariantImage422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateVariantImageRequest struct {
	PathParams UpdateVariantImagePathParams
	Request    shared.ImageInput `request:"mediaType=application/json"`
	Security   UpdateVariantImageSecurity
}

type UpdateVariantImageResponse struct {
	ContentType                                string
	StatusCode                                 int64
	Image                                      *shared.Image
	UpdateVariantImage401ApplicationJSONObject *UpdateVariantImage401ApplicationJSON
	UpdateVariantImage404ApplicationJSONObject *UpdateVariantImage404ApplicationJSON
	UpdateVariantImage422ApplicationJSONObject *UpdateVariantImage422ApplicationJSON
}
