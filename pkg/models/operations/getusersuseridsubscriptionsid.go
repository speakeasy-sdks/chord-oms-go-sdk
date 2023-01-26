package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetUsersUserIDSubscriptionsIDPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type GetUsersUserIDSubscriptionsIDSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetUsersUserIDSubscriptionsID401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUsersUserIDSubscriptionsID404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUsersUserIDSubscriptionsIDRequest struct {
	PathParams GetUsersUserIDSubscriptionsIDPathParams
	Security   GetUsersUserIDSubscriptionsIDSecurity
}

type GetUsersUserIDSubscriptionsIDResponse struct {
	ContentType                                           string
	StatusCode                                            int64
	GetUsersUserIDSubscriptionsID401ApplicationJSONObject *GetUsersUserIDSubscriptionsID401ApplicationJSON
	GetUsersUserIDSubscriptionsID404ApplicationJSONObject *GetUsersUserIDSubscriptionsID404ApplicationJSON
	Subscription                                          *shared.Subscription
}
