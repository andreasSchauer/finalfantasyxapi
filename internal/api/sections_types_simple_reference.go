package api

type SimpleRef struct {
	ID            int32   `json:"id,omitempty"`
	Name          string  `json:"name"`
	Version       *int32  `json:"version,omitempty"`
	Specification *string `json:"specification,omitempty"`
}

func createSimpleRef(id int32, name string, version *int32, spec *string) SimpleRef {
	return SimpleRef{
		ID:            id,
		Name:          name,
		Version:       version,
		Specification: spec,
	}
}
