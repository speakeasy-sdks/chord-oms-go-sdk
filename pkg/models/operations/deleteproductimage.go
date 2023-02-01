package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteProductImagePathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type DeleteProductImageSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteProductImage401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductImage404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductImage422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductImageRequest struct {
	PathParams DeleteProductImagePathParams
	Security   DeleteProductImageSecurity
}

type DeleteProductImageResponse struct {
	ContentType                                string
	StatusCode                                 int64
	DeleteProductImage401ApplicationJSONObject *DeleteProductImage401ApplicationJSON
	DeleteProductImage404ApplicationJSONObject *DeleteProductImage404ApplicationJSON
	DeleteProductImage422ApplicationJSONObject *DeleteProductImage422ApplicationJSON
}
