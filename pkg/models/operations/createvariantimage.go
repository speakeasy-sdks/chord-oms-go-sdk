package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateVariantImagePathParams struct {
	VariantID string `pathParam:"style=simple,explode=false,name=variant_id"`
}

type CreateVariantImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateVariantImageRequest struct {
	PathParams CreateVariantImagePathParams
	Request    shared.ImageInput `request:"mediaType=application/json"`
	Security   CreateVariantImageSecurity
}

type CreateVariantImage422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateVariantImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateVariantImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateVariantImageResponse struct {
	ContentType                                string
	StatusCode                                 int
	CreateVariantImage401ApplicationJSONObject *CreateVariantImage401ApplicationJSON
	CreateVariantImage404ApplicationJSONObject *CreateVariantImage404ApplicationJSON
	CreateVariantImage422ApplicationJSONObject *CreateVariantImage422ApplicationJSON
	Image                                      *shared.Image
}
