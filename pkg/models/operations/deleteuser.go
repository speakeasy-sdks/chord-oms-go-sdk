package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type DeleteUserPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type DeleteUserSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type DeleteUser401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type DeleteUser404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteUser422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteUserRequest struct {
	PathParams DeleteUserPathParams
	Security   DeleteUserSecurity
}

type DeleteUserResponse struct {
	ContentType                        string
	StatusCode                         int64
	DeleteUser401ApplicationJSONObject *DeleteUser401ApplicationJSON
	DeleteUser404ApplicationJSONObject *DeleteUser404ApplicationJSON
	DeleteUser422ApplicationJSONObject *DeleteUser422ApplicationJSON
}
