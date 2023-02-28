package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetUsersUserIDSubscriptionsPathParams struct {
	UserID int64 `pathParam:"style=simple,explode=false,name=user_id"`
}

type GetUsersUserIDSubscriptionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetUsersUserIDSubscriptionsRequest struct {
	PathParams GetUsersUserIDSubscriptionsPathParams
	Security   GetUsersUserIDSubscriptionsSecurity
}

type GetUsersUserIDSubscriptions404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUsersUserIDSubscriptions401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUsersUserIDSubscriptions200ApplicationJSON struct {
	Subscriptions []shared.Subscription `json:"subscriptions,omitempty"`
}

type GetUsersUserIDSubscriptionsResponse struct {
	ContentType                                         string
	StatusCode                                          int
	GetUsersUserIDSubscriptions200ApplicationJSONObject *GetUsersUserIDSubscriptions200ApplicationJSON
	GetUsersUserIDSubscriptions401ApplicationJSONObject *GetUsersUserIDSubscriptions401ApplicationJSON
	GetUsersUserIDSubscriptions404ApplicationJSONObject *GetUsersUserIDSubscriptions404ApplicationJSON
}
