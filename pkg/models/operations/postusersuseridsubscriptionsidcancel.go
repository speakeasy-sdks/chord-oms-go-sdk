package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type PostUsersUserIDSubscriptionsIDCancelPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID int64  `pathParam:"style=simple,explode=false,name=user_id"`
}

type PostUsersUserIDSubscriptionsIDCancelSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type PostUsersUserIDSubscriptionsIDCancelRequest struct {
	PathParams PostUsersUserIDSubscriptionsIDCancelPathParams
	Security   PostUsersUserIDSubscriptionsIDCancelSecurity
}

type PostUsersUserIDSubscriptionsIDCancel404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PostUsersUserIDSubscriptionsIDCancelResponse struct {
	ContentType                                                  string
	StatusCode                                                   int
	PostUsersUserIDSubscriptionsIDCancel404ApplicationJSONObject *PostUsersUserIDSubscriptionsIDCancel404ApplicationJSON
	Subscription                                                 *shared.Subscription
}
