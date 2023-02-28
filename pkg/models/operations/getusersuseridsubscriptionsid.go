package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetUsersUserIDSubscriptionsIDPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID int64  `pathParam:"style=simple,explode=false,name=user_id"`
}

type GetUsersUserIDSubscriptionsIDSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetUsersUserIDSubscriptionsIDRequest struct {
	PathParams GetUsersUserIDSubscriptionsIDPathParams
	Security   GetUsersUserIDSubscriptionsIDSecurity
}

type GetUsersUserIDSubscriptionsID404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUsersUserIDSubscriptionsID401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUsersUserIDSubscriptionsIDResponse struct {
	ContentType                                           string
	StatusCode                                            int
	GetUsersUserIDSubscriptionsID401ApplicationJSONObject *GetUsersUserIDSubscriptionsID401ApplicationJSON
	GetUsersUserIDSubscriptionsID404ApplicationJSONObject *GetUsersUserIDSubscriptionsID404ApplicationJSON
	Subscription                                          *shared.Subscription
}
