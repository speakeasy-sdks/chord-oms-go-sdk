package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteStorePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteStoreSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteStoreRequest struct {
	PathParams DeleteStorePathParams
	Security   DeleteStoreSecurity
}

type DeleteStore422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStore404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStore401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStoreResponse struct {
	ContentType                         string
	StatusCode                          int
	DeleteStore401ApplicationJSONObject *DeleteStore401ApplicationJSON
	DeleteStore404ApplicationJSONObject *DeleteStore404ApplicationJSON
	DeleteStore422ApplicationJSONObject *DeleteStore422ApplicationJSON
}
