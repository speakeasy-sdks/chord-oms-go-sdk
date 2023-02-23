package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteProductPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteProductSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteProductRequest struct {
	PathParams DeleteProductPathParams
	Security   DeleteProductSecurity
}

type DeleteProduct422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProduct404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProduct401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductResponse struct {
	ContentType                           string
	StatusCode                            int
	DeleteProduct401ApplicationJSONObject *DeleteProduct401ApplicationJSON
	DeleteProduct404ApplicationJSONObject *DeleteProduct404ApplicationJSON
	DeleteProduct422ApplicationJSONObject *DeleteProduct422ApplicationJSON
}
