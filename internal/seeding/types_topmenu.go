package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Topmenu struct {
	ID   int32
	Name string `json:"name"`
}

func (t Topmenu) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", t),
		t.Name,
	}
}

func (t Topmenu) ToKeyFields() []any {
	return []any{
		t.Name,
	}
}

func (t Topmenu) GetID() int32 {
	return t.ID
}

func (t Topmenu) Error() string {
	return fmt.Sprintf("topmenu %s", t.Name)
}

func (t Topmenu) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   t.ID,
		Name: t.Name,
	}
}
