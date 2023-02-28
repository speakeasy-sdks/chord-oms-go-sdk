package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetCurrentUserSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type GetCurrentUserRequest struct {
	Security GetCurrentUserSecurity
}

type GetCurrentUser500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type GetCurrentUser404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCurrentUser401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetCurrentUserResponse struct {
	ContentType                            string
	StatusCode                             int
	GetCurrentUser401ApplicationJSONObject *GetCurrentUser401ApplicationJSON
	GetCurrentUser404ApplicationJSONObject *GetCurrentUser404ApplicationJSON
	GetCurrentUser500ApplicationJSONObject *GetCurrentUser500ApplicationJSON
	User                                   *shared.User
}
