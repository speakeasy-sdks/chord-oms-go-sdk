package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateStorePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateStoreSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type UpdateStoreRequest struct {
	PathParams UpdateStorePathParams
	Request    shared.StoreInput `request:"mediaType=application/json"`
	Security   UpdateStoreSecurity
}

type UpdateStore422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateStore404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStore401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStoreResponse struct {
	ContentType                         string
	StatusCode                          int
	Store                               *shared.Store
	UpdateStore401ApplicationJSONObject *UpdateStore401ApplicationJSON
	UpdateStore404ApplicationJSONObject *UpdateStore404ApplicationJSON
	UpdateStore422ApplicationJSONObject *UpdateStore422ApplicationJSON
}
