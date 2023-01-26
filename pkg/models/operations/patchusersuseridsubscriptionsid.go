package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type PatchUsersUserIDSubscriptionsIDPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type PatchUsersUserIDSubscriptionsIDSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PatchUsersUserIDSubscriptionsID401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PatchUsersUserIDSubscriptionsID404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PatchUsersUserIDSubscriptionsIDRequest struct {
	PathParams PatchUsersUserIDSubscriptionsIDPathParams
	Request    shared.SubscriptionInput `request:"mediaType=application/json"`
	Security   PatchUsersUserIDSubscriptionsIDSecurity
}

type PatchUsersUserIDSubscriptionsIDResponse struct {
	ContentType                                             string
	StatusCode                                              int64
	PatchUsersUserIDSubscriptionsID401ApplicationJSONObject *PatchUsersUserIDSubscriptionsID401ApplicationJSON
	PatchUsersUserIDSubscriptionsID404ApplicationJSONObject *PatchUsersUserIDSubscriptionsID404ApplicationJSON
	Subscription                                            *shared.Subscription
}
