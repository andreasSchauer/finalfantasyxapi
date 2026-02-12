package api

type SubRef struct {
	ID            int32   `json:"id,omitempty"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
}

func createSubReference(id int32, name string, version *int32, spec *string) SubRef {
	return SubRef{
		ID:            id,
		Name:          name,
		Version:       version,
		Specification: spec,
	}
}
