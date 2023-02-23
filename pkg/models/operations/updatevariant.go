package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateVariantPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateVariantRequest struct {
	PathParams UpdateVariantPathParams
	Request    shared.VariantInput `request:"mediaType=application/json"`
	Security   UpdateVariantSecurity
}

type UpdateVariant422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateVariant401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateVariantResponse struct {
	ContentType                           string
	StatusCode                            int
	UpdateVariant401ApplicationJSONObject *UpdateVariant401ApplicationJSON
	UpdateVariant404ApplicationJSONObject *UpdateVariant404ApplicationJSON
	UpdateVariant422ApplicationJSONObject *UpdateVariant422ApplicationJSON
	Variant                               *shared.Variant
}
