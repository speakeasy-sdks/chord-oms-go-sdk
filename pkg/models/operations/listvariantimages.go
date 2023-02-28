package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListVariantImagesPathParams struct {
	VariantID string `pathParam:"style=simple,explode=false,name=variant_id"`
}

type ListVariantImagesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type ListVariantImagesRequest struct {
	PathParams ListVariantImagesPathParams
	Security   ListVariantImagesSecurity
}

type ListVariantImages404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListVariantImages401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListVariantImages200ApplicationJSON struct {
	Images []shared.Image `json:"images,omitempty"`
}

type ListVariantImagesResponse struct {
	ContentType                               string
	StatusCode                                int
	ListVariantImages200ApplicationJSONObject *ListVariantImages200ApplicationJSON
	ListVariantImages401ApplicationJSONObject *ListVariantImages401ApplicationJSON
	ListVariantImages404ApplicationJSONObject *ListVariantImages404ApplicationJSON
}
