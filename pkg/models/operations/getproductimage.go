package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetProductImagePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type GetProductImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetProductImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductImageRequest struct {
	PathParams GetProductImagePathParams
	Security   GetProductImageSecurity
}

type GetProductImageResponse struct {
	ContentType                             string
	StatusCode                              int64
	GetProductImage401ApplicationJSONObject *GetProductImage401ApplicationJSON
	GetProductImage404ApplicationJSONObject *GetProductImage404ApplicationJSON
	Image                                   *shared.Image
}
