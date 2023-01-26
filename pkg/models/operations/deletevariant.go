package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteVariantPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteVariantSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteVariant204ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteVariant401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteVariant404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteVariant422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteVariantRequest struct {
	PathParams DeleteVariantPathParams
	Security   DeleteVariantSecurity
}

type DeleteVariantResponse struct {
	ContentType                           string
	StatusCode                            int64
	DeleteVariant204ApplicationJSONObject *DeleteVariant204ApplicationJSON
	DeleteVariant401ApplicationJSONObject *DeleteVariant401ApplicationJSON
	DeleteVariant404ApplicationJSONObject *DeleteVariant404ApplicationJSON
	DeleteVariant422ApplicationJSONObject *DeleteVariant422ApplicationJSON
}
