package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateProductImagePathParams struct {
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type CreateProductImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateProductImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateProductImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateProductImage422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateProductImageRequest struct {
	PathParams CreateProductImagePathParams
	Request    shared.ImageInput `request:"mediaType=application/json"`
	Security   CreateProductImageSecurity
}

type CreateProductImageResponse struct {
	ContentType                                string
	StatusCode                                 int64
	CreateProductImage401ApplicationJSONObject *CreateProductImage401ApplicationJSON
	CreateProductImage404ApplicationJSONObject *CreateProductImage404ApplicationJSON
	CreateProductImage422ApplicationJSONObject *CreateProductImage422ApplicationJSON
	Image                                      *shared.Image
}
