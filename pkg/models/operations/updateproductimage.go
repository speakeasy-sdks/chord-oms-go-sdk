package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateProductImagePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type UpdateProductImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateProductImage401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateProductImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateProductImage422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateProductImageRequest struct {
	PathParams UpdateProductImagePathParams
	Request    shared.ImageInput `request:"mediaType=application/json"`
	Security   UpdateProductImageSecurity
}

type UpdateProductImageResponse struct {
	ContentType                                string
	StatusCode                                 int64
	Image                                      *shared.Image
	UpdateProductImage401ApplicationJSONObject *UpdateProductImage401ApplicationJSON
	UpdateProductImage404ApplicationJSONObject *UpdateProductImage404ApplicationJSON
	UpdateProductImage422ApplicationJSONObject *UpdateProductImage422ApplicationJSON
}
