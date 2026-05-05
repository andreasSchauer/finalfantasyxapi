package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Location struct {
	ID           int32
	Name         string        `json:"location"`
	Sublocations []Sublocation `json:"sublocations"`
}

func (l Location) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", l),
		l.Name,
	}
}

func (l Location) ToKeyFields() []any {
	return []any{
		l.Name,
	}
}

func (l Location) GetID() int32 {
	return l.ID
}

func (l Location) Error() string {
	return fmt.Sprintf("location %s", l.Name)
}

func (l Location) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   l.ID,
		Name: l.Name,
	}
}
