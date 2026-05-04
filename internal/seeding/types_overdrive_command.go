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

func (oc OverdriveCommand) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", oc),
		oc.Name,
		oc.Description,
		oc.Rank,
		h.DerefOrNil(oc.TopmenuID),
		oc.OpenSubmenu,
		h.DerefOrNil(oc.CharClassID),
		h.DerefOrNil(oc.SubmenuID),
	}
}

func (oc OverdriveCommand) GetID() int32 {
	return oc.ID
}

func (oc OverdriveCommand) Error() string {
	return fmt.Sprintf("overdrive command %s", oc.Name)
}

func (o OverdriveCommand) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   o.ID,
		Name: o.Name,
	}
}
