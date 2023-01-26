package shared

type Reminder struct {
	CreatedAt    *string `json:"created_at,omitempty"`
	ReminderType *string `json:"reminder_type,omitempty"`
	SentAt       *string `json:"sent_at,omitempty"`
}
