package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type CreateUserSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type CreateUserRequest struct {
	Request  shared.CreateUserInput `request:"mediaType=application/json"`
	Security CreateUserSecurity
}

type CreateUser500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type CreateUser422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type CreateUser401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type CreateUserResponse struct {
	ContentType                        string
	Headers                            map[string][]string
	StatusCode                         int
	CreateUser401ApplicationJSONObject *CreateUser401ApplicationJSON
	CreateUser422ApplicationJSONObject *CreateUser422ApplicationJSON
	CreateUser500ApplicationJSONObject *CreateUser500ApplicationJSON
	User                               *shared.User
}
