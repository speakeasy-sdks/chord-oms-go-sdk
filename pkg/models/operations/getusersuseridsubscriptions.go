package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetUsersUserIDSubscriptionsPathParams struct {
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type GetUsersUserIDSubscriptionsSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetUsersUserIDSubscriptions200ApplicationJSON struct {
	Subscriptions []shared.Subscription `json:"subscriptions,omitempty"`
}

type GetUsersUserIDSubscriptions401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUsersUserIDSubscriptions404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetUsersUserIDSubscriptionsRequest struct {
	PathParams GetUsersUserIDSubscriptionsPathParams
	Security   GetUsersUserIDSubscriptionsSecurity
}

type GetUsersUserIDSubscriptionsResponse struct {
	ContentType                                         string
	StatusCode                                          int64
	GetUsersUserIDSubscriptions200ApplicationJSONObject *GetUsersUserIDSubscriptions200ApplicationJSON
	GetUsersUserIDSubscriptions401ApplicationJSONObject *GetUsersUserIDSubscriptions401ApplicationJSON
	GetUsersUserIDSubscriptions404ApplicationJSONObject *GetUsersUserIDSubscriptions404ApplicationJSON
}
