package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type PostUsersUserIDSubscriptionsPathParams struct {
	UserID int64 `pathParam:"style=simple,explode=false,name=user_id"`
}

type PostUsersUserIDSubscriptionsRequest struct {
	PathParams PostUsersUserIDSubscriptionsPathParams
	Request    shared.Subscription `request:"mediaType=application/json"`
}

type PostUsersUserIDSubscriptions404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PostUsersUserIDSubscriptions401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PostUsersUserIDSubscriptionsResponse struct {
	ContentType                                          string
	StatusCode                                           int
	PostUsersUserIDSubscriptions401ApplicationJSONObject *PostUsersUserIDSubscriptions401ApplicationJSON
	PostUsersUserIDSubscriptions404ApplicationJSONObject *PostUsersUserIDSubscriptions404ApplicationJSON
	Subscription                                         *shared.Subscription
}
