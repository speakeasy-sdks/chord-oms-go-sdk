package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetUserPathParams struct {
	UserID int64 `pathParam:"style=simple,explode=false,name=user_id"`
}

type GetUserSecurity struct {
	APIKey          *shared.SchemeAPIKey          `security:"scheme,type=http,subtype=bearer"`
	StorefrontLogin *shared.SchemeStorefrontLogin `security:"scheme,type=apiKey,subtype=header"`
}

type GetUserRequest struct {
	PathParams GetUserPathParams
	Security   GetUserSecurity
}

type GetUser500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type GetUser404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUser401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUserResponse struct {
	ContentType                     string
	StatusCode                      int
	GetUser401ApplicationJSONObject *GetUser401ApplicationJSON
	GetUser404ApplicationJSONObject *GetUser404ApplicationJSON
	GetUser500ApplicationJSONObject *GetUser500ApplicationJSON
	User                            *shared.User
}
