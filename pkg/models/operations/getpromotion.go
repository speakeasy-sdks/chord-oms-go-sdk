package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type GetPromotionPathParams struct {
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

type GetPromotionSecurity struct {
	APIKey shared.SchemeAPIKey `security:"scheme,type=apiKey,subtype=header"`
}

type GetPromotion401ApplicationJSON struct {
	Message *string `json:"message,omitempty"`
}

type GetPromotion404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type GetPromotionRequest struct {
	PathParams GetPromotionPathParams
	Security   GetPromotionSecurity
}

type GetPromotionResponse struct {
	ContentType                          string
	StatusCode                           int64
	GetPromotion401ApplicationJSONObject *GetPromotion401ApplicationJSON
	GetPromotion404ApplicationJSONObject *GetPromotion404ApplicationJSON
	Promotion                            *shared.Promotion
}
