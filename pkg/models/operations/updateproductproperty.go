package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateProductPropertyPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type UpdateProductPropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateProductPropertyRequest struct {
	PathParams UpdateProductPropertyPathParams
	Request    shared.ProductPropertyInput `request:"mediaType=application/json"`
	Security   UpdateProductPropertySecurity
}

type UpdateProductProperty422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateProductProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateProductProperty401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateProductPropertyResponse struct {
	ContentType                                   string
	StatusCode                                    int
	ProductProperty                               *shared.ProductProperty
	UpdateProductProperty401ApplicationJSONObject *UpdateProductProperty401ApplicationJSON
	UpdateProductProperty404ApplicationJSONObject *UpdateProductProperty404ApplicationJSON
	UpdateProductProperty422ApplicationJSONObject *UpdateProductProperty422ApplicationJSON
}
