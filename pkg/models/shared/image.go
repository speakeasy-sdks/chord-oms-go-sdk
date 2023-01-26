package shared

type Image struct {
	Alt                   *string `json:"alt,omitempty"`
	AttachmentContentType *string `json:"attachment_content_type,omitempty"`
	AttachmentFileName    *string `json:"attachment_file_name,omitempty"`
	AttachmentHeight      *int64  `json:"attachment_height,omitempty"`
	AttachmentUpdatedAt   *string `json:"attachment_updated_at,omitempty"`
	AttachmentWidth       *int64  `json:"attachment_width,omitempty"`
	ID                    *int64  `json:"id,omitempty"`
	LargeURL              *string `json:"large_url,omitempty"`
	MiniURL               *string `json:"mini_url,omitempty"`
	Position              *int64  `json:"position,omitempty"`
	ProductURL            *string `json:"product_url,omitempty"`
	SmallURL              *string `json:"small_url,omitempty"`
	Type                  *string `json:"type,omitempty"`
	ViewableID            *int64  `json:"viewable_id,omitempty"`
	ViewableType          *string `json:"viewable_type,omitempty"`
}
