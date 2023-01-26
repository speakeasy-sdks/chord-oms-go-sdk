package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListVariantImagesPathParams struct {
	VariantID string `pathParam:"style=simple,explode=false,name=variant_id"`
}

type ListVariantImagesSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type ListVariantImages200ApplicationJSON struct {
	Images []shared.Image `json:"images,omitempty"`
}

type ListVariantImages401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListVariantImages404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListVariantImagesRequest struct {
	PathParams ListVariantImagesPathParams
	Security   ListVariantImagesSecurity
}

type ListVariantImagesResponse struct {
	ContentType                               string
	StatusCode                                int64
	ListVariantImages200ApplicationJSONObject *ListVariantImages200ApplicationJSON
	ListVariantImages401ApplicationJSONObject *ListVariantImages401ApplicationJSON
	ListVariantImages404ApplicationJSONObject *ListVariantImages404ApplicationJSON
}
