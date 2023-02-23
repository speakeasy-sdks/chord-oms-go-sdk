package shared

// GiftNote
// An optional note to attach to the order. For some fulfillment services,
// this note might be printed and packaged with the order.
type GiftNote struct {
	From      *string `json:"from,omitempty"`
	Note      *string `json:"note,omitempty"`
	Recipient *string `json:"recipient,omitempty"`
}
