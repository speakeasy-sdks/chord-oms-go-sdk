package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetProductPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetProductSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetProductRequest struct {
	PathParams GetProductPathParams
	Security   GetProductSecurity
}

type GetProduct404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProduct401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductResponse struct {
	ContentType                        string
	StatusCode                         int
	GetProduct401ApplicationJSONObject *GetProduct401ApplicationJSON
	GetProduct404ApplicationJSONObject *GetProduct404ApplicationJSON
	Product                            *shared.Product
}
