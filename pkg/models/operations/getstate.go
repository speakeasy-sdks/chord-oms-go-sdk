package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetStatePathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetState401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetState404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetStateRequest struct {
	PathParams GetStatePathParams
}

type GetStateResponse struct {
	ContentType                      string
	StatusCode                       int64
	GetState401ApplicationJSONObject *GetState401ApplicationJSON
	GetState404ApplicationJSONObject *GetState404ApplicationJSON
	State                            *shared.State
}
