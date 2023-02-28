package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateProductSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateProductRequest struct {
	Request  shared.ProductInput `request:"mediaType=application/json"`
	Security CreateProductSecurity
}

type CreateProduct422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateProduct401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateProductResponse struct {
	ContentType                           string
	StatusCode                            int
	CreateProduct401ApplicationJSONObject *CreateProduct401ApplicationJSON
	CreateProduct422ApplicationJSONObject *CreateProduct422ApplicationJSON
	Product                               *shared.Product
}
