package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AeonCommand struct {
	ID                int32
	TopmenuID         *int32
	SubmenuID         *int32
	Name              string                `json:"name"`
	Description       string                `json:"description"`
	Effect            string                `json:"effect"`
	Topmenu           *string               `json:"topmenu"`
	OpenSubmenu       *string               `json:"open_submenu"`
	Cursor            *string               `json:"cursor"`
	PossibleAbilities []PossibleAbilityList `json:"possible_abilities"`
}

func (c AeonCommand) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", c),
		c.Name,
		c.Description,
		c.Effect,
		h.DerefOrNil(c.TopmenuID),
		h.DerefOrNil(c.Cursor),
		h.DerefOrNil(c.SubmenuID),
	}
}

func (c AeonCommand) ToKeyFields() []any {
	return []any{
		c.Name,
	}
}

func (c AeonCommand) GetID() int32 {
	return c.ID
}

func (c AeonCommand) Error() string {
	return fmt.Sprintf("aeon command %s", c.Name)
}

func (c AeonCommand) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   c.ID,
		Name: c.Name,
	}
}

type PossibleAbilityList struct {
	User      string             `json:"user"`
	Abilities []AbilityReference `json:"abilities"`
}

func (pa PossibleAbilityList) Error() string {
	return fmt.Sprintf("possible abilities for %s", pa.User)
}
