package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type CreateUserSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type CreateUser401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type CreateUser422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateUserRequest struct {
	Request  shared.UserInput `request:"mediaType=application/json"`
	Security CreateUserSecurity
}

type CreateUserResponse struct {
	ContentType                        string
	StatusCode                         int64
	CreateUser401ApplicationJSONObject *CreateUser401ApplicationJSON
	CreateUser422ApplicationJSONObject *CreateUser422ApplicationJSON
	User                               *shared.User
}
