package shared

type Taxon struct {
	ID         *int64  `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	ParentID   *int64  `json:"parent_id,omitempty"`
	Permalink  *string `json:"permalink,omitempty"`
	PrettyName *string `json:"pretty_name,omitempty"`
	TaxonomyID *int64  `json:"taxonomy_id,omitempty"`
	Taxons     []Taxon `json:"taxons,omitempty"`
}
