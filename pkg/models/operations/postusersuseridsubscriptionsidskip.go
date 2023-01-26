package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type PostUsersUserIDSubscriptionsIDSkipPathParams struct {
	ID     string `pathParam:"style=simple,explode=false,name=id"`
	UserID string `pathParam:"style=simple,explode=false,name=user_id"`
}

type PostUsersUserIDSubscriptionsIDSkipSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type PostUsersUserIDSubscriptionsIDSkipRequest struct {
	PathParams PostUsersUserIDSubscriptionsIDSkipPathParams
	Security   PostUsersUserIDSubscriptionsIDSkipSecurity
}

type PostUsersUserIDSubscriptionsIDSkipResponse struct {
	ContentType  string
	StatusCode   int64
	Subscription *shared.Subscription
}
