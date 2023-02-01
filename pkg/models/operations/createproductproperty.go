package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateProductPropertyPathParams struct {
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type CreateProductPropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateProductProperty401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateProductProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateProductProperty422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateProductPropertyRequest struct {
	PathParams CreateProductPropertyPathParams
	Request    shared.ProductPropertyInput `request:"mediaType=application/json"`
	Security   CreateProductPropertySecurity
}

type CreateProductPropertyResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	CreateProductProperty401ApplicationJSONObject *CreateProductProperty401ApplicationJSON
	CreateProductProperty404ApplicationJSONObject *CreateProductProperty404ApplicationJSON
	CreateProductProperty422ApplicationJSONObject *CreateProductProperty422ApplicationJSON
	ProductProperty                               *shared.ProductProperty
}
