package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetProductVariantPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type GetProductVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetProductVariantRequest struct {
	PathParams GetProductVariantPathParams
	Security   GetProductVariantSecurity
}

type GetProductVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductVariant401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductVariantResponse struct {
	ContentType                               string
	StatusCode                                int
	GetProductVariant401ApplicationJSONObject *GetProductVariant401ApplicationJSON
	GetProductVariant404ApplicationJSONObject *GetProductVariant404ApplicationJSON
	Variant                                   *shared.Variant
}
