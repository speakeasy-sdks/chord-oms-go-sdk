package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteVariantPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteVariantRequest struct {
	PathParams DeleteVariantPathParams
	Security   DeleteVariantSecurity
}

type DeleteVariant422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteVariant401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteVariantResponse struct {
	ContentType                           string
	StatusCode                            int
	DeleteVariant401ApplicationJSONObject *DeleteVariant401ApplicationJSON
	DeleteVariant404ApplicationJSONObject *DeleteVariant404ApplicationJSON
	DeleteVariant422ApplicationJSONObject *DeleteVariant422ApplicationJSON
}
