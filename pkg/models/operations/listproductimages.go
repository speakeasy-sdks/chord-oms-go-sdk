package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListProductImagesPathParams struct {
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type ListProductImagesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListProductImagesRequest struct {
	PathParams ListProductImagesPathParams
	Security   ListProductImagesSecurity
}

type ListProductImages404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListProductImages401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListProductImages200ApplicationJSON struct {
	Images []shared.Image `json:"images,omitempty"`
}

type ListProductImagesResponse struct {
	ContentType                               string
	StatusCode                                int
	ListProductImages200ApplicationJSONObject *ListProductImages200ApplicationJSON
	ListProductImages401ApplicationJSONObject *ListProductImages401ApplicationJSON
	ListProductImages404ApplicationJSONObject *ListProductImages404ApplicationJSON
}
