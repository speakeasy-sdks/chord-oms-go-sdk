package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListProductImagesPathParams struct {
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type ListProductImagesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListProductImages200ApplicationJSON struct {
	Images []shared.Image `json:"images,omitempty"`
}

type ListProductImages401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListProductImages404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListProductImagesRequest struct {
	PathParams ListProductImagesPathParams
	Security   ListProductImagesSecurity
}

type ListProductImagesResponse struct {
	ContentType                               string
	StatusCode                                int64
	ListProductImages200ApplicationJSONObject *ListProductImages200ApplicationJSON
	ListProductImages401ApplicationJSONObject *ListProductImages401ApplicationJSON
	ListProductImages404ApplicationJSONObject *ListProductImages404ApplicationJSON
}
