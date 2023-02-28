package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type DeleteUserPathParams struct {
	UserID int64 `pathParam:"style=simple,explode=false,name=user_id"`
}

type DeleteUserSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type DeleteUserRequest struct {
	PathParams DeleteUserPathParams
	Security   DeleteUserSecurity
}

type DeleteUser500ApplicationJSON struct {
	Error  *string  `json:"error,omitempty"`
	Status *float64 `json:"status,omitempty"`
}

type DeleteUser422ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteUser404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteUser401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type DeleteUserResponse struct {
	ContentType                        string
	StatusCode                         int
	DeleteUser401ApplicationJSONObject *DeleteUser401ApplicationJSON
	DeleteUser404ApplicationJSONObject *DeleteUser404ApplicationJSON
	DeleteUser422ApplicationJSONObject *DeleteUser422ApplicationJSON
	DeleteUser500ApplicationJSONObject *DeleteUser500ApplicationJSON
}
