package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type PatchUsersUserIDSubscriptionsIDPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID int64  `pathParam:"style=simple,explode=false,name=user_id"`
}

type PatchUsersUserIDSubscriptionsIDSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type PatchUsersUserIDSubscriptionsIDRequest struct {
	PathParams PatchUsersUserIDSubscriptionsIDPathParams
	Request    shared.SubscriptionInput `request:"mediaType=application/json"`
	Security   PatchUsersUserIDSubscriptionsIDSecurity
}

type PatchUsersUserIDSubscriptionsID404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PatchUsersUserIDSubscriptionsID401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PatchUsersUserIDSubscriptionsIDResponse struct {
	ContentType                                             string
	StatusCode                                              int
	PatchUsersUserIDSubscriptionsID401ApplicationJSONObject *PatchUsersUserIDSubscriptionsID401ApplicationJSON
	PatchUsersUserIDSubscriptionsID404ApplicationJSONObject *PatchUsersUserIDSubscriptionsID404ApplicationJSON
	Subscription                                            *shared.Subscription
}
