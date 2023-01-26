package shared

type ImageInput struct {
	Alt          *string `json:"alt,omitempty"`
	Attachment   *string `json:"attachment,omitempty"`
	Position     *int64  `json:"position,omitempty"`
	ViewableID   *int64  `json:"viewable_id,omitempty"`
	ViewableType *string `json:"viewable_type,omitempty"`
}
