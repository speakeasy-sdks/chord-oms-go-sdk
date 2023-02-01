package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetProductPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetProductSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetProduct401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProduct404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetProductRequest struct {
	PathParams GetProductPathParams
	Security   GetProductSecurity
}

type GetProductResponse struct {
	ContentType                        string
	StatusCode                         int64
	GetProduct401ApplicationJSONObject *GetProduct401ApplicationJSON
	GetProduct404ApplicationJSONObject *GetProduct404ApplicationJSON
	Product                            *shared.Product
}
