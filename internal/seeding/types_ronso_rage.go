package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type RonsoRage struct {
	ID int32
	Overdrive
}

func (r RonsoRage) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", r),
		r.Overdrive.ID,
	}
}

func (r RonsoRage) ToKeyFields() []any {
	return []any{
		r.Name,
	}
}

func (r RonsoRage) GetID() int32 {
	return r.ID
}

func (r RonsoRage) Error() string {
	return fmt.Sprintf("ronso rage %s", r.Overdrive.Name)
}

func (r RonsoRage) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   r.ID,
		Name: r.Name,
	}
}
