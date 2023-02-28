package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type ListStates401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type ListStates200ApplicationJSON struct {
	States *shared.State `json:"states,omitempty"`
}

type ListStatesResponse struct {
	ContentType                        string
	StatusCode                         int
	ListStates200ApplicationJSONObject *ListStates200ApplicationJSON
	ListStates401ApplicationJSONObject *ListStates401ApplicationJSON
}
