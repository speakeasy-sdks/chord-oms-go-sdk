package operations

import (
	"github.com/speakeasy-sdks/chord-oms-go-sdk/pkg/models/shared"
)

type RedeemGiftCard404ApplicationJSON struct {
	Error *string `json:"error,omitempty"`
}

type RedeemGiftCardRequest struct {
	Request shared.RedeemInput `request:"mediaType=application/json"`
}

type RedeemGiftCardResponse struct {
	ContentType                            string
	StatusCode                             int64
	GiftCard                               *shared.GiftCard
	RedeemGiftCard404ApplicationJSONObject *RedeemGiftCard404ApplicationJSON
}
