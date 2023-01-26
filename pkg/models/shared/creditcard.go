package shared

type CreditCard struct {
	Address    *interface{} `json:"address,omitempty"`
	CcType     *string      `json:"cc_type,omitempty"`
	ID         *int64       `json:"id,omitempty"`
	LastDigits *string      `json:"last_digits,omitempty"`
	Month      *string      `json:"month,omitempty"`
	Name       *string      `json:"name,omitempty"`
	Year       *string      `json:"year,omitempty"`
}
