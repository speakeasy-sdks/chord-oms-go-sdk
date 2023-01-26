package shared

type Promotion struct {
	Advertise   *bool   `json:"advertise,omitempty"`
	Description *string `json:"description,omitempty"`
	ExpiresAt   *string `json:"expires_at,omitempty"`
	ID          *int64  `json:"id,omitempty"`
	MatchPolicy *string `json:"match_policy,omitempty"`
	Name        *string `json:"name,omitempty"`
	Path        *string `json:"path,omitempty"`
	StartsAt    *string `json:"starts_at,omitempty"`
	Type        *string `json:"type,omitempty"`
	UsageLimit  *int64  `json:"usage_limit,omitempty"`
}
