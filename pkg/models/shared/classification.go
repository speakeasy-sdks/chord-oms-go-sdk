package shared

type Classification struct {
	Position *int64 `json:"position,omitempty"`
	Taxon    *Taxon `json:"taxon,omitempty"`
	TaxonID  *int64 `json:"taxon_id,omitempty"`
}
