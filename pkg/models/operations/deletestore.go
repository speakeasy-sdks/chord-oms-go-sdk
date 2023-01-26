package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteStorePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteStoreSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteStore401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteStore404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStore422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteStoreRequest struct {
	PathParams DeleteStorePathParams
	Security   DeleteStoreSecurity
}

type DeleteStoreResponse struct {
	ContentType                         string
	StatusCode                          int64
	DeleteStore401ApplicationJSONObject *DeleteStore401ApplicationJSON
	DeleteStore404ApplicationJSONObject *DeleteStore404ApplicationJSON
	DeleteStore422ApplicationJSONObject *DeleteStore422ApplicationJSON
}
