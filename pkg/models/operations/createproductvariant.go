package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateProductVariantPathParams struct {
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type CreateProductVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateProductVariantRequest struct {
	PathParams CreateProductVariantPathParams
	Request    shared.VariantInput `request:"mediaType=application/json"`
	Security   CreateProductVariantSecurity
}

type CreateProductVariant422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateProductVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateProductVariant401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateProductVariantResponse struct {
	ContentType                                  string
	StatusCode                                   int
	CreateProductVariant401ApplicationJSONObject *CreateProductVariant401ApplicationJSON
	CreateProductVariant404ApplicationJSONObject *CreateProductVariant404ApplicationJSON
	CreateProductVariant422ApplicationJSONObject *CreateProductVariant422ApplicationJSON
	Variant                                      *shared.Variant
}
