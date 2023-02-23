package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetStatePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetStateRequest struct {
	PathParams GetStatePathParams
}

type GetState404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetState401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStateResponse struct {
	ContentType                      string
	StatusCode                       int
	GetState401ApplicationJSONObject *GetState401ApplicationJSON
	GetState404ApplicationJSONObject *GetState404ApplicationJSON
	State                            *shared.State
}
