package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateVariantRequest struct {
	Request  shared.VariantInput `request:"mediaType=application/json"`
	Security CreateVariantSecurity
}

type CreateVariant422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateVariant401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateVariantResponse struct {
	ContentType                           string
	StatusCode                            int
	CreateVariant401ApplicationJSONObject *CreateVariant401ApplicationJSON
	CreateVariant422ApplicationJSONObject *CreateVariant422ApplicationJSON
	Variant                               *shared.Variant
}
