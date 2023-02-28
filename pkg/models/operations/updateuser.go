package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type UpdateUserPathParams struct {
	UserID int64 `pathParam:"style=simple,explode=false,name=user_id"`
}

type UpdateUserSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type UpdateUserRequest struct {
	PathParams UpdateUserPathParams
	Request    map[string]interface{} `request:"mediaType=application/json"`
	Security   UpdateUserSecurity
}

type UpdateUser500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type UpdateUser422ApplicationJSON struct {
	Error  *string                `json:"error,omitempty"`
	Errors map[string]interface{} `json:"errors,omitempty"`
}

type UpdateUser404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateUser401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type UpdateUserResponse struct {
	ContentType                        string
	StatusCode                         int
	UpdateUser401ApplicationJSONObject *UpdateUser401ApplicationJSON
	UpdateUser404ApplicationJSONObject *UpdateUser404ApplicationJSON
	UpdateUser422ApplicationJSONObject *UpdateUser422ApplicationJSON
	UpdateUser500ApplicationJSONObject *UpdateUser500ApplicationJSON
	User                               *shared.User
}
