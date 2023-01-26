package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetUsersMeSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetUsersMe401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetUsersMe404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUsersMeRequest struct {
	Security GetUsersMeSecurity
}

type GetUsersMeResponse struct {
	ContentType                        string
	StatusCode                         int64
	GetUsersMe401ApplicationJSONObject *GetUsersMe401ApplicationJSON
	GetUsersMe404ApplicationJSONObject *GetUsersMe404ApplicationJSON
	User                               *shared.User
}
