package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateProductVariantPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type UpdateProductVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateProductVariantRequest struct {
	PathParams UpdateProductVariantPathParams
	Request    *shared.VariantInput `request:"mediaType=application/json"`
	Security   UpdateProductVariantSecurity
}

type UpdateProductVariant422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateProductVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateProductVariant401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateProductVariantResponse struct {
	ContentType                                  string
	StatusCode                                   int
	UpdateProductVariant401ApplicationJSONObject *UpdateProductVariant401ApplicationJSON
	UpdateProductVariant404ApplicationJSONObject *UpdateProductVariant404ApplicationJSON
	UpdateProductVariant422ApplicationJSONObject *UpdateProductVariant422ApplicationJSON
	Variant                                      *shared.Variant
}
