package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetProductVariantPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type GetProductVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetProductVariant401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetProductVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductVariantRequest struct {
	PathParams GetProductVariantPathParams
	Security   GetProductVariantSecurity
}

type GetProductVariantResponse struct {
	ContentType                               string
	StatusCode                                int64
	GetProductVariant401ApplicationJSONObject *GetProductVariant401ApplicationJSON
	GetProductVariant404ApplicationJSONObject *GetProductVariant404ApplicationJSON
	Variant                                   *shared.Variant
}
