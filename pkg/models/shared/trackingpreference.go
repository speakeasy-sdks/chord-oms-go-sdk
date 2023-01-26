package shared

type TrackingPreference struct {
	Destination *string `json:"destination,omitempty"`
	Preference  *bool   `json:"preference,omitempty"`
}
