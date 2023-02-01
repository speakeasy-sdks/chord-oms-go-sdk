package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateStorePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateStoreSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateStore401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStore404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateStore422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateStoreRequest struct {
	PathParams UpdateStorePathParams
	Request    shared.StoreInput `request:"mediaType=application/json"`
	Security   UpdateStoreSecurity
}

type UpdateStoreResponse struct {
	ContentType                         string
	StatusCode                          int64
	Store                               *shared.Store
	UpdateStore401ApplicationJSONObject *UpdateStore401ApplicationJSON
	UpdateStore404ApplicationJSONObject *UpdateStore404ApplicationJSON
	UpdateStore422ApplicationJSONObject *UpdateStore422ApplicationJSON
}
