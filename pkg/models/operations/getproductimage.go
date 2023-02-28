package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetProductImagePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type GetProductImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetProductImageRequest struct {
	PathParams GetProductImagePathParams
	Security   GetProductImageSecurity
}

type GetProductImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductImageResponse struct {
	ContentType                             string
	StatusCode                              int
	GetProductImage401ApplicationJSONObject *GetProductImage401ApplicationJSON
	GetProductImage404ApplicationJSONObject *GetProductImage404ApplicationJSON
	Image                                   *shared.Image
}
