package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type ListStates200ApplicationJSON struct {
	States *shared.State `json:"states,omitempty"`
}

type ListStates401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type ListStatesResponse struct {
	ContentType                        string
	StatusCode                         int64
	ListStates200ApplicationJSONObject *ListStates200ApplicationJSON
	ListStates401ApplicationJSONObject *ListStates401ApplicationJSON
}
