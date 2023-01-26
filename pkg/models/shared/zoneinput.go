package shared

type ZoneInputZoneMembersAttributes struct {
	ZoneableID   *int64  `json:"zoneable_id,omitempty"`
	ZoneableType *string `json:"zoneable_type,omitempty"`
}

type ZoneInput struct {
	Description           *string                          `json:"description,omitempty"`
	Name                  *string                          `json:"name,omitempty"`
	ZoneMembersAttributes []ZoneInputZoneMembersAttributes `json:"zone_members_attributes,omitempty"`
}
