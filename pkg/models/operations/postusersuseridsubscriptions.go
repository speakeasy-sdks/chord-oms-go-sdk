package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type PostUsersUserIDSubscriptionsPathParams struct {
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type PostUsersUserIDSubscriptions401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PostUsersUserIDSubscriptions404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PostUsersUserIDSubscriptionsRequest struct {
	PathParams PostUsersUserIDSubscriptionsPathParams
	Request    shared.Subscription `request:"mediaType=application/json"`
}

type PostUsersUserIDSubscriptionsResponse struct {
	ContentType                                          string
	StatusCode                                           int64
	PostUsersUserIDSubscriptions401ApplicationJSONObject *PostUsersUserIDSubscriptions401ApplicationJSON
	PostUsersUserIDSubscriptions404ApplicationJSONObject *PostUsersUserIDSubscriptions404ApplicationJSON
	Subscription                                         *shared.Subscription
}
