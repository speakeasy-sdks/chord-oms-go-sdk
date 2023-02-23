package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateProductPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateProductSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateProductRequest struct {
	PathParams UpdateProductPathParams
	Request    shared.ProductInput `request:"mediaType=application/json"`
	Security   UpdateProductSecurity
}

type UpdateProduct422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateProduct404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateProduct401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateProductResponse struct {
	ContentType                           string
	StatusCode                            int
	Product                               *shared.Product
	UpdateProduct401ApplicationJSONObject *UpdateProduct401ApplicationJSON
	UpdateProduct404ApplicationJSONObject *UpdateProduct404ApplicationJSON
	UpdateProduct422ApplicationJSONObject *UpdateProduct422ApplicationJSON
}
