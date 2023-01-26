package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type PostUsersUserIDSubscriptionsIDCancelPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type PostUsersUserIDSubscriptionsIDCancelSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostUsersUserIDSubscriptionsIDCancel404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type PostUsersUserIDSubscriptionsIDCancelRequest struct {
	PathParams PostUsersUserIDSubscriptionsIDCancelPathParams
	Security   PostUsersUserIDSubscriptionsIDCancelSecurity
}

type PostUsersUserIDSubscriptionsIDCancelResponse struct {
	ContentType                                                  string
	StatusCode                                                   int64
	PostUsersUserIDSubscriptionsIDCancel404ApplicationJSONObject *PostUsersUserIDSubscriptionsIDCancel404ApplicationJSON
	Subscription                                                 *shared.Subscription
}
