package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetUserPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetUserSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetUser401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetUser404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUserRequest struct {
	PathParams GetUserPathParams
	Security   GetUserSecurity
}

type GetUserResponse struct {
	ContentType                     string
	StatusCode                      int64
	GetUser401ApplicationJSONObject *GetUser401ApplicationJSON
	GetUser404ApplicationJSONObject *GetUser404ApplicationJSON
	User                            *shared.User
}
