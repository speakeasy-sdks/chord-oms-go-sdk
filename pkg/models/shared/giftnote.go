package shared

type GiftNote struct {
	From      *string `json:"from,omitempty"`
	Note      *string `json:"note,omitempty"`
	Recipient *string `json:"recipient,omitempty"`
}
