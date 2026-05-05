package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type OverdriveCommand struct {
	ID          int32
	CharClassID *int32
	TopmenuID   *int32
	SubmenuID   *int32
	Name        string  `json:"name"`
	Description string  `json:"description"`
	User        string  `json:"user"`
	Rank        int32   `json:"rank"`
	Topmenu     *string `json:"topmenu"`
	OpenSubmenu string  `json:"open_submenu"`
}

func (o OverdriveCommand) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", o),
		o.Name,
		o.Description,
		o.Rank,
		h.DerefOrNil(o.TopmenuID),
		o.OpenSubmenu,
		h.DerefOrNil(o.CharClassID),
		h.DerefOrNil(o.SubmenuID),
	}
}

func (o OverdriveCommand) ToKeyFields() []any {
	return []any{
		o.Name,
	}
}

func (o OverdriveCommand) GetID() int32 {
	return o.ID
}

func (o OverdriveCommand) Error() string {
	return fmt.Sprintf("overdrive command %s", o.Name)
}

func (o OverdriveCommand) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   o.ID,
		Name: o.Name,
	}
}
