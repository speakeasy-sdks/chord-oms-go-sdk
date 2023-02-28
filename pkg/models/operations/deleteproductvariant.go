package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteProductVariantPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type DeleteProductVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteProductVariantRequest struct {
	PathParams DeleteProductVariantPathParams
	Security   DeleteProductVariantSecurity
}

type DeleteProductVariant422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductVariant401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductVariantResponse struct {
	ContentType                                  string
	StatusCode                                   int
	DeleteProductVariant401ApplicationJSONObject *DeleteProductVariant401ApplicationJSON
	DeleteProductVariant404ApplicationJSONObject *DeleteProductVariant404ApplicationJSON
	DeleteProductVariant422ApplicationJSONObject *DeleteProductVariant422ApplicationJSON
}
