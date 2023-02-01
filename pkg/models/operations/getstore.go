package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetStorePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetStoreSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetStore401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStore404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStoreRequest struct {
	PathParams GetStorePathParams
	Security   GetStoreSecurity
}

type GetStoreResponse struct {
	ContentType                      string
	StatusCode                       int64
	GetStore401ApplicationJSONObject *GetStore401ApplicationJSON
	GetStore404ApplicationJSONObject *GetStore404ApplicationJSON
	Store                            *shared.Store
}
