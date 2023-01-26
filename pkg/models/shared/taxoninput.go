package shared

type TaxonInput struct {
	ChildIndex      *int64  `json:"child_index,omitempty"`
	Description     *string `json:"description,omitempty"`
	Icon            *string `json:"icon,omitempty"`
	MetaDescription *string `json:"meta_description,omitempty"`
	MetaKeywords    *string `json:"meta_keywords,omitempty"`
	MetaTitle       *string `json:"meta_title,omitempty"`
	Name            *string `json:"name,omitempty"`
	ParentID        *int64  `json:"parent_id,omitempty"`
	Permalink       *string `json:"permalink,omitempty"`
	Position        *int64  `json:"position,omitempty"`
	TaxonomyID      *int64  `json:"taxonomy_id,omitempty"`
}
