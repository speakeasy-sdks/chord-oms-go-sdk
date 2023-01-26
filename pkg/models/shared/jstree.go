package shared

type JstreeAttr struct {
	ID   *int64  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

type Jstree struct {
	Attr  *JstreeAttr `json:"attr,omitempty"`
	Data  *string     `json:"data,omitempty"`
	State *string     `json:"state,omitempty"`
}
