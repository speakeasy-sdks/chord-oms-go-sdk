package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateVariant401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateVariant422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateVariantRequest struct {
	Request  shared.VariantInput `request:"mediaType=application/json"`
	Security CreateVariantSecurity
}

type CreateVariantResponse struct {
	ContentType                           string
	StatusCode                            int64
	CreateVariant401ApplicationJSONObject *CreateVariant401ApplicationJSON
	CreateVariant422ApplicationJSONObject *CreateVariant422ApplicationJSON
	Variant                               *shared.Variant
}
