package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetProductPropertyPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type GetProductPropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetProductProperty401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductPropertyRequest struct {
	PathParams GetProductPropertyPathParams
	Security   GetProductPropertySecurity
}

type GetProductPropertyResponse struct {
	ContentType                                string
	StatusCode                                 int64
	GetProductProperty401ApplicationJSONObject *GetProductProperty401ApplicationJSON
	GetProductProperty404ApplicationJSONObject *GetProductProperty404ApplicationJSON
	ProductProperty                            *shared.ProductProperty
}
