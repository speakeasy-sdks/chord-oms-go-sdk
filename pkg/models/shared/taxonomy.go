package shared

type Taxonomy struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Root *Taxon  `json:"root,omitempty"`
}
