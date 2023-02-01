package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteProductPropertyPathParams struct {
	ID        string `pathParam:"style=simple,explode=false,name=id"`
	ProductID string `pathParam:"style=simple,explode=false,name=product_id"`
}

type DeleteProductPropertySecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteProductProperty401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductProperty404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductProperty422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteProductPropertyRequest struct {
	PathParams DeleteProductPropertyPathParams
	Security   DeleteProductPropertySecurity
}

type DeleteProductPropertyResponse struct {
	ContentType                                   string
	StatusCode                                    int64
	DeleteProductProperty401ApplicationJSONObject *DeleteProductProperty401ApplicationJSON
	DeleteProductProperty404ApplicationJSONObject *DeleteProductProperty404ApplicationJSON
	DeleteProductProperty422ApplicationJSONObject *DeleteProductProperty422ApplicationJSON
}
