package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type PostUsersUserIDSubscriptionsIDSkipPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID int64  `pathParam:"style=simple,explode=false,name=user_id"`
}

type PostUsersUserIDSubscriptionsIDSkipSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type PostUsersUserIDSubscriptionsIDSkipRequest struct {
	PathParams PostUsersUserIDSubscriptionsIDSkipPathParams
	Security   PostUsersUserIDSubscriptionsIDSkipSecurity
}

type PostUsersUserIDSubscriptionsIDSkipResponse struct {
	ContentType  string
	StatusCode   int
	Subscription *shared.Subscription
}
