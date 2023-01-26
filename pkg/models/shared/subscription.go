package shared

type Subscription struct {
	ActionableDate *string                `json:"actionable_date,omitempty"`
	BillAddress    *Address               `json:"bill_address,omitempty"`
	CreatedAt      *string                `json:"created_at,omitempty"`
	EndDate        *string                `json:"end_date,omitempty"`
	ID             *float64               `json:"id,omitempty"`
	Interval       *string                `json:"interval,omitempty"`
	IntervalLength *float64               `json:"interval_length,omitempty"`
	IntervalUnits  *string                `json:"interval_units,omitempty"`
	LastRemindedAt *string                `json:"last_reminded_at,omitempty"`
	LineItems      []SubscriptionLineItem `json:"line_items,omitempty"`
	Reminders      []Reminder             `json:"reminders,omitempty"`
	ShipAddress    *Address               `json:"ship_address,omitempty"`
	State          *string                `json:"state,omitempty"`
}
