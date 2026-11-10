package sections

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"

func createSimpleRef(id int32, name string, version *int32, spec *string) sectiontypes.SimpleRef {
	return sectiontypes.SimpleRef{
		ID:            id,
		Name:          name,
		Version:       version,
		Specification: spec,
	}
}
