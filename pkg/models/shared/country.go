package shared

type Country struct {
	ID      *int64  `json:"id,omitempty"`
	Iso     *string `json:"iso,omitempty"`
	Iso3    *string `json:"iso3,omitempty"`
	IsoName *string `json:"iso_name,omitempty"`
	Name    *string `json:"name,omitempty"`
	Numcode *int64  `json:"numcode,omitempty"`
}
