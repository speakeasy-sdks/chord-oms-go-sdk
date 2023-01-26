package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateProductSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateProduct401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateProduct422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateProductRequest struct {
	Request  shared.ProductInput `request:"mediaType=application/json"`
	Security CreateProductSecurity
}

type CreateProductResponse struct {
	ContentType                           string
	StatusCode                            int64
	CreateProduct401ApplicationJSONObject *CreateProduct401ApplicationJSON
	CreateProduct422ApplicationJSONObject *CreateProduct422ApplicationJSON
	Product                               *shared.Product
}
