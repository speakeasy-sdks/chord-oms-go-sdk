package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/v2/pkg/models/shared"
)

type RedeemGiftCardRequest struct {
	Request shared.RedeemInput `request:"mediaType=application/json"`
}

type RedeemGiftCard404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type RedeemGiftCardResponse struct {
	ContentType                            string
	StatusCode                             int
	GiftCard                               *shared.GiftCard
	RedeemGiftCard404ApplicationJSONObject *RedeemGiftCard404ApplicationJSON
}
