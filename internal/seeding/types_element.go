package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Element struct {
	ID                int32
	Name              string  `json:"name"`
	OppositeElement   *string `json:"opposite_element"`
	OppositeElementID *int32
}

func (e Element) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
		e.Name,
		h.DerefOrNil(e.OppositeElementID),
	}
}

func (e Element) ToKeyFields() []any {
	return []any{
		e.Name,
	}
}

func (e Element) GetID() int32 {
	return e.ID
}

func (e Element) Error() string {
	return fmt.Sprintf("element %s", e.Name)
}

func (e Element) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   e.ID,
		Name: e.Name,
	}
}
