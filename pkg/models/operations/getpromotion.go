package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type GetPromotionPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetPromotionSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=http,subtype=bearer"`
}

type GetPromotionRequest struct {
	PathParams GetPromotionPathParams
	Security   GetPromotionSecurity
}

type GetPromotion404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetPromotion401ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetPromotionResponse struct {
	ContentType                          string
	StatusCode                           int
	GetPromotion401ApplicationJSONObject *GetPromotion401ApplicationJSON
	GetPromotion404ApplicationJSONObject *GetPromotion404ApplicationJSON
	Promotion                            *shared.Promotion
}
