package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteProductVariantPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type DeleteProductVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteProductVariant401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteProductVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductVariant422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductVariantRequest struct {
	PathParams DeleteProductVariantPathParams
	Security   DeleteProductVariantSecurity
}

type DeleteProductVariantResponse struct {
	ContentType                                  string
	StatusCode                                   int64
	DeleteProductVariant401ApplicationJSONObject *DeleteProductVariant401ApplicationJSON
	DeleteProductVariant404ApplicationJSONObject *DeleteProductVariant404ApplicationJSON
	DeleteProductVariant422ApplicationJSONObject *DeleteProductVariant422ApplicationJSON
}
