package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateProductPropertyPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type UpdateProductPropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateProductProperty401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateProductProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateProductProperty422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateProductPropertyRequest struct {
	PathParams UpdateProductPropertyPathParams
	Request    shared.ProductPropertyInput `request:"mediaType=application/json"`
	Security   UpdateProductPropertySecurity
}

type UpdateProductPropertyResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	ProductProperty                               *shared.ProductProperty
	UpdateProductProperty401ApplicationJSONObject *UpdateProductProperty401ApplicationJSON
	UpdateProductProperty404ApplicationJSONObject *UpdateProductProperty404ApplicationJSON
	UpdateProductProperty422ApplicationJSONObject *UpdateProductProperty422ApplicationJSON
}
