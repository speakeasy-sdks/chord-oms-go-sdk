package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type UpdateUserPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type UpdateUserSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateUser401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type UpdateUser404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateUser422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateUserRequest struct {
	PathParams UpdateUserPathParams
	Request    shared.UserInput `request:"mediaType=application/json"`
	Security   UpdateUserSecurity
}

type UpdateUserResponse struct {
	ContentType                        string
	StatusCode                         int64
	UpdateUser401ApplicationJSONObject *UpdateUser401ApplicationJSON
	UpdateUser404ApplicationJSONObject *UpdateUser404ApplicationJSON
	UpdateUser422ApplicationJSONObject *UpdateUser422ApplicationJSON
	User                               *shared.User
}
